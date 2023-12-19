package controllers

import (
	"encoding/json"
	"fmt"
	"io"

	"log"
	"net/http"

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
	resp := services.GenerateRecipe(nil)
	if resp == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": resp})
	}

	var recipe models.Recipe
	if err := json.Unmarshal([]byte(fmt.Sprintf("%+v", resp)), &recipe); err != nil { // Parse []byte to go struct pointer
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

func CreateRecipe(c *gin.Context) {
	var imageUrl FoodImageUrlInput
	if error := c.ShouldBindJSON(&imageUrl); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	imageBytes, err := loadImageFromURL(imageUrl.Image)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	resp := services.GenerateRecipe(imageBytes)
	if resp == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": resp})
	}

	var recipe models.Recipe
	if err := json.Unmarshal([]byte(fmt.Sprintf("%+v", resp)), &recipe); err != nil { // Parse []byte to go struct pointer
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, recipe)

	//c.JSON(http.StatusOK, gin.H{"status": "success", "message": imageUrl.Image})
}

func loadImageFromURL(URL string) ([]byte, error) {

	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	//  img, _, err := image.Decode(response.Body)
	//  if err != nil {
	//	return nil, err
	// }

	imageData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Encode the image data as a base64 string.
	//imageBase64 := base64.StdEncoding.EncodeToString(imageData)

	return imageData, nil
}
