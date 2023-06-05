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

// TODO : categorization
func (r *Repository) GetSamples(ctx context.Context, id int) (model.Sample, usermodel.User, model.Sample_tag, error) {
	query := `SELECT samples.*,users.*,sample_tags FROM samples LEFT JOIN users ON users.id = samples.user_id JOIN sample_tags ON samples.id = sample_tags.sample.id
	WHERE sample_tags.tag_id IN (1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20);`

	var sample model.Sample
	var author usermodel.User
	var sample_tag model.Sample_tag

	err := r.db.QueryRow(ctx, query, id).
		Scan(
			&sample_tag.Tag_ID,
		)
	if err != nil {
		return model.Sample{}, usermodel.User{}, model.Sample_tag{}, err
	}

	return sample, author, sample_tag, nil
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
			&sample.CoverURL,
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
