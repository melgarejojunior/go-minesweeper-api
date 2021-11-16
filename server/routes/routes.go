package routes

import (
	"minesweeper/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		minesweeper := main.Group("minesweeper")
		{
			minesweeper.POST("/start", controllers.ConfigMinesweeper)
		}

		game := main.Group("game")
		{
			game.POST("/:id/play", controllers.Play)
		}
	}
	return router
}
