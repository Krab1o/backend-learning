package cmd

import (
	"fmt"
	"os"
	"task-tracker/cmd/add"
	"task-tracker/cmd/delete"
	"task-tracker/cmd/list"
	"task-tracker/cmd/mark"
	"task-tracker/cmd/tag"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
    Use:   "task-tasker",
    Short: "Task-tracker is a cli tool for tracking your tasks",
    Long:  "Task-tracker is a cli tool for tracking your tasks. some more blah-blah",
    Run: func(cmd *cobra.Command, args []string) {
		
    },
}

func init() {
	rootCmd.AddCommand(add.AddCmd)
	rootCmd.AddCommand(list.ListCmd)
    rootCmd.AddCommand(mark.MarkCmd)
    rootCmd.AddCommand(delete.DeleteCmd)
    rootCmd.AddCommand(tag.TagCmd)
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Oops. An error while executing task-tracker '%s'\n", err)
        os.Exit(1)
    }
}