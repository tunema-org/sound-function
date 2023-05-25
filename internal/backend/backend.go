package backend

import (
	"github.com/tunema-org/sound-function/internal/clients"
	"github.com/tunema-org/sound-function/internal/config"
	"github.com/tunema-org/sound-function/internal/repository"
)

type Backend struct {
	clients *clients.Clients
	repo    *repository.Repository
	cfg     *config.Config
}

func New(clients *clients.Clients, repo *repository.Repository, cfg *config.Config) *Backend {
	return &Backend{
		clients: clients,
		repo:    repo,
		cfg:     cfg,
	}
}
