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
		v1.POST("/signin", ctl.SignIn)

		v1.GET("/users/:id", ctl.GetUserByID)
		v1.GET("/users", ctl.GetAllUsers)

		// ファイルアップロード系エンドポイント
		v1.POST("/icons", ctl.CreateIcon)
		v1.POST("/audios", ctl.CreateAudio)

		v1.POST("/musics", ctl.CreateMusic)
		v1.GET("/musics", ctl.GetAllMusics)
		v1.PUT("/musics/:id/incCountListen", ctl.IncrementMusicCntListen)

		v1.GET("/genres", ctl.GetAllGenres)
	}
	router = r
}
