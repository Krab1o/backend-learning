package mark

import (
	category "task-tracker/cmd/mark/categories"

	"github.com/spf13/cobra"
)

var MarkCmd = &cobra.Command {
    Use:    "mark",
    Short:   `Changes your task status`,
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func init() {
    MarkCmd.AddCommand(category.MarkDoneCmd)
	MarkCmd.AddCommand(category.MarkInProgressCmd)
	MarkCmd.AddCommand(category.MarkToDoCmd)
}