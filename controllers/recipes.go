package controllers

import (
	"encoding/json"
	"fmt"
	"io"

	"log"
	"net/http"
	"os"
	"path/filepath"

	"example/gemini-recipes/models"
	"example/gemini-recipes/services"

	"github.com/gin-gonic/gin"
)

type FoodImageUrlInput struct {
	Image string `json:"url_image" binding:"required"`
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Health check successful! The server is up and running smoothly."})

}

func GetRecipe(c *gin.Context) {

	imageBytes, err := os.ReadFile(filepath.Join("assets", "pizza.jpeg"))
	if err != nil {
		log.Fatal("error reading image: ", err)
	}
	resp := services.AnalyzeFoodImage(imageBytes)
	if resp == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": resp})
	}

	var recipe models.Recipe
	if err := json.Unmarshal([]byte(fmt.Sprintf("%+v", resp)), &recipe); err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

func GenerateRecipe(c *gin.Context) {
	var imageUrl FoodImageUrlInput
	if error := c.ShouldBindJSON(&imageUrl); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	imageBytes, err := getImageBytes(imageUrl.Image)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	resp := services.AnalyzeFoodImage(imageBytes)
	if resp == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": resp})
	}

	var recipe models.Recipe
	if err := json.Unmarshal([]byte(fmt.Sprintf("%+v", resp)), &recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

func getImageBytes(URL string) ([]byte, error) {
	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	imageBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return imageBytes, nil
}
