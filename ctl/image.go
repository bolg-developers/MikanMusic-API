package ctl

import (
	"log"
	"net/http"

	"github.com/bolg-developers/MikanMusic-API/svc"
	"github.com/gin-gonic/gin"
)

func CreateIcon(c *gin.Context) {
	mf, err := c.MultipartForm()
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Printf("BadRequest: %+v", err)
		return
	}

	hflist, ok := mf.File["image"]
	if !ok {
		c.Status(http.StatusBadRequest)
		log.Printf("BadRequest: imageが指定されていません")
		return
	}

	f, err := hflist[0].Open()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		log.Printf("InternalServerError: %+v", err)
		return
	}

	url, err := svc.CreateImage(f, 400, 400)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		log.Printf("InternalServerError: %+v", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"url": url})
}
