package domain

type ListMovie []Music

type Music struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Performer   string `json:"performer"`
	ReleaseDate int    `json:"release_data"`
	Genre       string `json:"genre"`
}
