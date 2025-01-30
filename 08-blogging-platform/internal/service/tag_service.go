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
	CreateTag(newPost *model.Tag) error
	GetTags() ([]model.Tag, error)
	GetTagByID(id uint) (*model.Tag, error)
	DeleteTag(id uint) error
	UpdateTag(id uint, updatedPost *model.Tag) error
}

func NewTagService() TagService {
	return &tagService{
		tagRepository: repository.NewTagRepositoryPostgres(),
	}
}

func (b *tagService) CreateTag(newPost *model.Tag) error {
	err := b.tagRepository.Add(newPost)
	if err != nil {
		log.Printf("Service: failed to create post, %v", err)
	}
	return err
}

func (b *tagService) GetTags() ([]model.Tag, error) {
	posts, err := b.tagRepository.Get()
	if err != nil {
		log.Printf("Service: failed to get posts, %v", err)
	}
	return posts, err
}

func (b *tagService) GetTagByID(id uint) (*model.Tag, error) {
	post, err := b.tagRepository.GetOne(id)
	if err != nil {
		log.Printf("Service: failed to get post, %v", err)
	}
	return post, err
}

func (b *tagService) DeleteTag(id uint) error {
	err := b.tagRepository.Delete(id)
	if err != nil {
		log.Printf("Service: failed to delete post, %v", err)
	}
	return err
}

func (b *tagService) UpdateTag(id uint, updatedPost *model.Tag) error {
	err := b.tagRepository.Update(id, updatedPost)
	if err != nil {
		log.Printf("Service: failed to update post, %v", err)
	}
	return err
}