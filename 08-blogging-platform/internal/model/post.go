package model

import "time"

type Post struct {
	ID			uint		`json:"id"`
	Title		string		`json:"title"`
	Content		string		`json:"content"`
	Category	string		`json:"category"`
	Tags		[]string	`json:"tags"`
	CreatedAt	time.Time	`json:"createdAt"`
	UpdatedAt	time.Time	`json:"updatedAt"`
}

type Tag struct {
	ID 			uint		`json:"id"`
	Title		string		`json:"tag"`
}

type Link struct {
	PostID		uint
	TagID		uint
}