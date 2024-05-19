package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"task/manage"
)

func init() {
	rootCmd.AddCommand(list)
}

var list = &cobra.Command{
	Use:   "list",
	Short: "list tasks",
	Long:  `This command prints a greeting message`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks := manage.ListTasks()
		for _, task := range tasks {
			fmt.Printf("%v- %v \n", task.Id, task.Details)
		}
	},
}
