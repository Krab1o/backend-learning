package task

import "github.com/guregu/null"

type Task struct {
	ID			uint32		`json:"id"`
	Title 		string		`json:"task"`
	Description	null.String `json:"description"`
	Status		string		`json:"status"`
	AddedDate 	string		`json:"added_date"`
	UpdatedDate null.String	`json:"updated_date"`
	Tags		[]string	`json:"tags"`
}