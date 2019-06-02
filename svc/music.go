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

func DeleteMusic(id string) error {
	return errors.WithStack(infra.DB().Delete(model.Music{}, "id = ?", id).Error)
}
