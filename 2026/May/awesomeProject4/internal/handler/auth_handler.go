package handler

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"unicode/utf8"

	"app/internal/auth"
	appmw "app/internal/middleware"
	"app/internal/repository"

	"github.com/labstack/echo/v4"
)

var usernameRE = regexp.MustCompile(`^[a-zA-Z0-9_]{3,32}$`)

type AuthHandler struct {
	users    *repository.UserRepository
	sessions *auth.SessionStore
}

func NewAuthHandler(users *repository.UserRepository, sessions *auth.SessionStore) *AuthHandler {
	return &AuthHandler{users: users, sessions: sessions}
}

func (h *AuthHandler) RegisterPublic(e *echo.Echo) {
	g := e.Group("", appmw.RedirectIfAuth(h.sessions))
	g.GET("/login", h.LoginForm)
	g.POST("/login", h.Login)
	g.GET("/register", h.RegisterForm)
	g.POST("/register", h.Register)
}

func (h *AuthHandler) RegisterProtected(e *echo.Group) {
	e.POST("/logout", h.Logout)
}

func (h *AuthHandler) LoginForm(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", PageData{
		Title: "Login",
		Flash: c.QueryParam("flash"),
	})
}

func (h *AuthHandler) RegisterForm(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", PageData{Title: "Register"})
}

func (h *AuthHandler) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	pd := PageData{Title: "Login", FormUsername: username}

	if username == "" || password == "" {
		pd.Error = "Username and password are required."
		return c.Render(http.StatusBadRequest, "login.html", pd)
	}

	user, err := h.users.GetByUsername(c.Request().Context(), username)
	if err != nil {
		if !errors.Is(err, repository.ErrNotFound) {
			log.Printf("login: %v", err)
			pd.Error = "Something went wrong. Try again."
			return c.Render(http.StatusInternalServerError, "login.html", pd)
		}
		pd.Error = "Invalid username or password."
		return c.Render(http.StatusUnauthorized, "login.html", pd)
	}
	if !auth.CheckPassword(user.PasswordHash, password) {
		pd.Error = "Invalid username or password."
		return c.Render(http.StatusUnauthorized, "login.html", pd)
	}

	if err := h.sessions.SetUserID(c, user.ID); err != nil {
		log.Printf("session: %v", err)
		pd.Error = "Could not start session."
		return c.Render(http.StatusInternalServerError, "login.html", pd)
	}

	return c.Redirect(http.StatusSeeOther, "/?flash=Welcome+back")
}

func (h *AuthHandler) Register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	confirm := c.FormValue("password_confirm")

	pd := PageData{Title: "Register", FormUsername: username}

	if msg := validateRegistration(username, password, confirm); msg != "" {
		pd.Error = msg
		return c.Render(http.StatusBadRequest, "register.html", pd)
	}

	hash, err := auth.HashPassword(password)
	if err != nil {
		log.Printf("hash password: %v", err)
		pd.Error = "Something went wrong. Try again."
		return c.Render(http.StatusInternalServerError, "register.html", pd)
	}

	user, err := h.users.Create(c.Request().Context(), username, hash)
	if errors.Is(err, repository.ErrExists) {
		pd.Error = "Username is already taken."
		return c.Render(http.StatusConflict, "register.html", pd)
	}
	if err != nil {
		log.Printf("register: %v", err)
		pd.Error = "Could not create account."
		return c.Render(http.StatusInternalServerError, "register.html", pd)
	}

	if err := h.sessions.SetUserID(c, user.ID); err != nil {
		log.Printf("session: %v", err)
		pd.Error = "Account created but login failed. Please sign in."
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	return c.Redirect(http.StatusSeeOther, "/?flash=Account+created")
}

func (h *AuthHandler) Logout(c echo.Context) error {
	if err := h.sessions.Clear(c); err != nil {
		log.Printf("logout: %v", err)
	}
	return c.Redirect(http.StatusSeeOther, "/login?flash=Logged+out")
}

func validateRegistration(username, password, confirm string) string {
	if username == "" || password == "" || confirm == "" {
		return "All fields are required."
	}
	if !usernameRE.MatchString(username) {
		return "Username must be 3–32 characters: letters, numbers, underscore."
	}
	if utf8.RuneCountInString(password) < 6 {
		return "Password must be at least 6 characters."
	}
	if password != confirm {
		return "Passwords do not match."
	}
	return ""
}
