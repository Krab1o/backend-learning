package service

import (
	"blogging-platform/internal/model"
	"blogging-platform/internal/repository"
	"log"
	"time"
)

type BasicService struct{
	repository repository.PostRepository
}

type PostService interface {
	CreatePost(newPost *model.Post) error
	GetPosts() ([]model.Post, error)
	GetPostByID(id uint) (*model.Post, error)
	DeletePost(id uint) error
	UpdatePost(id uint, updatedPost *model.Post) error
}

func NewPostService() PostService {
	return &BasicService{
		repository: repository.NewPostRepositoryPostgres(),
	}
}

func (b *BasicService) CreatePost(newPost *model.Post) error {
	newPost.CreatedAt = time.Now()
	newPost.UpdatedAt = newPost.CreatedAt
	err := b.repository.CreatePost(newPost)
	if err != nil {
		log.Printf("Service: failed to create post, %v", err)
	}
	return err
}

func (b *BasicService) GetPosts() ([]model.Post, error) {
	posts, err := b.repository.GetPosts()
	if err != nil {
		log.Printf("Service: failed to get posts, %v", err)
	}
	return posts, err
}

func (b *BasicService) GetPostByID(id uint) (*model.Post, error) {
	post, err := b.repository.GetPostByID(id)
	if err != nil {
		log.Printf("Service: failed to get post, %v", err)
	}
	return post, err
}

func (b *BasicService) DeletePost(id uint) error {
	err := b.repository.DeletePost(id)
	if err != nil {
		log.Printf("Service: failed to delete post, %v", err)
	}
	return err
}

func (b *BasicService) UpdatePost(id uint, updatedPost *model.Post) error {
	updatedPost.UpdatedAt = time.Now()
	err := b.repository.UpdatePost(id, updatedPost)
	if err != nil {
		log.Printf("Service: failed to update post, %v", err)
	}
	return err
}
