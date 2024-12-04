package agent

import (
	"context"
	"log/slog"

	"github.com/AlexBlackNn/authloyalty/client/internal/config/configagent"
	"github.com/AlexBlackNn/authloyalty/client/internal/services/client"
	"github.com/AlexBlackNn/metrics/app/agent/encryption"
)

type clientService interface {
	Register(ctx context.Context)
	Sync(ctx context.Context)
	ChangeData(ctx context.Context)
}

// AppMonitor service consists all service layers.
type App struct {
	clientService clientService
}

// NewAppMonitor creates App.
func New(
	log *slog.Logger,
	cfg *configagent.Config,
) (*App, error) {

	encryptor, err := encryption.NewEncryptor(cfg.CryptoKeyPath)
	if err != nil {
		return nil, err
	}

	client := client.New(
		log,
		cfg,
		encryptor,
	)
	return &App{clientService: client}, nil
}
