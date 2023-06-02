package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tunema-org/sound-function/model"
	usermodel "github.com/tunema-org/user-function/model"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

// TODO: update
func (r *Repository) FindSampleByID(ctx context.Context, id int) (model.Sample, usermodel.User, error) {
	query := `SELECT samples.*, users.* FROM samples LEFT JOIN users ON users.id = samples.user_id WHERE id = $1`
	var sample model.Sample
	var author usermodel.User
	err := r.db.QueryRow(ctx, query, id).
		Scan(
			&sample.ID,
			&sample.UserID,
			&sample.Name,
			&sample.BPM,
			&sample.Key,
			&sample.KeyScale,
			&sample.Time,
			&sample.FileURL,
			&sample.CoverUrl,
			&sample.Price,
			&sample.CreatedAt,
			&author.ID,
			&author.Username,
			&author.Email,
			nil, // ignore password
			&author.ProfileImgURL)
	if err != nil {

		return model.Sample{}, usermodel.User{}, err
	}

	return sample, author, nil
}
