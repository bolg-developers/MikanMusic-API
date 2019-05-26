package router

import (
	"github.com/bolg-developers/MikanMusic-API/ctl"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Router() *gin.Engine {
	return router
}

func init() {
	r := gin.Default()

	r.GET("/alive", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "🍊",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.POST("/signup", ctl.SignUp)

		v1.GET("/users/:id", ctl.GetUserByID)

		v1.POST("/icons", ctl.CreateIcon)

	}
	router = r
}
