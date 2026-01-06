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
	7: `In two concise sets of bulleted points in a professional encouraging tone highlight the club achievements last year. 
		The first section with a heading praises the club for their overall achievements and membership from July through June. The  
		second section with a heading talks about the the goals that need to be set for the toastmaster year. Add a reminders 
		to complete the club success plan by Sep 30 and the Smedley Award for adding five new, dual, or reinstated members 
		with a join date between Aug 1 and Sep 30. Add a link to the document about distinguished club
		program https://content.toastmasters.org/image/upload/1111-distinguished-club-program.pdf.
		`,
	8: `In two concise sets of bulleted points in a professional encouraging tone highlight the club achievements. 
		The first section with a heading praises the club for their overall achievements in the month of July. The  
		second section with a heading talks about the the goals that need to be set for the toastmaster year. Add a reminders 
		to complete the club success plan by Sep 30 and the Smedley Award for adding five new, dual, or reinstated members 
		with a join date between Aug 1 and Sep 30. Add a link to the document about distinguished club
		program https://content.toastmasters.org/image/upload/1111-distinguished-club-program.pdf.
		`,
	9: `In two concise sets of bulleted points in a professional encouraging tone highlight the club achievements. 
		The first section with a heading praises the club for their overall achievements from Jul to Aug. The second
		section with a heading talks about the the goals that need to be completed from Sep to Jun, leave out the June-Aug training. 
		Add a reminders to complete the club success plan by Sep 30 and the Smedley Award for adding five new, dual, or 
		reinstated members with a join date between Aug 1 and Sep 30. Add a link to the document about distinguished club
		program https://content.toastmasters.org/image/upload/1111-distinguished-club-program.pdf.
		`,
	10: `In two concise sets of bulleted points in a professional encouraging tone highlight the club achievements. 
		The first section with a heading praises the club for their overall achievements and membership from Jul to Sep. The second 
		section with a heading talks about the goals that need to be completed from Oct to Jun, leave out the June-Aug training. 
		Highlight the fact if the club has attained an overall distinguished status. Add a link to the document about distinguished club
		program https://content.toastmasters.org/image/upload/1111-distinguished-club-program.pdf.
		`,
	11: `In two concise sets of bulleted points in a joyful yet professional encouraging tone highlight the club achievements. 
		The first section with a heading praises the club for their overall achievements and membership from Jul to Oct. The second 
		section with a heading talks about the goals that need to be completed from Nov to Jun, leave out the June-Aug training. 
		Highlight the fact if the club has attained an overall distinguished status. Add a link to the document about distinguished club
		program https://content.toastmasters.org/image/upload/1111-distinguished-club-program.pdf.
		`,
	12: `In two concise sets of bulleted points in a joyful yet professional encouraging tone highlight the club achievements. 
		The first section with a heading praises the club for their overall achievements and membership from Jul to Nov. The second 
		section with a heading talks about the goals that need to be completed from Dec to Jun, leave out the June-Aug training. 
		Highlight the fact if the club has attained an overall distinguished status. Add a link to the document about distinguished club
		program https://content.toastmasters.org/image/upload/1111-distinguished-club-program.pdf. 
		`,
	1: `In two concise sets of bulleted points in a professional encouraging tone highlight the club achievements. 
		The first section with a heading praises the club for their overall achievements and membership from Jul to Dec. The second 
		section with a heading talks about the goals that need to be completed from Jan to Jun, leave out the June-Aug training. 
		Add a reminder of the Talk Up Toastmasters member drive for adding five new, dual, or reinstated members with a join date between 
		Feb 1 and Mar 31. Highlight the fact if the club has attained an overall distinguished status. Add a link to the document about 
		distinguished club program https://content.toastmasters.org/image/upload/1111-distinguished-club-program.pdf. 
		`,
	2: `In two concise sets of bulleted points in a professional encouraging tone highlight the club achievements. 
		The first section with a heading praises the club for their overall achievements and membership from Jul to Jan. The second 
		section with a heading talks about the goals that need to be completed from Feb to Jun, leave out the June-Aug training. 
		Add a reminder of the Talk Up Toastmasters member drive for adding five new, dual, or reinstated members with a join date between 
		Feb 1 and Mar 31. Highlight the fact if the club has attained an overall distinguished status. Add a link to the document about 
		distinguished club program https://content.toastmasters.org/image/upload/1111-distinguished-club-program.pdf. 
		`,
	3: `In two concise sets of bulleted points in a professional encouraging tone highlight the club achievements. 
		The first section with a heading praises the club for their overall achievements and membership from Jul to Feb. The second 
		section with a heading talks about the goals that need to be completed from Mar to Jun, leave out the June-Aug and Nov-Feb trainings. 
		Add a reminder of the Talk Up Toastmasters member drive for adding five new, dual, or reinstated members with a join date between 
		Feb 1 and Mar 31. Highlight the fact if the club has attained an overall distinguished status. Add a link to the document about 
		distinguished club program https://content.toastmasters.org/image/upload/1111-distinguished-club-program.pdf. 
		`,
	4: `In two concise sets of bulleted points in a professional encouraging tone highlight the club achievements. 
		The first section with a heading praises the club for their overall achievements and membership from Jul to Mar. The second 
		section with a heading talks about the goals that need to be completed from Apr to Jun, leave out the June-Aug and Nov-Feb trainings. 
		Add a reminder of the Beat the Clock member drive for adding five new, dual, or reinstated members with a join date between 
		May 1 and Jun 30. Highlight the fact if the club has attained an overall distinguished status. Add a link to the document about 
		distinguished club program https://content.toastmasters.org/image/upload/1111-distinguished-club-program.pdf. 
		`,
	5: `In two concise sets of bulleted points in a professional encouraging tone highlight the club achievements. 
		The first section with a heading praises the club for their overall achievements and membership from Jul to Apr. The second 
		section with a heading talks about the goals that need to be completed from May to Jun, leave out the June-Aug and Nov-Feb trainings. 
		Add a reminder of the Beat the Clock member drive for adding five new, dual, or reinstated members with a join date between 
		May 1 and Jun 30. Highlight the fact if the club has attained an overall distinguished status. Add a link to the document about 
		distinguished club program https://content.toastmasters.org/image/upload/1111-distinguished-club-program.pdf. 
		`,
	6: `In two concise sets of bulleted points in a professional encouraging tone highlight the club achievements. 
		The first section with a heading praises the club for their overall achievements and membership from Jul to May. The second 
		section with a heading talks about the goals that need to be completed in Jun, leave out the Nov-Feb training. 
		Add a reminder of the Beat the Clock member drive for adding five new, dual, or reinstated members with a join date between 
		May 1 and Jun 30. Highlight the fact if the club has attained an overall distinguished status. Add a link to the document about 
		distinguished club program https://content.toastmasters.org/image/upload/1111-distinguished-club-program.pdf. 
		`,
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
