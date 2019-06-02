package ctl

import (
	"log"

	"github.com/bolg-developers/MikanMusic-API/svc"
	"github.com/gin-gonic/gin"
)

func CreateAudio(c *gin.Context) {
	mf, err := c.MultipartForm()
	if err != nil {
		c.Status(400)
		log.Printf("BadRequest: %+v", err)
		return
	}

	hflist, ok := mf.File["audio"]
	if !ok {
		c.Status(400)
		log.Printf("BadRequest: audioが指定されていません")
		return
	}

	f, err := hflist[0].Open()
	if err != nil {
		c.Status(500)
		log.Printf("InternalServerError: %+v", err)
		return
	}
	defer f.Close()

	url, err := svc.CreateAudio(f)
	if err != nil {
		c.Status(500)
		log.Printf("InternalServerError: %+v", err)
		return
	}
	c.JSON(201, gin.H{"url": url})
}
