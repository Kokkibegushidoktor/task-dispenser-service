package handlers

import (
	"bytes"
	"fmt"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"os"
)

const (
	maxUploadSize = 5 << 20 // 5 mb
)

var acceptedTypes = map[string]struct{}{
	"image/png":  {},
	"image/jpeg": {},
	"image/webp": {},
	"image/bmp":  {},
}

func (h *Handlers) UploadImage(c echo.Context) error {
	c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, maxUploadSize)

	file, fileHeader, err := c.Request().FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Error().Msgf("error closing file, err: %v", err)
		}
	}()

	buffer := make([]byte, fileHeader.Size)

	if _, err = file.Read(buffer); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	contentType := http.DetectContentType(buffer)
	if _, ex := acceptedTypes[contentType]; !ex {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: "file type not accepted"})
	}

	t, _ := c.Get(userCtx).(*jwt.Token)
	user, err := t.Claims.GetSubject()
	if err != nil {

	}
	tmpFilename := fmt.Sprintf("%s.%s", user, fileHeader.Filename)

	err = func() error {
		f, err := os.OpenFile(tmpFilename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &errResponse{Err: "could not create tmp file"})
		}

		defer func() {
			if err = f.Close(); err != nil {
				log.Error().Msgf("error closing temp file, err: %v", err)
			}
		}()

		if _, err = io.Copy(f, bytes.NewReader(buffer)); err != nil {
			return c.JSON(http.StatusInternalServerError, &errResponse{Err: "could not write to tmp file"})
		}

		return nil
	}()
	if err != nil {
		return err
	}

	res, err := h.services.Files.SaveAndUpload(c.Request().Context(), &models.File{
		Name:        tmpFilename,
		Type:        models.Image,
		ContentType: contentType,
		Size:        fileHeader.Size,
		Uploader:    user,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &errResponse{Err: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"url": res,
	})
}
