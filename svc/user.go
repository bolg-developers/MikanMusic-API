package svc

import (
	"github.com/bolg-developers/MikanMusic-API/infra"
	"github.com/bolg-developers/MikanMusic-API/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func CreateUser(u *model.User) error {
	u.ID = uuid.New().String()
	return errors.WithStack(infra.DB().Create(u).Error)
}

func GetUserByID(id string) (*model.UserCore, error) {
	var u model.User
	err := infra.DB().Where("id = ?", id).First(&u).Error
	return &u.UserCore, errors.WithStack(err)
}
