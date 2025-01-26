package repository

import (
	"blogging-platform/internal/db"
	"blogging-platform/internal/model"
	"database/sql"
	"log"
)

const insertStmt = `INSERT INTO posts (title, content, category, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5)`

const selectStmt = `SELECT * FROM posts`
const selectOneStmt = `SELECT * FROM posts WHERE id = $1`

const deleteStmt = `DELETE FROM posts WHERE id = $1`

const updateStmt = `UPDATE posts SET 
	title = $2,
	content = $3,
	category = $4,
	created_at = $5,
	updated_at = $6
WHERE id = $1
`

type PostRepositoryPostgres struct{}

type PostRepository interface {
	CreatePost(newPost *model.Post) error
	GetPosts() ([]model.Post, error)
	GetPostByID(id uint) (*model.Post, error)
	DeletePost(id uint) error
	UpdatePost(id uint, updatedPost *model.Post) error
}

func NewPostRepositoryPostgres() PostRepository {
	return &PostRepositoryPostgres{}
}
func (s *PostRepositoryPostgres) CreatePost(newPost *model.Post) error {	
	_, err := db.DB.Exec(insertStmt,
		newPost.Title,
		newPost.Content,
		newPost.Category,
		newPost.CreatedAt,
		newPost.UpdatedAt,
	)
	if err != nil {
		log.Printf("Store: %v", err)
	} else {
		log.Println("Store: inserted successfully")
	}
	return nil
}

func (s *PostRepositoryPostgres) GetPosts() ([]model.Post, error) {
	rows, err := db.DB.Query(selectStmt)
	if err != nil {
		log.Printf("Store: %v", err)
	} else {
		log.Println("Store: selected successfully")
	}
	defer rows.Close()

	var posts []model.Post
	var post model.Post

	for rows.Next() {
		err := rows.Scan(
			&post.ID, 
			&post.Title, 
			&post.Content, 
			&post.Category, 
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if (err != nil) {
			log.Printf("Store: %v", err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *PostRepositoryPostgres) GetPostByID(id uint) (*model.Post, error) {
	post := &model.Post{}
	err := db.DB.QueryRow(selectOneStmt, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Category,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		log.Printf("Store: nothing was found")
	} else if err != nil{
		log.Printf("Store: failed to get post, %v", err)
	} else {
		log.Println("Store: selected by id successfully")
	}
	return post, err
}

func (s *PostRepositoryPostgres) DeletePost(id uint) error {
	_, err := db.DB.Exec(deleteStmt, id)
	if err != nil {
		log.Printf("Store: %v", err)
	} else {
		log.Println("Store: deleted successfully")
	}
	return err
}

func (s *PostRepositoryPostgres) UpdatePost(id uint, updatedPost *model.Post) error {
	_, err := db.DB.Exec(updateStmt, 
		id,
		updatedPost.Title,
		updatedPost.Content,
		updatedPost.Category,
		updatedPost.CreatedAt,
		updatedPost.UpdatedAt,
	)
	if err != nil {
		log.Printf("Store: %v", err)
	} else {
		log.Println("Store: updated successfully")
	}
	return err
}

