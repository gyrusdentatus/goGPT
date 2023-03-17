package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	client := openai.NewClient("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")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to GoGPT! Please choose a role: writer, programmer, or teacher")
	role, _ := reader.ReadString('\n')
	role = strings.TrimSpace(strings.ToLower(role))

	for {
		fmt.Println("Enter your message (type 'exit' to quit):")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if userInput == "exit" {
			break
		}

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleSystem,
						Content: fmt.Sprintf("You are a %s now. I am here to assist you with any questions you have.", role),
					},
					{
						Role:    openai.ChatMessageRoleUser,
						Content: userInput,
					},
				},
			},
		)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			return
		}

		fmt.Printf("GoGPT (%s): %s\n", role, resp.Choices[0].Message.Content)
	}
}

