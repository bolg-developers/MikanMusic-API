package svc

import (
	"github.com/bolg-developers/MikanMusic-API/infra"
	"github.com/bolg-developers/MikanMusic-API/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func CreateMusic(m *model.Music) error {
	m.ID = uuid.New().String()
	return errors.WithStack(infra.DB().Create(m).Error)
}

// musicとgenreをjoinにして返す
func GetAllMusics() (model.Musics, error) {
	ret := make(model.Musics, 0)
	err := errors.WithStack(infra.DB().Find(&ret).Error)
	return ret, err
}

func DeleteMusic(id string) error {
	return errors.WithStack(infra.DB().Delete(model.Music{}, "id = ?", id).Error)
}
