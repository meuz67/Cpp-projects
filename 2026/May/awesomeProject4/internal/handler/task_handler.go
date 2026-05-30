package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"app/internal/models"
	"app/internal/repository"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	repo *repository.TaskRepository
}

func NewTaskHandler(repo *repository.TaskRepository) *TaskHandler {
	return &TaskHandler{repo: repo}
}

func (h *TaskHandler) Register(g *echo.Group) {
	g.GET("/", h.List)
	g.GET("/tasks/new", h.NewForm)
	g.POST("/tasks", h.Create)
	g.GET("/tasks/:id/edit", h.EditForm)
	g.POST("/tasks/:id", h.Update)
	g.POST("/tasks/:id/delete", h.Delete)
}

func (h *TaskHandler) List(c echo.Context) error {
	userID := userIDFromContext(c)
	tasks, err := h.repo.List(c.Request().Context(), userID)
	if err != nil {
		log.Printf("list tasks: %v", err)
		pd := pageFromContext(c, "Tasks")
		pd.Error = "Failed to load tasks."
		return c.Render(http.StatusInternalServerError, "index.html", pd)
	}

	pd := pageFromContext(c, "Tasks")
	pd.Tasks = tasks
	pd.Flash = c.QueryParam("flash")
	return c.Render(http.StatusOK, "index.html", pd)
}

func (h *TaskHandler) NewForm(c echo.Context) error {
	pd := pageFromContext(c, "New Task")
	pd.Task = &models.Task{}
	return c.Render(http.StatusOK, "form.html", pd)
}

func (h *TaskHandler) Create(c echo.Context) error {
	userID := userIDFromContext(c)
	title := c.FormValue("title")
	description := c.FormValue("description")

	if title == "" {
		pd := pageFromContext(c, "New Task")
		pd.Task = &models.Task{Title: title, Description: description}
		pd.Error = "Title is required."
		return c.Render(http.StatusBadRequest, "form.html", pd)
	}

	if _, err := h.repo.Create(c.Request().Context(), userID, title, description); err != nil {
		log.Printf("create task: %v", err)
		pd := pageFromContext(c, "New Task")
		pd.Task = &models.Task{Title: title, Description: description}
		pd.Error = "Failed to create task."
		return c.Render(http.StatusInternalServerError, "form.html", pd)
	}

	return c.Redirect(http.StatusSeeOther, "/?flash=Task+created")
}

func (h *TaskHandler) EditForm(c echo.Context) error {
	userID := userIDFromContext(c)
	id, err := parseID(c)
	if err != nil {
		return err
	}

	task, err := h.repo.GetByID(c.Request().Context(), userID, id)
	if errors.Is(err, repository.ErrNotFound) {
		return c.Redirect(http.StatusSeeOther, "/?flash=Task+not+found")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to load task")
	}

	pd := pageFromContext(c, "Edit Task")
	pd.Task = task
	return c.Render(http.StatusOK, "form.html", pd)
}

func (h *TaskHandler) Update(c echo.Context) error {
	userID := userIDFromContext(c)
	id, err := parseID(c)
	if err != nil {
		return err
	}

	title := c.FormValue("title")
	description := c.FormValue("description")

	if title == "" {
		task, _ := h.repo.GetByID(c.Request().Context(), userID, id)
		if task == nil {
			task = &models.Task{ID: id, Description: description}
		}
		task.Title = title
		task.Description = description
		pd := pageFromContext(c, "Edit Task")
		pd.Task = task
		pd.Error = "Title is required."
		return c.Render(http.StatusBadRequest, "form.html", pd)
	}

	_, err = h.repo.Update(c.Request().Context(), userID, id, title, description)
	if errors.Is(err, repository.ErrNotFound) {
		return c.Redirect(http.StatusSeeOther, "/?flash=Task+not+found")
	}
	if err != nil {
		log.Printf("update task: %v", err)
		pd := pageFromContext(c, "Edit Task")
		pd.Task = &models.Task{ID: id, Title: title, Description: description}
		pd.Error = "Failed to update task."
		return c.Render(http.StatusInternalServerError, "form.html", pd)
	}

	return c.Redirect(http.StatusSeeOther, "/?flash=Task+updated")
}

func (h *TaskHandler) Delete(c echo.Context) error {
	userID := userIDFromContext(c)
	id, err := parseID(c)
	if err != nil {
		return err
	}

	if err := h.repo.Delete(c.Request().Context(), userID, id); errors.Is(err, repository.ErrNotFound) {
		return c.Redirect(http.StatusSeeOther, "/?flash=Task+not+found")
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete task")
	}

	return c.Redirect(http.StatusSeeOther, "/?flash=Task+deleted")
}

func userIDFromContext(c echo.Context) int {
	id, _ := c.Get("userID").(int)
	return id
}

func parseID(c echo.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "invalid task id")
	}
	return id, nil
}
