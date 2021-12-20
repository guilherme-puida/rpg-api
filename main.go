package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpg/controllers"
	"rpg/models"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/characters", controllers.FindCharacters)
	r.POST("/characters", controllers.CreateCharacter)
	r.GET("/characters/:id", controllers.FindCharacter)
	r.PATCH("/characters/:id", controllers.UpdateCharacter)
	r.DELETE("/characters/:id", controllers.DeleteCharacter)

	models.ConnectDatabase()

	r.Run()
}
