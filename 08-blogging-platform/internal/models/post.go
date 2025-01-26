package models

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

type PostStorage interface {
	Add(newPost Post)
	Get()
	GetByID(id uint)
	Delete()
	Update()
}