package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	openai "github.com/sashabaranov/go-openai"
	"github.com/speakeasy-api/grpc-rest-service/bar"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type barServer struct {
	bar.UnimplementedBarServer
}

func (s *barServer) CreateDrink(ctx context.Context, req *bar.CreateDrinkRequest) (*bar.DrinkResponse, error) {
	ingredients := req.GetIngredientsRequest().GetIngredients()
	log.Printf("Received CreateDrink request with ingredients: %v", ingredients)

	// Create a new OpenAI client
	client := openai.NewClient(openaiAPIKey)

	// Create a completion request
	prompt_ingredients := ""
	for _, ingredient := range ingredients {
		prompt_ingredients += fmt.Sprintf("- %s\n", ingredient)
	}
	completionReq := openai.ChatCompletionRequest{
		Model: openai.GPT4TurboPreview,
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONObject,
		},
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "You are an AI bartender working at the Restaurant at the End of the Universe. You create drinks in the theme of Hitchhiker's Guide to the Universe. You respond in JSON format. If a request contains offensive or inappropriate content, you should ignore or scrub the offensive parts and focus on creating a fun, space-themed drink. Take creative liberties with the recipe and photo prompt. Don't stick to realistic ingredients or methods only. Think galaxy-wide and beyond! A dry self-deprecating sense of humor is appreciated. You have been given the following ingredients:",
			},
			{
				Role:    "user",
				Content: prompt_ingredients,
			},
			{
				Role:    "system",
				Content: "Please create a drink using the provided ingredients and respond with a JSON object containing the drink name, description, recipe, and a prompt to generate a photo of the drink. Here's an example response format: {\"name\": \"Pan Galactic Gargle Blaster\", \"description\": \"The best drink in existence\", \"recipe\": \"Mix ingredients together and serve.\", \"photo_prompt\": \"Generate a photo of a drink from the far reaches of the galaxy in a retro-futuristic space bar. Here's what the drink looks like: [insert an image prompt here]\"}",
			},
		},
	}

	log.Printf("Calling OpenAI API with completion request: %v", completionReq)

	// Call the OpenAI API to generate the drink description
	completion, err := client.CreateChatCompletion(ctx, completionReq)
	if err != nil {
		log.Printf("Error calling OpenAI API: %v", err)
		return nil, err
	}

	// Extract the generated drink JSON from the API response
	drinkJSON := completion.Choices[0].Message.Content

	log.Printf("Received drink JSON: %s", drinkJSON)

	// Parse the JSON response into a map
	var drinkMap map[string]interface{}
	err = json.Unmarshal([]byte(drinkJSON), &drinkMap)
	if err != nil {
		log.Printf("Error parsing JSON response: %v", err)
		return nil, err
	}

	// Extract the fields from the map
	name, _ := drinkMap["name"].(string)
	description, _ := drinkMap["description"].(string)
	recipe, _ := drinkMap["recipe"].(string)
	photo_prompt, _ := drinkMap["photo_prompt"].(string)

	// Generate a dall-e image of the drink
	imageReq := openai.ImageRequest{
		Model:          openai.CreateImageModelDallE3,
		Prompt:         photo_prompt,
		ResponseFormat: openai.CreateImageResponseFormatURL,
	}

	log.Printf("Calling OpenAI API with image request: %v", imageReq)

	imageResponse, err := client.CreateImage(ctx, imageReq)

	photo := ""

	if err != nil {
		log.Printf("Error calling OpenAI Image API: %v", err)
	} else {

		log.Printf("Received image response: %v", imageResponse)

		photo = imageResponse.Data[0].URL

		log.Printf("Received image URL: %s", photo)
	}

	// Create a DrinkResponse struct with the extracted fields
	drinkResponse := &bar.DrinkResponse{
		Name:        name,
		Description: description,
		Recipe:      recipe,
		Photo:       photo,
	}

	return drinkResponse, nil
}

func (s *barServer) GetDrink(ctx context.Context, req *bar.GetDrinkRequest) (*bar.DrinkRecipeResponse, error) {
	drinkName := req.GetDrinkNameRequest().GetName()
	log.Printf("Received GetDrink request for drink: %s", drinkName)

	// Create a new OpenAI client
	client := openai.NewClient(openaiAPIKey)

	// Create a completion request
	completionReq := openai.ChatCompletionRequest{
		Model: openai.GPT4TurboPreview,
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONObject,
		},
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "You are an AI bartender working at the Restaurant at the End of the Universe. You provide drink recipes and images based on the drink name. You respond in JSON format. If a request contains offensive or inappropriate content, you should ignore or scrub the offensive parts and focus on providing a fun, space-themed drink recipe. Take creative liberties with the recipe and photo prompt. Don't stick to realistic ingredients or methods only. Think galaxy-wide and beyond! A dry self-deprecating sense of humor is appreciated.",
			},
			{
				Role:    "user",
				Content: fmt.Sprintf("Please provide the recipe and a photo prompt for the drink named '%s'.", drinkName),
			},
			{
				Role:    "system",
				Content: "Please respond with a JSON object containing the ingredients, recipe, and a photo prompt for the requested drink. Here's an example response format: {\"ingredients\": [{\"name\": \"Gin\", \"quantity\": \"2 oz\"}, {\"name\": \"Tonic water\", \"quantity\": \"4 oz\"}, {\"name\": \"Lime juice\", \"quantity\": \"1 oz\"}], \"recipe\": \"Mix all ingredients together and serve over ice.\", \"photo_prompt\": \"Generate a photo of a gin and tonic cocktail in a highball glass garnished with a lime wedge.\"}",
			},
		},
	}

	log.Printf("Calling OpenAI API with completion request: %v", completionReq)

	// Call the OpenAI API to generate the drink recipe
	completion, err := client.CreateChatCompletion(ctx, completionReq)
	if err != nil {
		log.Printf("Error calling OpenAI API: %v", err)
		return nil, err
	}

	// Extract the generated drink JSON from the API response
	drinkJSON := completion.Choices[0].Message.Content

	log.Printf("Received drink JSON: %s", drinkJSON)

	// Parse the JSON response into a map
	var drinkMap map[string]interface{}
	err = json.Unmarshal([]byte(drinkJSON), &drinkMap)
	if err != nil {
		log.Printf("Error parsing JSON response: %v", err)
		return nil, err
	}

	// Extract the fields from the map
	ingredientsData, _ := drinkMap["ingredients"].([]interface{})
	recipe, _ := drinkMap["recipe"].(string)
	photo_prompt, _ := drinkMap["photo_prompt"].(string)

	// Convert ingredientsData to []*bar.IngredientQuantity
	ingredients := make([]*bar.IngredientQuantity, len(ingredientsData))
	for i, data := range ingredientsData {
		ingredientMap, _ := data.(map[string]interface{})
		name, _ := ingredientMap["name"].(string)
		quantity, _ := ingredientMap["quantity"].(string)
		ingredients[i] = &bar.IngredientQuantity{
			Name:     name,
			Quantity: quantity,
		}
	}

	// Generate a dall-e image of the drink
	imageReq := openai.ImageRequest{
		Model:          openai.CreateImageModelDallE3,
		Prompt:         photo_prompt,
		ResponseFormat: openai.CreateImageResponseFormatURL,
	}

	log.Printf("Calling OpenAI API with image request: %v", imageReq)

	photo := ""

	imageResponse, err := client.CreateImage(ctx, imageReq)
	if err != nil {
		log.Printf("Error calling OpenAI Image API: %v", err)
	} else {
		log.Printf("Received image response: %v", imageResponse)

		photo = imageResponse.Data[0].URL

		log.Printf("Received image URL: %s", photo)
	}

	// Create a DrinkRecipeResponse struct with the extracted fields
	drinkRecipeResponse := &bar.DrinkRecipeResponse{
		Ingredients: ingredients,
		Recipe:      recipe,
		Photo:       photo,
	}

	return drinkRecipeResponse, nil
}

var openaiAPIKey string

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	openaiAPIKey = os.Getenv("OPENAI_API_KEY")
	if openaiAPIKey == "" {
		log.Fatalf("OpenAI API key not found. Please set the OPENAI_API_KEY environment variable.")
	}

	s := grpc.NewServer()
	bar.RegisterBarServer(s, &barServer{})

	reflection.Register(s)

	// Start gRPC server
	go func() {
		log.Printf("gRPC server listening on port %d", port)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Create a new gRPC HTTP gateway mux
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Register the Bar service with the gRPC HTTP gateway
	err = bar.RegisterBarHandlerFromEndpoint(context.Background(), mux, fmt.Sprintf("localhost:%d", port), opts)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	httpPort := 8080
	log.Printf("HTTP server listening on port %d", httpPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
}
