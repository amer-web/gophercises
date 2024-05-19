package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "manage",
	Short: "manage is a CLI for managing your TODOs.",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from MyCLI!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
