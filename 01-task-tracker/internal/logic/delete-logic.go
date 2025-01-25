package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	jsonstruct "task-tracker/internal/json_struct"
	"task-tracker/internal/task"
)

func remove(s []task.Task, i int) []task.Task {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func DeleteTask(ID uint32) {
	file, err := os.ReadFile(dataPath)

	switch {
	case errors.Is(err, os.ErrNotExist):
		log.Println("No tasks to delete!")
	case err != nil:
		log.Fatal(err)
	default:
		// Left blank intentionally
	}

	data := jsonstruct.Data{}
	json.Unmarshal(file, &data)

	var deleteID uint32
	
	for i := 0; i < len(data.Tasks); i++ {
		if (data.Tasks[i].ID == ID) {
			deleteID = data.Tasks[i].ID
			data.Tasks = remove(data.Tasks, i)
		}
	}

	if (deleteID == 0) {
		fmt.Println("There is no task with this ID!")
	} else {
		dataBinary, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		
		writeFile(dataBinary)	
		fmt.Printf("Task number %d deleted successfully!\n", deleteID)
	}
}