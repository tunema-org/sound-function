package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tunema-org/sound-function/internal/backend"
)

type M gin.H

type handler struct {
	backend *backend.Backend
}

func NewHandler(ctx context.Context, backend *backend.Backend) *gin.Engine {
	h := &handler{
		backend: backend,
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/sounds/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})

	r.GET("/sounds", h.ListSamples)
	r.GET("/sounds/search", h.SearchSamples)
	r.GET("/sounds/tags", h.ListTags)
	r.GET("/sounds/categories", h.ListCategories)
	r.GET("/sounds/categories/:id/tags", h.ListCategoryTags)
	r.GET("/sounds/users/:id", h.ListUserSamples)
	r.POST("/sounds", h.CreateSample)
	r.GET("/sounds/:id", h.GetSampleByID)
	r.PATCH("/sounds/:id", h.UpdateSample)
	r.DELETE("/sounds/:id", h.DeleteSample)

	return r
}
