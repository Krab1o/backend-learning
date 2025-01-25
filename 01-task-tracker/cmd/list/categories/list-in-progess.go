package category

import (
	"task-tracker/internal/logic"
	"task-tracker/internal/status"

	"github.com/spf13/cobra"
)

var ListInProgressCmd = &cobra.Command {
    Use:    "in-progress",
    Short:  "Shows all in-progress tasks",
    Long:   
`Shows all tasks with "in-progress" group.`,
    Run: func(cmd *cobra.Command, args []string) {
        logic.ListTask([]string{status.InProgress.String()})
    },
}

func init() {
    
}