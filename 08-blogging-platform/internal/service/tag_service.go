package service

import (
	"blogging-platform/internal/model"
	"blogging-platform/internal/repository"
	"log"
)

type tagService struct{
	tagRepository 	repository.TagRepository
}

type TagService interface {
	CreateTags(titles []string) ([]uint, error)
}

func NewTagService() TagService {
	return &tagService{
		tagRepository: repository.NewTagRepository(),
	}
}

func (b *tagService) CreateTags(titles []string) ([]uint, error) {
	tags := make([]model.Tag, 0, len(titles))
	for _, title := range titles {
		tags = append(tags, model.Tag{Title: title})
	}
	tagIDs, err := b.tagRepository.Add(tags)
	log.Println("Service tag: ", tagIDs)
	if err != nil {
		log.Printf("Service tag: failed to create tag, %v", err)
	}
	return tagIDs, err
}