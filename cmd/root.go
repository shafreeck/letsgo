package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "letsgo",
	Short: "letsgo helps you to code in go more productively",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
