package svc

import (
	"github.com/bolg-developers/MikanMusic-API/infra"
	"github.com/bolg-developers/MikanMusic-API/model"
	"github.com/pkg/errors"
)

func CreateMusicGenreRelation(mgr *model.MusicGenreRelation) error {
	return errors.WithStack(infra.DB().Create(mgr).Error)
}

func DeleteMusicGenreRelationByMusicID(id string) error {
	return errors.WithStack(infra.DB().Delete(model.MusicGenreRelation{}, "music_id = ?", id).Error)
}
