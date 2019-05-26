package ctl

import (
	"log"
	"net/http"

	"github.com/bolg-developers/MikanMusic-API/model"
	"github.com/bolg-developers/MikanMusic-API/svc"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func SignUp(c *gin.Context) {
	u := new(model.User)
	if err := c.BindJSON(u); err != nil {
		c.Status(http.StatusBadRequest)
		log.Printf("BadRequest: %+v", errors.WithStack(err))
		return
	}

	hash, err := svc.HashPassword(u.Password)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		log.Printf("InternalServerError: %+v", err)
		return
	}
	u.Password = hash

	u.LogoutAt = nil

	if err := svc.CreateUser(u); err != nil {
		c.Status(http.StatusInternalServerError)
		log.Printf("InternalServerError: %+v", err)
		return
	}

	token, err := svc.CreateUserToken(u.ID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		log.Printf("InternalServerError: %+v", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Status(http.StatusBadRequest)
		log.Printf("BadRequest: idが見つかりません")
		return
	}

	user, err := svc.GetUserByID(id)
	if err != nil {
		if gorm.IsRecordNotFoundError(errors.Cause(err)) {
			c.Status(http.StatusNotFound)
			log.Printf("NotFound: %+v", err)
			return
		}
		c.Status(http.StatusInternalServerError)
		log.Printf("InternalServerError: %+v", err)
		return
	}

	c.JSON(http.StatusOK, user)
}
