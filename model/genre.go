package model

type Genre struct {
	ID   string `json:"id"`
	Name string `json:"json" binding:"max=255"`
}

type Genres []*Genre
