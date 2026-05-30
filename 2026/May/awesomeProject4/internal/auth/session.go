package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

const sessionName = "taskcrud_session"
const userIDKey = "user_id"

type SessionStore struct {
	store *sessions.CookieStore
}

func NewSessionStore(secret string) *SessionStore {
	store := sessions.NewCookieStore([]byte(secret))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 7,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	return &SessionStore{store: store}
}

func (s *SessionStore) Get(c echo.Context) (*sessions.Session, error) {
	return s.store.Get(c.Request(), sessionName)
}

func (s *SessionStore) SetUserID(c echo.Context, userID int) error {
	sess, err := s.Get(c)
	if err != nil {
		return err
	}
	sess.Values[userIDKey] = userID
	return sess.Save(c.Request(), c.Response())
}

func (s *SessionStore) UserID(c echo.Context) (int, bool) {
	sess, err := s.Get(c)
	if err != nil {
		return 0, false
	}
	v, ok := sess.Values[userIDKey]
	if !ok {
		return 0, false
	}
	switch id := v.(type) {
	case int:
		return id, true
	case int64:
		return int(id), true
	case float64:
		return int(id), true
	default:
		return 0, false
	}
}

func (s *SessionStore) Clear(c echo.Context) error {
	sess, err := s.Get(c)
	if err != nil {
		return err
	}
	sess.Options.MaxAge = -1
	delete(sess.Values, userIDKey)
	return sess.Save(c.Request(), c.Response())
}
