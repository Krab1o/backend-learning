package jsonstruct

import "task-tracker/internal/task"

type Data struct {
	ID 			uint32		`json:"ID"`
	Tasks		[]task.Task `json:"tasks"`	
}