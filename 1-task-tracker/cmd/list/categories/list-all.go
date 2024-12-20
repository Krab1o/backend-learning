package category

import (
	"task-tracker/internal/logic"
	"task-tracker/internal/status"

	"github.com/spf13/cobra"
)

var ListAllCmd = &cobra.Command {
    Use:    "all",
    Short:  "Shows all tasks",
    Long:   
`Shows all tasks from 3 different category.`,
    Run: func(cmd *cobra.Command, args []string) {
		logic.ListTask(
			[]string{
				status.ToDo.String(), 
				status.InProgress.String(), 
				status.Done.String(),
			})
    },
}

func init() {
    
}