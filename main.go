package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetResponse(client gpt3.Client, ctx context.Context, quesiton string) {
	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			quesiton,
		},
		MaxTokens:   gpt3.IntPtr(4000), // This cost more money more the token ChatGPT will charge more
		Temperature: gpt3.Float32Ptr(0),
	}, func(resp *gpt3.CompletionResponse) {
		fmt.Print(resp.Choices[0].Text)
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(13)
	}
	fmt.Printf("\n")
}

type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

func GetAPIKeyFromEnv(file string) (key string) {
	viper.SetConfigFile(file)
	viper.ReadInConfig()
	thekey := viper.GetString("API_KEY")
	if thekey == "" {
		panic("Missing API KEY")
	}
	return thekey
}

func main() {
	log.SetOutput(new(NullWriter))
	/* Read API Key from .env file */
	apiKey := GetAPIKeyFromEnv(".env")

	/* Set Context and GPT Client */
	ctx := context.Background()
	gpt3Client := gpt3.NewClient(apiKey)

	/* Read CLI message untill user enters quit */
	cliCommad := &cobra.Command{
		Use:   "chatgpt",
		Short: "Chat with ChatGPT in console.",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false
			for !quit {
				fmt.Print("How can I help you ? ('quit' to end): ")
				if !scanner.Scan() {
					break
				}
				question := scanner.Text()
				switch question {
				case "quit":
					quit = true
				default:
					GetResponse(gpt3Client, ctx, question)
					fmt.Println("/n")
				}
			}
		},
	}
	cliCommad.Execute()
}
