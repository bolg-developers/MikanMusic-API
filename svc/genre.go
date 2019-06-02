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

func ValidateGenres(genres model.Genres) error {
	all, err := GetAllGenres()
	if err != nil {
		return err
	}
	for _, g := range genres {
		var exists bool
		for _, one := range all {
			if one == g {
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
