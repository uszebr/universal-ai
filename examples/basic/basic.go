package main

import (
	"context"
	"fmt"

	"github.com/uszebr/universal-ai/base"
	"github.com/uszebr/universal-ai/util"
)

func main() {

	// Same request is used for Gemini, OpenAI
	req := base.AIRequest{
		Messages: []base.Message{
			{Role: "user", Content: util.StrPtr("What is the population of Europe?")},
		},
		N: util.Ptr(2),
	}

	service := base.NewAIService("OpenAiAPIKey", "https://api.openai.com/v1/chat/completions", "gpt-4.1")

	response, err := service.Request(context.Background(), req)
	if err != nil {
		panic(err)
	}

	if len(response.Choices) == 0 {
		fmt.Println("No choices returned")
		return
	}
	fmt.Printf("==== OPENAI Basic ===============\n")
	fmt.Printf("Quantity of choices: %v\n", len(response.Choices))
	fmt.Printf("Content 0: %v\n", *response.Choices[0].Message.Content)
	fmt.Printf("Content 1: %v\n", *response.Choices[1].Message.Content)

	serviceGemini := base.NewAIService("GeminiAPIKey", "https://generativelanguage.googleapis.com/v1beta:chatCompletions", "gemini-2.5-flash-preview-04-17")

	responseGemini, err := serviceGemini.Request(context.Background(), req)
	if err != nil {
		panic(err)
	}
	if len(responseGemini.Choices) == 0 {
		fmt.Println("No choices returned")
		return
	}
	fmt.Printf("==== Gemini Basic ===============\n")
	fmt.Printf("Quantity of choices: %v\n", len(responseGemini.Choices))
	fmt.Printf("Content 0: %v\n", *responseGemini.Choices[0].Message.Content)
	fmt.Printf("Content 1: %v\n", *responseGemini.Choices[1].Message.Content)

	serviceDeepSeek := base.NewAIService("DeepSeekAPIKey", "https://api.deepseek.com/chat/completions", "deepseek-chat")

	reqDeepSeek := base.AIRequest{
		Messages: []base.Message{
			{Role: "user", Content: util.StrPtr("What is the population of Europe?")},
		},
		N: util.Ptr(1), // only 1 is allowed for deepsek(or omit)
	}
	responseDeepSeek, err := serviceDeepSeek.Request(context.Background(), reqDeepSeek)
	if err != nil {
		panic(err)
	}
	if len(responseDeepSeek.Choices) == 0 {
		fmt.Println("No choices returned")
		return
	}
	fmt.Printf("==== DeepSeek Basic ===============\n")
	fmt.Printf("Quantity of choices: %v\n", len(responseDeepSeek.Choices))
	fmt.Printf("Content 0: %v\n", *responseDeepSeek.Choices[0].Message.Content)

}
