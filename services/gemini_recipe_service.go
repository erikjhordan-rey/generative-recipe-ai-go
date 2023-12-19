package services

import (
	"context"
	"log"
	"os"

	//"path/filepath"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func GenerateRecipe(data []byte) genai.Part {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error reading .env file")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro-vision")

	//data, err := os.ReadFile(filepath.Join("assets", "salad.jpeg"))
	//if err != nil {
	//	log.Fatal("error reading image: ", err)
	//	}

	prompt := genai.Text("Provide a list of attributes: name, description, ingredients, type of cuisine, vegetarian or not as string type, in file.json")
	imageData := genai.ImageData("jpeg", data)
	resp, err := model.GenerateContent(ctx, prompt, imageData)
	if err != nil {
		log.Fatal("error generating content: ", err)
	}

	return getResponse(resp)

}

func getResponse(resp *genai.GenerateContentResponse) genai.Part {
	for _, cand := range resp.Candidates {
		for _, part := range cand.Content.Parts {
			return part
		}
	}
	return nil
}
