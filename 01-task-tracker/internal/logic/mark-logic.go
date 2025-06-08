package logic

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	jsonstruct "task-tracker/internal/json_struct"
	"task-tracker/internal/status"
	"time"

	"github.com/guregu/null"
)

var markedToDo =
"Task number %d now needs to be done!\n"

var markedInProgress =
"Task number %d is in progress now!\n"

var markedDone = 
"Task number %d done successfully!\n"

var markedError =
"There is no task with %d ID!\n"

func MarkTask(ID uint32, markStatus string) {
	file, err := os.ReadFile(dataPath)

	if err != nil {
		log.Fatal(err)
	}

	data := jsonstruct.Data{}
	json.Unmarshal(file, &data)

	for i := range data.Tasks {
		if (data.Tasks[i].ID == ID) {
			data.Tasks[i].Status = markStatus
			data.Tasks[i].UpdatedDate = null.StringFrom(time.Now().Format(time.DateTime))
		}
	}

	dataBinary, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	
	writeFile(dataBinary)
	switch markStatus {
	case status.ToDo.String():
		fmt.Printf(markedToDo, ID)
	case status.InProgress.String():
		fmt.Printf(markedInProgress, ID)
	case status.Done.String():
		fmt.Printf(markedDone, ID)
	default:
		log.Print(markedError, ID)
	}
}