package helpers

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func LogVerbose[T any](verbose bool, msg T) {
	if verbose {
		fmt.Println(msg)
	}
}

func Validate(key string) {
	if len(key) == 0 {
		fmt.Println("The `key` flag or `OPENAI_KEY` is not set")
		os.Exit(1)
	}
	if !strings.HasPrefix(key, "sk-") {
		fmt.Println("The api key isn't in the right format, (must start with `sk-`")
		os.Exit(1)
	}
}

func GenerateMan(cmd *cobra.Command) {
	header := &doc.GenManHeader{
		Title:   "MINE",
		Section: "1",
	}
	err := doc.GenManTree(cmd, header, "./man")
	if err != nil {
		fmt.Println("Error generating man", err)
	}
}
