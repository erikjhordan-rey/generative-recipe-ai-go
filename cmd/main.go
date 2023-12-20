package main

import (
	"example/gemini-recipes/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/healthcheck", controllers.HealthCheck)
	router.POST("/recipes", controllers.GenerateRecipe)
	router.GET("/recipes", controllers.GetRecipe)
	router.Run()
}
