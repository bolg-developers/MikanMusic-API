package svc

import (
	"github.com/bolg-developers/MikanMusic-API/infra"
	"github.com/bolg-developers/MikanMusic-API/model"
	"github.com/pkg/errors"
)

var (
	ErrUnknownGenre = errors.New("不明なジャンルが指定されています")
)

func GetAllGenres() (model.Genres, error) {
	ret := make(model.Genres, 0)
	err := errors.WithStack(infra.DB().Find(&ret).Error)
	return ret, err
}

func GetGenresByMusicID(id string) (model.Genres, error) {
	ret := make(model.Genres, 0)
	err := errors.WithStack(infra.DB().Table("musics").Select("genres.id, genres.name").
		Joins("left join music_genre_relations on music_genre_relations.music_id = musics.id").
		Joins("left join genres on genres.id = music_genre_relations.genre_id").
		Where("musics.id = ?", id).
		Find(&ret).Error)
	return ret, err
}

func ValidateGenres(genres model.Genres) error {
	all, err := GetAllGenres()
	if err != nil {
		return err
	}
	for _, g := range genres {
		var exists bool
		for _, one := range all {
			if one.ID == g.ID {
				exists = true
				break
			}
		}
		if !exists {
			return ErrUnknownGenre
		}
	}
	return nil
}
