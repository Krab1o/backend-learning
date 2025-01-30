package repository

import (
	"blogging-platform/internal/db"
	"blogging-platform/internal/model"
	"database/sql"
	"log"
)

type TagRepositoryPostgres struct{}

type TagRepository interface {
	Add(newPost *model.Tag) error
	Get() ([]model.Tag, error)
	GetOne(id uint) (*model.Tag, error)
	Delete(id uint) error
	Update(id uint, updatedPost *model.Tag) error
}

func NewTagRepositoryPostgres() TagRepository {
	return &TagRepositoryPostgres{}
}

func (s *TagRepositoryPostgres) Add(newTag *model.Tag) error {
	const insertStmt = 
	`INSERT INTO tags (title) 
	VALUES ($1)`	
	_, err := db.DB.Exec(insertStmt,
		newTag.Title,
	)
	if err != nil {
		log.Printf("Store: %v", err)
	} else {
		log.Println("Store: inserted successfully")
	}
	return nil
}

func (s *TagRepositoryPostgres) Get() ([]model.Tag, error) {
	const selectStmt = `SELECT * FROM tags`
	rows, err := db.DB.Query(selectStmt)
	if err != nil {
		log.Printf("Store: %v", err)
	} else {
		log.Println("Store: selected successfully")
	}
	defer rows.Close()

	var tags []model.Tag
	var tag model.Tag

	for rows.Next() {
		err := rows.Scan(
			&tag.ID, 
			&tag.Title,
		)
		if (err != nil) {
			log.Printf("Store: %v", err)
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func (s *TagRepositoryPostgres) GetOne(id uint) (*model.Tag, error) {
	tag := &model.Tag{}
	const selectOneStmt = `SELECT * FROM tags WHERE id = $1`
	err := db.DB.QueryRow(selectOneStmt, id).Scan(
		&tag.ID,
		&tag.Title,
	)
	if err == sql.ErrNoRows {
		log.Printf("Store: nothing was found")
	} else if err != nil{
		log.Printf("Store: failed to get post, %v", err)
	} else {
		log.Println("Store: selected by id successfully")
	}
	return tag, err
}

func (s *TagRepositoryPostgres) Delete(id uint) error {
	const deleteStmt = `DELETE FROM tags WHERE id = $1`
	_, err := db.DB.Exec(deleteStmt, id)
	if err != nil {
		log.Printf("Store: %v", err)
	} else {
		log.Println("Store: deleted successfully")
	}
	return err
}

func (s *TagRepositoryPostgres) Update(id uint, updatedTag *model.Tag) error {
	const updateStmt = `UPDATE tags SET 
	title = $2,
WHERE id = $1
`
	_, err := db.DB.Exec(updateStmt, 
		id,
		updatedTag.Title,
	)
	if err != nil {
		log.Printf("Store: %v", err)
	} else {
		log.Println("Store: updated successfully")
	}
	return err
}