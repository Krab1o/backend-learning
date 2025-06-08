package category

import (
	"strconv"
	"task-tracker/internal/logic"
	"task-tracker/internal/status"

	"github.com/spf13/cobra"
)

var MarkToDoCmd = &cobra.Command {
    Use:    "to-do",
    Short:  "Marks task as to-do",
    Long:   
`Marks task with particular ID as "to-do".`,
    Run: func(cmd *cobra.Command, args []string) {
		val, _ := strconv.Atoi(args[0])
        logic.MarkTask(uint32(val), status.ToDo.String())
    },
}

func init() {
    
}