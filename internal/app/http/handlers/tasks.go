package handlers

import (
	"errors"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type createTaskInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

type createTaskResponse struct {
	Id string `json:"id"`
}

func (h *Handlers) CreateTask(c echo.Context) error {
	var inp createTaskInput
	if err := c.Bind(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := c.Validate(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	res, err := h.services.Tasks.Create(c.Request().Context(), service.CreateTaskInput{
		Title:       inp.Title,
		Description: inp.Description,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &errResponse{Err: err.Error()})
	}

	return c.JSON(http.StatusCreated, &createTaskResponse{Id: res.Hex()})
}

type updateTaskInput struct {
	ID          string `param:"id" validate:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *Handlers) UpdateTask(c echo.Context) error {
	var inp updateTaskInput
	if err := c.Bind(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := c.Validate(&inp); err != nil {
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
	ID string `param:"id" validate:"required"`
}

func (h *Handlers) DeleteTask(c echo.Context) error {
	var inp deleteTaskInput
	if err := c.Bind(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := c.Validate(&inp); err != nil {
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

type getTaskInput struct {
	ID string `param:"id" validate:"required"`
}

type getTaskResponse struct {
	Task *models.Task `json:"task"`
}

func (h *Handlers) GetTask(c echo.Context) error {
	var inp getTaskInput

	if err := c.Bind(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := c.Validate(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	res, err := h.services.Tasks.GetById(c.Request().Context(), inp.ID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, &errResponse{Err: err.Error()})
	}

	return c.JSON(http.StatusOK, &getTaskResponse{Task: res})
}
