package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
	"task/manage"
)

func init() {
	rootCmd.AddCommand(add)
}

var add = &cobra.Command{
	Use:   "add",
	Short: "add manage",
	Long:  `This command prints a greeting message`,
	Run: func(cmd *cobra.Command, args []string) {
		if ar := len(args); ar == 0 {
			log.Fatal("should add manage")
		}
		details := strings.Join(args, " ")
		task := manage.Task{Details: details, Status: false}
		if task.CreateTask(&task) != nil {
			log.Fatal("manage not add yet")
		}
		fmt.Println("task added")
	},
}
