package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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
	fmt.Println(prompt)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
