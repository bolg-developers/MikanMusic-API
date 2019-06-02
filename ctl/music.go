package ctl

import (
	"log"

	"github.com/bolg-developers/MikanMusic-API/model"
	"github.com/bolg-developers/MikanMusic-API/svc"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func CreateMusic(c *gin.Context) {
	var mr model.MusicReq
	if err := c.BindJSON(&mr); err != nil {
		c.Status(400)
		log.Printf("BadRequest: %+v", errors.WithStack(err))
		return
	}

	if err := svc.ValidateGenres(mr.Genres); err != nil {
		if err == svc.ErrUnknownGenre {
			c.Status(400)
			log.Printf("BadRequest: %+v", err)
			return
		}
		c.Status(500)
		log.Printf("InternalServerError: %+v", err)
		return
	}

	if err := svc.CreateMusic(&mr.Music); err != nil {
		c.Status(500)
		log.Printf("InternalServerError: %+v", err)
		return
	}

	for _, g := range mr.Genres {
		if err := svc.CreateMusicGenreRelation(&model.MusicGenreRelation{
			MusicID: mr.ID,
			GenreID: g.ID,
		}); err != nil {
			log.Printf("InternalServerError: %+v", err)
			if err := svc.DeleteMusic(mr.ID); err != nil {
				log.Printf("InternalServerError: %+v", err)
			}
			if err := svc.DeleteMusicGenreRelationByMusicID(mr.ID); err != nil {
				log.Printf("InternalServerError: %+v", err)
			}
			c.Status(500)
			return
		}
	}

	c.Status(201)
}
