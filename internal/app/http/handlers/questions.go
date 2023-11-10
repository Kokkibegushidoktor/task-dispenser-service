package handlers

import (
	"errors"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type addQuestionInput struct {
	LevelID     string `json:"levelId" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

func (h *Handlers) AddQuestion(c echo.Context) error {
	var inp addQuestionInput
	if err := c.Bind(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := c.Validate(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	res, err := h.services.Questions.Create(c.Request().Context(), service.AddQuestionInput{
		LevelID:     inp.LevelID,
		Title:       inp.Title,
		Description: inp.Description,
	})
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, &errResponse{Err: err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"id": res.Hex(),
	})
}

type updateQuestionInput struct {
	ID          string `json:"id" validate:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *Handlers) UpdateQuestion(c echo.Context) error {
	var inp updateQuestionInput
	if err := c.Bind(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := c.Validate(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := h.services.Questions.Update(c.Request().Context(), service.UpdateQuestionInput{
		ID:          inp.ID,
		Title:       inp.Title,
		Description: inp.Description,
	}); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, &errResponse{Err: err.Error()})
	}

	return c.JSON(http.StatusOK, &emptyResponse{})
}

type deleteQuestionInput struct {
	ID string `json:"id" validate:"required"`
}

func (h *Handlers) DeleteQuestion(c echo.Context) error {
	var inp deleteQuestionInput
	if err := c.Bind(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := c.Validate(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := h.services.Questions.Delete(c.Request().Context(), inp.ID); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, &errResponse{Err: err.Error()})
	}

	return c.JSON(http.StatusOK, &emptyResponse{})
}
