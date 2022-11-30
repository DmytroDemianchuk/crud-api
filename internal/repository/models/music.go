package models

import "github.com/dmytrodemianchuk/crud-api/internal/domain"

type Music struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Performer   string `json:"performer"`
	ReleaseDate int    `json:"release_data"`
	Genre       string `json:"genre"`
}

func (m Music) ToDomain() domain.Music {
	return domain.Music{
		ID:          m.ID,
		Name:        m.Name,
		Performer:   m.Performer,
		ReleaseDate: m.ReleaseDate,
		Genre:       m.Genre,
	}
}
