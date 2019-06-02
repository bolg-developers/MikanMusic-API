package model

type MusicAndGenre struct {
	Music
	Genres Genres `json:"genres"`
}

type MusicAndGenreList []*MusicAndGenre
