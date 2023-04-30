package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)
var key string
var rootCmd = &cobra.Command{
	Use:   "gptcli PROMPT [-k|key] key",
	Short: "Gpt CLI is an Ai powered way to ask GPT4 on the CLI",
	Long: `
An AI powered CLI to ask GPT4 any type of question, ranging from CLI to everyday questions.
    `,
	Args: cobra.ExactArgs(1),
	Run:  Run,
}

func Run(cmd *cobra.Command, args []string) {
	prompt := args[0];
	if len(key) == 0 {
		fmt.Println("The `key` flag or `OPENAI_KEY` is not set")
		os.Exit(1)
	}
	if !strings.HasPrefix(key, "sk-") {
		fmt.Println("The api key isn't in the right format, (must start with `sk-`")
		os.Exit(1)
	}
	fmt.Println(prompt, key)
}

func Execute() {
	rootCmd.Flags().StringVarP(&key, "key", "k", os.Getenv("OPENAI_KEY"), "The api key for openapi")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
