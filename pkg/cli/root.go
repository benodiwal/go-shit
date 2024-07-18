package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(uglifyCmd)
	
	uglifyCmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "", "Input file to uglify")
	uglifyCmd.MarkFlagRequired("input")

	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

var rootCmd = &cobra.Command {
	Use: "go-shit",
	Short: "go-shit makes your code ugly",
	Long: `go-shit is a tool that "uglifies" your code by removing formatting, adding unnecessary spaces/characters, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
