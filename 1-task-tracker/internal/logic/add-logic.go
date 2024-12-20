package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	jsonstruct "task-tracker/internal/json_struct"
	"task-tracker/internal/status"
	"task-tracker/internal/task"
	"time"
)

func AddTask(taskName string) {
	if _, err := os.Stat(directoryPath); errors.Is(err, os.ErrNotExist) {
		createDirectory()
	}

	file, err := os.ReadFile(dataPath)
	switch {
	case errors.Is(err, os.ErrNotExist):
	if _, createErr := os.Create(dataPath); createErr != nil {
		log.Fatal(createErr)
	}
	case err != nil:
		log.Fatal(err)
	default:
		// Left blank intentionally
	}

	data := jsonstruct.Data{}
	json.Unmarshal(file, &data)

	ID := data.ID + 1 
	
	currentTime := time.Now().Format(time.DateTime)
	addedTask := &task.Task {
		ID: 		ID,
		Title: 		taskName,
		Status: 	status.ToDo.String(),
		AddedDate: 	currentTime,
	}

	data.Tasks = append(data.Tasks, *addedTask)
	data.ID = ID

	dataBinary, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	
	writeFile(dataBinary)	
	fmt.Printf("Task %s added successfully!\n", taskName)
}