package repository

import (
	"blogging-platform/internal/db"
	"blogging-platform/internal/model"
	"fmt"
	"strings"
)

type postTagRepositoryPostgres struct{}

type PostTagRepository interface {
	AddLinks(links[]model.Link) error
	RemoveLinksByID(id uint) error
}

func NewPostTagRepository() PostTagRepository {
	return &postTagRepositoryPostgres{}
}

func (r *postTagRepositoryPostgres) AddLinks(links []model.Link) error {
	if len(links) == 0 {
		return nil
	}

	valueStrings := make([]string, 0, len(links))
	valueArgs := make([]any, 0, len(links) * 2)

	for ind, link := range links {
		valueStrings = append(valueStrings, fmt.Sprintf(
			"($%d, $%d)",
			ind * 2 + 1,
			ind * 2 + 2, 
			))
		valueArgs = append(valueArgs, link.PostID, link.TagID)
	}
	
	query := fmt.Sprintf("INSERT INTO posts_tags (id_post, id_tag) VALUES %s", strings.Join(valueStrings, ","))
	_, err := db.DB.Exec(query, valueArgs...)
	return err
}

func (r *postTagRepositoryPostgres) RemoveLinksByID(id uint) error {
	query := `
	DELETE FROM posts_tags
	WHERE id_post = $1
	`
	_, err := db.DB.Exec(query, id)
	return err
}