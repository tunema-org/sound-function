package api

import (
	"context"

	"github.com/tunema-org/sound-function/internal/backend"
)

type handler struct {
	backend *backend.Backend
}

func NewHandler(ctx context.Context, backend *backend.Backend) func(event)
