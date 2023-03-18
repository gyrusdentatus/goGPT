package main

import (
    "bufio"
    "context"
    "fmt"
    "os"
    "strings"
    openai "github.com/sashabaranov/go-openai"
    "github.com/fatih/color"
)

var (
    promptStyle       = color.New(color.FgHiMagenta, color.Bold)
    userInputStyle    = color.New(color.FgHiCyan)
    goGPTResponseStyle = color.New(color.FgHiGreen)
    errorStyle        = color.New(color.FgHiRed, color.Bold)
)

func PrintStyled(c *color.Color, format string, a ...interface{}) {
    c.PrintfFunc()(format, a...)
}

func main() {
    client := openai.NewClient("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")

    reader := bufio.NewReader(os.Stdin)

    PrintStyled(promptStyle, "Welcome to GoGPT! Please choose a role: writer, programmer, or teacher\n")
    role, _ := reader.ReadString('\n')
    role = strings.TrimSpace(strings.ToLower(role))

    for {
        PrintStyled(promptStyle, "Enter your message (type 'exit' to quit):\n")
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
            PrintStyled(errorStyle, "ChatCompletion error: %v\n", err)
            return
        }

        PrintStyled(userInputStyle, "You: %s\n", userInput)
        PrintStyled(goGPTResponseStyle, "GoGPT (%s): %s\n", role, resp.Choices[0].Message.Content)
    }
}

