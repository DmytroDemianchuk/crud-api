package models

import "github.com/lukinairina90/crud_movies/internal/domain"

type Music struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Performer   string `json:"performer"`
	ReleaseDate int    `json:"release_data"`
	Genre       string `json:"genre"`
}

func (m Music) ToDomain() domain.Movie {
	return domain.Movie{
		ID:          m.ID,
		Name:        m.Name,
		Performer:   m.Performer,
		ReleaseDate: m.ReleaseDate,
		Genre:       m.Genre,
	}
}
