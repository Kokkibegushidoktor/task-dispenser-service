package handlers

import (
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type addLevelInput struct {
	TaskId        string `json:"taskId"`
	Title         string `json:"title"`
	VarQuestCount int    `json:"varQuestCount"`
}

func (h *Handlers) AddTaskLevel(c echo.Context) error {
	var inp addLevelInput
	if err := c.Bind(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := validateAddLevelInput(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	res, err := h.services.Levels.Create(c.Request().Context(), service.CreateLevelInput{
		TaskId:        inp.TaskId,
		Title:         inp.Title,
		VarQuestCount: inp.VarQuestCount,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"id": res.Hex(),
	})
}

type updateLevelInput struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	VarQuestCount int    `json:"varQuestCount"`
}

func (h *Handlers) UpdateTaskLevel(c echo.Context) error {
	var inp updateLevelInput
	if err := c.Bind(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := validateUpdateLevelInput(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	err := h.services.Levels.Update(c.Request().Context(), service.UpdateLevelInput{
		ID:            inp.ID,
		Title:         inp.Title,
		VarQuestCount: inp.VarQuestCount,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	return c.JSON(http.StatusOK, &emptyResponse{})
}

type deleteLevelInput struct {
	ID string `json:"id"`
}

func (h *Handlers) DeleteTaskLevel(c echo.Context) error {
	var inp deleteLevelInput
	if err := c.Bind(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := validateDeleteLevelInput(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := h.services.Levels.Delete(c.Request().Context(), inp.ID); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	return c.JSON(http.StatusOK, &emptyResponse{})
}
