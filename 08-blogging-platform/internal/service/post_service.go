package service

import (
	"blogging-platform/internal/model"
	"blogging-platform/internal/repository"
	"log"
	"time"
)

type postService struct{
	postRepository 	repository.PostRepository
}

type PostService interface {
	CreatePost(newPost *model.Post) error
	GetPosts() ([]model.Post, error)
	GetPostByID(id uint) (*model.Post, error)
	DeletePost(id uint) error
	UpdatePost(id uint, updatedPost *model.Post) error
}

func NewPostService() PostService {
	return &postService{
		postRepository: repository.NewPostRepositoryPostgres(),
	}
}

func (b *postService) CreatePost(newPost *model.Post) error {
	newPost.CreatedAt = time.Now()
	newPost.UpdatedAt = newPost.CreatedAt
	err := b.postRepository.Add(newPost)
	if err != nil {
		log.Printf("Service: failed to create post, %v", err)
	}
	return err
}

func (b *postService) GetPosts() ([]model.Post, error) {
	posts, err := b.postRepository.Get()
	if err != nil {
		log.Printf("Service: failed to get posts, %v", err)
	}
	return posts, err
}

func (b *postService) GetPostByID(id uint) (*model.Post, error) {
	post, err := b.postRepository.GetOne(id)
	if err != nil {
		log.Printf("Service: failed to get post, %v", err)
	}
	return post, err
}

func (b *postService) DeletePost(id uint) error {
	err := b.postRepository.Delete(id)
	if err != nil {
		log.Printf("Service: failed to delete post, %v", err)
	}
	return err
}

func (b *postService) UpdatePost(id uint, updatedPost *model.Post) error {
	updatedPost.UpdatedAt = time.Now()
	err := b.postRepository.Update(id, updatedPost)
	if err != nil {
		log.Printf("Service: failed to update post, %v", err)
	}
	return err
}
