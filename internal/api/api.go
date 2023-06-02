package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/tunema-org/sound-function/internal/backend"
)

type handler struct {
	backend *backend.Backend
}

func NewHandler(ctx context.Context, backend *backend.Backend) *gin.Engine {
	h := &handler{
		backend: backend,
	}

	r := gin.Default()

	r.GET("/sounds", h.ListSamples)
	r.GET("/sounds/search", h.SearchSample)
	r.GET("/sounds/tags", h.ListTags)
	r.GET("/sounds/categories", h.ListCategories)
	r.GET("/sounds/categories/:id/tags", h.ListCategoryTags)
	r.POST("/sounds", h.CreateSample)
	r.GET("/sounds/:id", h.GetSampleByID)
	r.PATCH("/sounds/:id", h.UpdateSample)
	r.DELETE("/sounds/:id", h.DeleteSample)

	return r
}
