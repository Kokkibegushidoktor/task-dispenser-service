package http

import (
	"context"
	"errors"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/pkg/auth"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"

	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/app/http/handlers"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/config"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/repository"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/tech/closer"
)

type Server struct {
	server   *echo.Echo
	handlers *handlers.Handlers
	cfg      *config.Config
}

func New(cfg *config.Config, repo repository.Repository) *Server {
	server := echo.New()
	server.Use(middleware.Recover())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	server.HideBanner = true
	server.HidePort = true

	manager, err := auth.NewManager(cfg.JwtSecret)
	if err != nil {
		log.Error().Msgf("Error creating a token manager, err: %v", err)
		manager, _ = auth.NewManager("amogus")
	}

	return &Server{
		server:   server,
		handlers: handlers.New(repo, manager),
		cfg:      cfg,
	}
}

func (s *Server) Start() {
	s.setupRoutes()

	go func() {
		log.Info().Msgf("Starting listening http server at %s", s.cfg.HttpAddr)
		err := s.server.Start(s.cfg.HttpAddr)
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err)
		}
	}()

	closer.Add(s.close)
}

func (s *Server) close() error {
	if err := s.server.Shutdown(context.TODO()); err != nil {
		log.Error().Msgf("Error stopping http server, err: %v", err)
		return err
	}

	log.Info().Msgf("Http server shutdown is done")

	return nil
}
