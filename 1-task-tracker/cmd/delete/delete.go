package delete

import (
	"strconv"
	"task-tracker/internal/logic"

	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command {
    Use:    "delete",
    Short:  "Deletes a task",
    Long:   
`Deletes a task from the list of your tasks. 
Deletes it by id.`,
    Run: func(cmd *cobra.Command, args []string) {
        val, _ := strconv.Atoi(args[0])
		logic.DeleteTask(uint32(val))
    },
}

func init() {
    
}