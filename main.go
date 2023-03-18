package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
	"github.com/fatih/color"
	"GoGPT/db"
)

var (
	promptStyle       = color.New(color.FgHiMagenta, color.Bold)
	userInputStyle    = color.New(color.FgHiCyan)
	goGPTResponseStyle = color.New(color.FgHiGreen)
	errorStyle        = color.New(color.FgHiRed, color.Bold)
)

func printStyled(c *color.Color, format string, a ...interface{}) {
	c.PrintfFunc()(format, a...)
}

func main() {
	client := openai.NewClient("your token")

	reader := bufio.NewReader(os.Stdin)

	printStyled(promptStyle, "Welcome to GoGPT! Please choose a role: writer, programmer, or teacher\n")
	role, _ := reader.ReadString('\n')
	role = strings.TrimSpace(strings.ToLower(role))

	// Initialize the database
	db := initDB()
	defer db.Close()

	for {
		printStyled(promptStyle, "Enter your message (type 'exit' to quit):\n")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if userInput == "exit" {
			break
		}

		// Fetch the conversation history
		messages := getConversationHistory(db)

		// Add user message to the history
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: userInput,
		})

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    openai.GPT3Dot5Turbo,
				Messages: messages,
			},
		)

		if err != nil {
			printStyled(errorStyle, "ChatCompletion error: %v\n", err)
			return
		}

		printStyled(userInputStyle, "You: %s\n", userInput)
		printStyled(goGPTResponseStyle, "GoGPT (%s): %s\n", role, resp.Choices[0].Message.Content)

		// Add messages to the database
		addMessage(db, openai.ChatMessageRoleUser, userInput)
		addMessage(db, openai.ChatMessageRoleAssistant, resp.Choices[0].Message.Content)
	}
}

