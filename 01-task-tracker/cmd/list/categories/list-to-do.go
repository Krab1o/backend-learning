package category

import (
	"task-tracker/internal/logic"
	"task-tracker/internal/status"

	"github.com/spf13/cobra"
)

var ListToDoCmd = &cobra.Command {
    Use:    "to-do",
    Short:  "Shows all to-do tasks",
    Long:   
`Shows all tasks with "to-do" group.`,
    Run: func(cmd *cobra.Command, args []string) {
        
        logic.ListTask([]string{status.ToDo.String()})
    },
}

func init() {
    
}