package model

type MusicReq struct {
	Music
	Genres Genres `json:"genres"`
}
