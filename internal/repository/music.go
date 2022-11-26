package psql

import (
	"context"
	"database/sql"

	"github.com/dmytrodemianchuk/internal/domain"
	"github.com/dmytrodemiancuk/crud-api/internal/repository/models"
)

type Books struct {
	db *sql.DB
}

func NewBooks(db *sql.DB) *Books {
	return &Books{db}
}

func (m *Music) List(ctx context.Context) (domain.ListMusic, error) {
	var list []models.Movie
	if err := m.db.SelectContext(ctx, &list, "SELECT * FROM music"); err != nil {
		return nil, err
	}

	dlist := make(domain.ListMovie, 0, len(list))
	for _, music := range list {
		dlist = append(dlist, music.ToDomain())
	}

	return dlist, nil
}

func (b *Music) Get(ctx context.Context, id int64) (domain.Music, error) {
	var music domain.Music
	if err := m.db.GetContext(ctx, &music, "SELECT * FROM  movie WHERE id=$1", id); err != nil {
		return domain.Movie{}, err
	}

	return music.ToDomain(), nil
}

func (m *Music) Create(ctx context.Context, movie domain.Music) (domain.Music, error) {
	mMusic := models.Music{
		Name:        music.Name,
		Performer:   music.Performer,
		ReleaseDate: music.ReleaseData,
		Genre:       music.Genre,
	}

	if err := m.db.QueryRowxContext(ctx, "INSERT INTO movie (name, description, production_year, genre, actors, poster) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *", mMovie.Name, mMovie.Description, mMovie.ProductionYear, mMovie.Genre, mMovie.Actors, mMovie.Poster).StructScan(&mMovie); err != nil {
		return domain.Movie{}, err
	}

	return mMusic.ToDomain(), nil
}

func (m *Music) Update(ctx context.Context, id int, movie domain.Movie) (domain.Movie, error) {
	mMusic := models.Music{
		Name:        music.Name,
		Performer:   music.Performer,
		ReleaseDate: music.ReleaseData,
		Genre:       music.Genre,
	}

	if err := m.db.QueryRowxContext(ctx, "UPDATE movie SET name=$1, performer=$2, release_date=$3, genre=$4, WHERE id=$7 RETURNING *",
		mMusic.Name, mMusic.Performer, mMusic.ReleaseData, mMusic.Genre, id).StructScan(&mMusic); err != nil {
		return domain.Movie{}, err
	}

	return mMusic.ToDomain(), nil
}
func (m *Music) Delete(ctx context.Context, id int) error {
	if _, err := m.db.ExecContext(ctx, "DELETE FROM music WHERE id=$1", id); err != nil {
		return err
	}
	return nil
}
