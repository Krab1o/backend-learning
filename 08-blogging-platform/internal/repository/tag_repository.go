package repository

import (
	"blogging-platform/internal/db"
	"blogging-platform/internal/model"
	"fmt"
	"log"
	"strings"
)

type tagRepositoryPostgres struct{}

type TagRepository interface {
	Add(newTags []model.Tag) ([]uint, error)
}

func NewTagRepository() TagRepository {
	return &tagRepositoryPostgres{}
}

func (s *tagRepositoryPostgres) Add(newTags []model.Tag) ([]uint, error) {
	valueStrings := make([]string, 0, len(newTags))
	valueArgs := make([]any, 0, len(newTags))

	for ind, tag := range newTags {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d)", ind + 1))
		valueArgs = append(valueArgs, tag.Title)
	}
	log.Println("TAG REPO TAGS: ", newTags)
	query := fmt.Sprintf(
		`INSERT INTO tags (tag) 
		VALUES %s 
		ON CONFLICT (tag)
		DO UPDATE SET tag = EXCLUDED.tag
		RETURNING id
		`, strings.Join(valueStrings, ","))
	
	rows, err := db.DB.Query(query, valueArgs...)
	if err != nil {
		fmt.Println(err)
	}

	var ids = make([]uint, 0, len(newTags))
	var id uint
	for rows.Next() {
		rows.Scan(&id)
		ids = append(ids, id)
	}
	log.Println("TAG REPO IDs: ", ids)
	if err != nil {
		log.Printf("Store: %v", err)
	} else {
		log.Println("Store: tags inserted successfully")
	}
	return ids, err
}