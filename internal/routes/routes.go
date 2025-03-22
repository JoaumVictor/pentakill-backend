package routes

import (
	"github.com/JoaumVictor/pentakill-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/games", handlers.GetGames)
	r.POST("/games", handlers.CreateGame)
	r.PUT("/games/:id", handlers.UpdateGame)
	r.DELETE("/games/:id", handlers.DeleteGame)

	return r
}
