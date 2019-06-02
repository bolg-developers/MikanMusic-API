package model

import "time"

type Music struct {
	ID          string    `json:"id"`
	ArtistID    string    `json:"artistID" binding:"required,uuid4"`
	LyristID    *string   `json:"lyristID" binding:"omitempty,uuid4"`
	ComposerID  *string   `json:"composerID" binding:"omitempty,uuid4"`
	ArrangerID  *string   `json:"arrangerID" binding:"omitempty,uuid4"`
	Name        string    `json:"name" binding:"required,max=255"`
	Lyric       string    `json:"lyric" binding:"max=2000"`
	Description string    `json:"description" binding:"max=500"`
	CountLike   uint      `json:"countLike" binding:"eq=0"`
	CountListen uint64    `json:"countListen" binding:"eq=0"`
	ResourceURL string    `json:"resourceURL" binding:"required,url,max=255"`
	ArtworkURL  string    `json:"artworkURL" binding:"omitempty,url,max=255"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
