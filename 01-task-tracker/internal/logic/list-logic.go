package logic

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	jsonstruct "task-tracker/internal/json_struct"
	"task-tracker/internal/status"
	"text/tabwriter"
)

var taskHeader =
"Number\tTask\tDescription\tStatus\tCreationDate\tUpdateDate\tTags"

var taskTemplate = 
"#%v\t%s\t%s\t%s\t%s\t%s\t%s\n"

var noTasksYetMsg = 
`There is no tasks yet! Create one with "COMMAND EXAMPLE"`

var noTasksToDoMsg =
`There are no tasks! You may go chill.`

var noTasksDoneMsg =
`There are no completed tasks! Are you lazy or you have nothing to do?`

var noTasksInProgressMsg =
`Nothing in progress...`

var noTasks =
`No tasks at all! Let's create a new one...`

func ListTask(statuses []string) {
	
	file, err := os.ReadFile(dataPath)
	if err == os.ErrNotExist {
		fmt.Print(noTasksYetMsg)
	}

	data := jsonstruct.Data{}
	json.Unmarshal(file, &data)

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.Debug)

	var isTask bool
	for _, task := range data.Tasks {
		if (slices.Contains(statuses, task.Status)) {
			if (!isTask) {
				isTask = true
				fmt.Fprintln(w, taskHeader)
			}
			if (task.Description.String == "") {
				task.Description.String = "---"
			}
			if (task.UpdatedDate.String == "") {
				task.UpdatedDate.String = "---"
			}

			fmt.Fprintf(w, 
				taskTemplate,
				task.ID, 
				task.Title, 
				task.Description.String,
				task.Status,
				task.AddedDate,
				task.UpdatedDate.String,
				task.Tags,
			)
		}
	}

	if (!isTask) {
		if (len(statuses) != 1) {
			fmt.Println(noTasks)
		} else {
			switch statuses[0] {
			case status.ToDo.String():
				fmt.Println(noTasksToDoMsg)
			case status.InProgress.String():
				fmt.Println(noTasksInProgressMsg)
			case status.Done.String():
				fmt.Println(noTasksDoneMsg)
			default:
				// Intentionally blank (cannot reach this)
			}
		}
	} else {
		w.Flush()
	}
}