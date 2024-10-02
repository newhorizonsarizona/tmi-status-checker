package util

import (
	"context"
	"log"
	"os"
	"time"

	openai "github.com/sashabaranov/go-openai"
	"golang.org/x/time/rate"
)

func Test() {
	answer := Chat(`
				In one concise to the point paragraph with a formal encouraging tone highlight the club achievements. 
				The first part praises the club for their overall achievements and membership, and 
				the second part commends the club for Distinguished Club Program goals achieved in each category. 
Distinguished Club Program Report:
    Administration:
        Club officer list on time:
            achieved: "1"
            status: ""
            target: "Y"
        Membership-renewal dues on time:
            achieved: "1"
            status: Achieved
            target: "Y"
    DCP Status:
        Membership:
            Base: "22"
            Required: "20"
            To Date: "26"
        Overall:
            Current: "4"
            Distinguished: "No"
            Target: "10"
            Year: 2024-2025
    Education:
        Level 1 awards:
            achieved: "2"
            status: 2 Level 1s needed
            target: "4"
        Level 2 awards:
            achieved: "0"
            status: 2 Level 2s needed
            target: "2"
        Level 3 awards:
            achieved: "0"
            status: 2 Level 3s needed
            target: "2"
        Level 4; Level 5; or DTM award:
            achieved: "1"
            status: Achieved
            target: "1"
        More Level 2 awards:
            achieved: "0"
            status: 2 Level 2s needed
            target: "2"
        One more Level 4; Level 5; or DTM award:
            achieved: "1"
            status: Achieved
            target: "1"
    Membership:
        More new members:
            achieved: "0"
            status: 4 New Members needed
            target: "4"
        New members:
            achieved: "4"
            status: Achieved
            target: "4"
    Training:
        Club officers trained June-August:
            achieved: "4"
            status: First Training Period Achieved
            target: "4"
        Club officers trained November-February:
            achieved: "0"
            status: Second Training Period 4 needed
            target: "4"

				`)
	log.Println("Answer: ", answer)
}

func Chat(question string) string {
	api_key := os.Getenv("OPENAI_API_KEY")
	if api_key == "" {
		log.Println("API Key error.")
		return "API Key error"
	}
	// Initialize the rate limiter (e.g., 1 request per second)
	limiter := rate.NewLimiter(rate.Every(1*time.Second), 1)
	client := openai.NewClient(api_key)
	// Define your request
	request := openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: question,
			},
		},
	}

	answer, err := executeRequestWithThrottle(client, request, limiter)

	if err != nil {
		log.Printf("Chat Completion error: %v\n", err)
		return "Chat Completion error"
	}

	return answer

}

func executeRequestWithThrottle(client *openai.Client, request openai.ChatCompletionRequest, limiter *rate.Limiter) (string, error) {
	// Wait for the rate limiter to allow a request
	if err := limiter.Wait(context.Background()); err != nil {
		return "rate limit wait error", err
	}

	// Perform the request
	response, err := client.CreateChatCompletion(context.Background(), request)
	if err != nil {
		return "request error", err
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "None", nil
}
