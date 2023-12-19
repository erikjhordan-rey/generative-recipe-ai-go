package main

import (
	"example/gemini-recipes/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/healthcheck", controllers.HealthCheck)
	router.GET("/recipes", controllers.GetRecipe)
	router.POST("/recipes", controllers.CreateRecipe)
	router.Run()
}
