package ctl

import (
	"log"
	"github.com/bolg-developers/MikanMusic-API/svc"
	"github.com/gin-gonic/gin"
)

func GetAllGenres(c *gin.Context) {
	genres, err := svc.GetAllGenres()
	if err != nil {
		c.Status(500)
		log.Printf("InternalServerError: %+v", err)
		return
	}
	c.JSON(200, gin.H{"genres": genres})
}