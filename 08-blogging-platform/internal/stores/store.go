package stores

import (
	"blogging-platform/internal/db"
	"blogging-platform/internal/models"
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

type PostStorageDB struct{}

type PostStorage interface {
	Add(newPost *models.Post)
	Get() []models.Post
	GetByID(id uint) (models.Post, bool)
	Delete(id uint) bool
	Update(id uint, updatedPost *models.Post) bool
}

func NewPostStorageDB() *PostStorageDB {
	return &PostStorageDB{}
}

func (s *PostStorageDB) Add(newPost *models.Post) {	
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
}

func (s *PostStorageDB) Get() []models.Post {
	rows, err := db.DB.Query(selectStmt)
	if err != nil {
		log.Printf("Store: %v", err)
	} else {
		log.Println("Store: selected successfully")
	}
	defer rows.Close()

	var posts []models.Post
	var post models.Post

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
	return posts
}

//TODO: condition on failed search
func (s *PostStorageDB) GetByID(id uint) (models.Post, bool) {
	post := models.Post{}
	err := db.DB.QueryRow(selectOneStmt, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Category,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		log.Printf("Store: %v", err)
	} else {
		log.Println("Store: selected by id successfully")
	}
	return post, true
}

//TODO: condition on failed search
func (s *PostStorageDB) Delete(id uint) bool {
	_, err := db.DB.Exec(deleteStmt, id)
	if err != nil {
		log.Printf("Store: %v", err)
	} else {
		log.Println("Store: deleted successfully")
	}
	return true
}

//TODO: condition on failed search
func (s *PostStorageDB) Update(id uint, updatedPost *models.Post) bool {
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
	return true
}

