package tag

import (
	"strconv"
	"task-tracker/internal/logic"

	"github.com/spf13/cobra"
)

var TagCmd = &cobra.Command {
    Use:    "tag",
    Short:  "Adds tag to your task",
    Run: func(cmd *cobra.Command, args []string) {
		val, _ := strconv.Atoi(args[0])
        logic.AddTag(uint32(val), args[1])
    },
}

func init() {

}