package types

import "time"

type User struct {
	Email 		string	`json:"email"`
	Password 	string	`json:"password"`
}

type Article struct {
	ID				uint		`json:"id"`
	Title			string		`json:"title"`
	CreatingTime	time.Time	`json:"creating_time"`
	EditingTime		time.Time	`json:"editing_time"`
	Content			string		`json:"content"`
}

type ArticleStorage []Article