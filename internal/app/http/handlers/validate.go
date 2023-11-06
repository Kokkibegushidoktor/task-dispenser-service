package handlers

import (
	"errors"
	"strings"
)

func validateSignInInput(inp *signInInput) error {
	if strings.TrimSpace(inp.Username) == "" &&
		strings.TrimSpace(inp.Password) == "" {
		return errors.New("empty request")
	}

	if strings.TrimSpace(inp.Username) == "" {
		return errors.New("empty username")
	}

	if strings.TrimSpace(inp.Password) == "" {
		return errors.New("empty password")
	}

	return nil
}

func validateCreateTaskInput(inp *createTaskInput) error {
	if strings.TrimSpace(inp.Title) == "" {
		return errors.New("empty title")
	}

	return nil
}

func validateUpdateTaskInput(inp *updateTaskInput) error {
	if strings.TrimSpace(inp.ID) == "" {
		return errors.New("empty id")
	}

	return nil
}

func validateDeleteTaskInput(inp *deleteTaskInput) error {
	if strings.TrimSpace(inp.ID) == "" {
		return errors.New("empty id")
	}

	return nil
}

func validateAddLevelInput(inp *addLevelInput) error {
	if strings.TrimSpace(inp.TaskId) == "" {
		return errors.New("empty task id")
	}

	if inp.VarQuestCount < 1 {
		return errors.New("invalid var question count")
	}

	return nil
}

func validateUpdateLevelInput(inp *updateLevelInput) error {
	if strings.TrimSpace(inp.ID) == "" {
		return errors.New("empty id")
	}

	if inp.VarQuestCount < 0 {
		return errors.New("invalid var question count")
	}

	return nil
}

func validateDeleteLevelInput(inp *deleteLevelInput) error {
	if strings.TrimSpace(inp.ID) == "" {
		return errors.New("empty id")
	}

	return nil
}

func validateAddQuestionInput(inp *addQuestionInput) error {
	if strings.TrimSpace(inp.LevelID) == "" {
		return errors.New("empty level id")
	}

	return nil
}

func validateUpdateQuestionInput(inp *updateQuestionInput) error {
	if strings.TrimSpace(inp.ID) == "" {
		return errors.New("empty id")
	}

	return nil
}

func validateDeleteQuestionInput(inp *deleteQuestionInput) error {
	if strings.TrimSpace(inp.ID) == "" {
		return errors.New("empty id")
	}

	return nil
}
