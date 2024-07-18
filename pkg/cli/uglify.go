package cli

import (
	"fmt"
	"os"

	"benodiwal/github.com/go-shit/pkg/utils"

	"github.com/spf13/cobra"
)

var inputFile string

var uglifyCmd = &cobra.Command{
	Use: "uglify",
	Short: "Uglifies your code",
	Long: `Uglifies your code by removing formatting and adding unnecessary spaces/characters.`,
	Run: func(cmd *cobra.Command, args []string) {
		if inputFile == "" {
			fmt.Println("Input File is required")
			os.Exit(1)
		} else {
			fmt.Println(string(utils.ReadFile(inputFile)))
		}
	},
}
