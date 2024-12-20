package category

import (
	"task-tracker/internal/logic"
	"task-tracker/internal/status"

	"github.com/spf13/cobra"
)

var ListDoneCmd = &cobra.Command {
    Use:    "done",
    Short:  "Shows all done tasks",
    Long:   
`Shows all tasks with "done" group.`,
    Run: func(cmd *cobra.Command, args []string) {
        logic.ListTask([]string{status.Done.String()})
    },
}

func init() {
    
}