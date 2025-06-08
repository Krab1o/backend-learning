package category

import (
	"strconv"
	"task-tracker/internal/logic"
	"task-tracker/internal/status"

	"github.com/spf13/cobra"
)

var MarkDoneCmd = &cobra.Command {
    Use:    "done",
    Short:  "Marks task as done",
    Long:   
`Marks task with particular ID as "done".`,
    Run: func(cmd *cobra.Command, args []string) {
		val, _ := strconv.Atoi(args[0])
        logic.MarkTask(uint32(val), status.Done.String())
    },
}

func init() {
    
}