package logic

import (
	"encoding/json"
	"log"
	"os"
	jsonstruct "task-tracker/internal/json_struct"
	"time"

	"github.com/guregu/null"
)

func AddTag (ID uint32, tag string) {
	file, err := os.ReadFile(dataPath)

	if err != nil {
		log.Fatal(err)
	}

	data := jsonstruct.Data{}
	json.Unmarshal(file, &data)

	for i := range data.Tasks {
		if (data.Tasks[i].ID == ID) {
			data.Tasks[i].Tags = append(data.Tasks[i].Tags, tag)
			data.Tasks[i].UpdatedDate = null.StringFrom(time.Now().Format(time.DateTime))
		}
	}

	dataBinary, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	writeFile(dataBinary)
}