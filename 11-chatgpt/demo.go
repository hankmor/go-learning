package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

// go-openai 库：
// https://github.com/sashabaranov/go-openai

func main() {
	var token = os.Getenv("OPENAI_API_KEY")
	if token == "" {
		fmt.Println("Not found an openai key, you should config your OPENAI_KEY to your environment variable first")
		os.Exit(-1)
	}
	client := openai.NewClient(token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
					//Content: "Hello, translate this word to chinese: message",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
