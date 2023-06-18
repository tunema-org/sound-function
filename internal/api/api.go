package api

import (
	"context"
	"net/http"

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

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	r.GET("/sounds/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})

	r.GET("/sounds", h.ListSamples)
	r.GET("/sounds/search", h.SearchSamples)
	r.GET("/sounds/tags", h.ListTagsAndCategories)
	r.GET("/sounds/categories", h.ListCategories)
	r.GET("/sounds/categories/:id/tags", h.ListCategoryTags)
	r.GET("/sounds/users/:id", h.ListUserSamples)
	r.POST("/sounds", h.CreateSample)
	r.GET("/sounds/:id", h.GetSampleByID)
	r.PATCH("/sounds/:id", h.UpdateSample)
	r.DELETE("/sounds/:id", h.DeleteSample)

	return r
}
