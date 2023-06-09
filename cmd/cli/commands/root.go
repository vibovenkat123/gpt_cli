package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/vibovenkat123/gpt_cli/pkg/helpers"
)

var key string
var max_tokens int
var verbose bool
var rootCmd = &cobra.Command{
	Use:   "gptcli [FLAGS] PROMPT",
	Short: "Gpt CLI is an Ai powered way to ask GPT4 on the CLI",
	Long: `A command-line interface for OpenAI's GPT-4 API.
Required argument:
  PROMPT              The prompt for the API to generate text.

Examples:
  gptcli "Hello, world!"
  gptcli -k sk-abc123 "Hello, world!"
  gptcli --verbose "Hello, world!"
  gptcli -mt 1024 "Hello, world!"
    `,
	Args: cobra.RangeArgs(1, 21),
	Run:  Run,
}

func Run(cmd *cobra.Command, args []string) {
	prompt := args[0]
	helpers.Validate(key)
	helpers.LogVerbose(verbose, "Generating Request Struct")
	reqJSON := helpers.ApiReq{
		Max_tokens: max_tokens,
		Model:      "gpt-4",
		Messages: []helpers.Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}
	helpers.LogVerbose(verbose, "Marshalling JSON")
	jsonParams, err := json.Marshal(reqJSON)
	if err != nil {
		fmt.Println("Error while marshalling JSON", err)
		os.Exit(1)
	}
	helpers.LogVerbose(verbose, "Converting JSON to bytes")
	reqBody := bytes.NewBuffer(jsonParams)
	helpers.LogVerbose(verbose, "Making new request")
	req, err := http.NewRequest("POST", helpers.Url, reqBody)
	if err != nil {
		fmt.Println("Error while making request", err)
		os.Exit(1)
	}
	helpers.LogVerbose(verbose, "Apply headers to request")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", key))
	client := &http.Client{}
	helpers.LogVerbose(verbose, "Executing request")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while sending request", err)
	}
	defer resp.Body.Close()
	helpers.LogVerbose(verbose, "Reading Body")
	body, _ := io.ReadAll(resp.Body)
	apiRes := helpers.ApiRes{}
	helpers.LogVerbose(verbose, "Unmarshalling Body")
	err = json.Unmarshal([]byte(string(body)), &apiRes)
	if err != nil {
		fmt.Println("Error while recieving request body", err)
		os.Exit(1)
	}
	helpers.LogVerbose(verbose, apiRes)
	helpers.LogVerbose(verbose, "Printing message, msg:")
	printRes(apiRes)
}

func Execute() {
	rootCmd.Flags().StringVarP(&key, "key", "k", os.Getenv("OPENAI_KEY"), "The API key to use for authentication.")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose mode for more detailed output.")
	rootCmd.Flags().IntVarP(&max_tokens, "max", "m", 2000, "The maximum number of tokens to generate. (default: 2000)")
	//	helpers.LogVerbose(verbose, "Generating man")
	//	helpers.GenerateMan(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printRes(res helpers.ApiRes) {
	for _, choice := range res.Choices {
		fmt.Println(choice.Message.Content)
	}
}
