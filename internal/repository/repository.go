package repository

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func getAllSamples(c *gin.Context) {
	samples := FindFile
	c.IndentedJSON(http.StatusOK, Sample)
}

func (r *Repository) FindSampleByID(ctx context.Context) (Sample, error) {
	query := `SELECT cover_url, name, sample_file_url, length, bpm, price  FROM samples WHERE name = $1`
	var sample Sample
	err := r.db.QueryRow(ctx, query).Scan(&sample.Name)
	if err != nil {

		return Sample{}, err
	}

	return sample, nil
}
