package util

import (
	"context"
	"log"
	"os"
	"time"

	openai "github.com/sashabaranov/go-openai"
	"golang.org/x/time/rate"
)

var QuestionBank = map[int]string{
	7: `In two paragraphs with a formal encouraging tone highlight the club achievements last term, July through June. 
		The first paragraph praises the club for their overall achievements last term, and the second
		encourages the club to create a Distinguished Club Success plan and work towards the goals for the next term.
		`,
	8: `In two concise to the point paragraphs with a jovial encouraging tone highlight the club achievements. 
		The first paragraph praises the club for their overall achievements in the first month of the new term, and the second
		encourages the club to work on the goals defined in the Distinguished Club Success plan in the ongoing term.
		`,
	9: `In one concise to the point paragraph with a jovial encouraging tone highlight the club achievements. 
		The first part praises the club for their overall achievements and membership in the first two months, and 
		the second part commends the club for the Distinguished Club Program goals achieved in each category. Highlight 
		the fact if the club has attained an overall distinguished status.
		`,
	10: `In one concise to the point paragraph with a formal encouraging tone highlight the club achievements. 
		The first part praises the club for their overall achievements and membership in the first three months, and 
		the second part commends the club for the Distinguished Club Program goals achieved in each category. Highlight 
		the fact if the club has attained an overall distinguished status.
		`,
	11: `In one concise to the point paragraph with a jovial encouraging tone highlight the club achievements. 
		The first part praises the club for their overall achievements and membership first four months, and 
		the second part commends the club for the Distinguished Club Program goals achieved in each category. Highlight 
		the fact if the club has attained an overall distinguished status.
		`,
	12: `In one concise paragraph with a casual holiday sprit, highlight the club achievements so far. 
		The first part praises the club for their overall achievements and membership, and 
		the second part commends the club for the Distinguished Club Program goals achieved in each category. Highlight 
		the fact if the club has attained an overall distinguished status.
		`,
	1: `In one concise to the point paragraph with a formal encouraging tone highlight the club achievements in the first
		six months of the current term. The first part praises the club for their overall achievements and membership so far, and 
		the second part commends the club for the Distinguished Club Program goals achieved in each category. Highlight 
		the fact if the club has attained an overall distinguished status.
		`,
	2: `In one concise to the point paragraph with an informal encouraging tone highlight the club achievements in the current 
		term, July through January. The first part praises the club for their overall achievements and membership so far, and 
		the second part commends the club for the Distinguished Club Program goals achieved in each category. Highlight 
		the fact if the club has attained an overall distinguished status.
		`,
	3: `In one concise to the point paragraph with an jovial encouraging tone highlight the club achievements over the last
		eigth months. The first part praises the club for their overall achievements and membership so far, and 
		the second part commends the club for the Distinguished Club Program goals achieved in each category. Highlight 
		the fact if the club has attained an overall distinguished status.
		`,
	4: `In one concise to the point paragraph with an formal encouraging tone highlight the club achievements over the last
		nine months. The first part praises the club for their overall achievements and membership so far, and 
		the second part commends the club for the Distinguished Club Program goals achieved in each category. Highlight 
		the fact if the club has attained an overall distinguished status.
		`,
	5: `In one concise to the point paragraph with an formal encouraging tone highlight the club achievements with just 
		two months to go in the current term. The first part praises the club for their overall achievements and membership so far, and 
		the second part commends the club for the Distinguished Club Program goals achieved in each category. Highlight 
		the fact if the club has attained an overall distinguished status.
		`,
	6: `In one concise to the point paragraph with an formal encouraging tone highlight the club achievements with just 
		one month to go in the current term. The first part praises the club for their overall achievements and membership so far, and 
		the second part commends the club for the Distinguished Club Program goals achieved in each category. Highlight 
		the fact if the club has attained an overall distinguished status.
		`,
}

func Test() {
	// Open the YAML file
	yaml, err := os.ReadFile("./reports/dcp_report.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	currentTime := time.Now()

	question := QuestionBank[int(currentTime.Month())] + string(yaml)
	answer := Chat(question)
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
		Model: openai.GPT4Turbo,
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
