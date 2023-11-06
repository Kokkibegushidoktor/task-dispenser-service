package handlers

import (
	"errors"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type createTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *Handlers) CreateTask(c echo.Context) error {
	var inp createTaskInput
	if err := c.Bind(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := validateCreateTaskInput(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	res, err := h.services.Tasks.Create(c.Request().Context(), service.CreateTaskInput{
		Title:       inp.Title,
		Description: inp.Description,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &errResponse{Err: err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"id": res.Hex(),
	})
}

type updateTaskInput struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *Handlers) UpdateTask(c echo.Context) error {
	var inp updateTaskInput
	if err := c.Bind(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := validateUpdateTaskInput(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	err := h.services.Tasks.Update(c.Request().Context(), service.UpdateTaskInput{
		ID:          inp.ID,
		Title:       inp.Title,
		Description: inp.Description,
	})
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, &errResponse{Err: err.Error()})
	}

	return c.JSON(http.StatusOK, &emptyResponse{})
}

type deleteTaskInput struct {
	ID string `json:"id"`
}

func (h *Handlers) DeleteTask(c echo.Context) error {
	var inp deleteTaskInput
	if err := c.Bind(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := validateDeleteTaskInput(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := h.services.Tasks.Delete(c.Request().Context(), inp.ID); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, &errResponse{Err: err.Error()})
	}

	return c.JSON(http.StatusOK, &emptyResponse{})
}
