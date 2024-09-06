package main

import (
	"context"
	"fmt"
	"os"
	"time"

	openai "github.com/sashabaranov/go-openai"
	"golang.org/x/time/rate"
)

func main() {
	api_key := os.Getenv("OPENAI_API_KEY")
	// Initialize the rate limiter (e.g., 1 request per second)
	limiter := rate.NewLimiter(rate.Every(1*time.Second), 1)
	client := openai.NewClient(api_key)
	// Define your request
	request := openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Hello!",
			},
		},
	}
	err := executeRequestWithThrottle(client, request, limiter)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

}

func executeRequestWithThrottle(client *openai.Client, request openai.ChatCompletionRequest, limiter *rate.Limiter) error {
	// Wait for the rate limiter to allow a request
	if err := limiter.Wait(context.Background()); err != nil {
		return fmt.Errorf("rate limit wait error: %w", err)
	}

	// Perform the request
	response, err := client.CreateChatCompletion(context.Background(), request)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}

	fmt.Println("Response:", response)
	return nil
}
