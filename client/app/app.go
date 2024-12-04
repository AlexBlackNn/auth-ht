package app

import (
	log "log/slog"

	"github.com/AlexBlackNn/authloyalty/client/internal/config"
	"github.com/AlexBlackNn/authloyalty/client/internal/logger"
	"github.com/AlexBlackNn/authloyalty/client/internal/service"
)

type servicer interface {
	Start() error
}

type App struct {
	Сfg     *config.Config
	Log     *log.Logger
	Service servicer
}

func New() (*App, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}
	log := logger.New(cfg.Env)

	return &App{
		Сfg:     cfg,
		Log:     log,
		Service: service.New(cfg, log),
	}, nil
}

func (a *App) Start() error {
	err := a.Service.Start()
	if err != nil {
		return err
	}
	return nil
}
