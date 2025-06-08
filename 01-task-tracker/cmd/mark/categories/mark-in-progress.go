package category

import (
	"strconv"
	"task-tracker/internal/logic"
	"task-tracker/internal/status"

	"github.com/spf13/cobra"
)

var MarkInProgressCmd = &cobra.Command {
    Use:    "in-progress",
    Short:  "Marks task as in-progress",
    Long:   
`Marks task with particular ID as "in-progress".`,
    Run: func(cmd *cobra.Command, args []string) {
        val, _ := strconv.Atoi(args[0])
        logic.MarkTask(uint32(val), status.InProgress.String())
    },
}

func init() {
    
}