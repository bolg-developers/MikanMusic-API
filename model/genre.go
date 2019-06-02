package model

type Genre struct {
	ID   string `json:"id"`
	Name string `json:"name" binding:"max=255"`
}

type Genres []*Genre
