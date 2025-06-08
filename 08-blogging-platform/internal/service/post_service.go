package service

import (
	"blogging-platform/internal/model"
	"blogging-platform/internal/repository"
	"log"
	"time"
)

type postService struct{
	tagService			TagService

	postRepository 		repository.PostRepository
	postTagRepository	repository.PostTagRepository
}

type PostService interface {
	CreatePost(newPost *model.Post) error
	GetPosts(searchTerm string) ([]model.Post, error)
	GetPostByID(id uint) (*model.Post, error)
	GetPostByTag(tag string) ([]model.Post, error)
	DeletePost(id uint) error
	UpdatePost(updatedPost *model.Post) error
}

func NewPostService() PostService {
	return &postService{
		tagService: NewTagService(),

		postTagRepository: repository.NewPostTagRepository(),
		postRepository: repository.NewPostRepository(),
	}
}

func (p *postService) CreatePost(newPost *model.Post) error {
	newPost.CreatedAt = time.Now()
	newPost.UpdatedAt = newPost.CreatedAt

	tagsID, _ := p.tagService.CreateTags(newPost.Tags)
	postID, err := p.postRepository.Add(newPost)

	links := make([]model.Link, 0, len(newPost.Tags))
	for _, id := range tagsID {
		links = append(links, model.Link{
			PostID: postID,
			TagID: id,
		})
	}
	p.postTagRepository.AddLinks(links)
	
	if err != nil {
		log.Printf("Service: failed to create post, %v", err)
	}
	return err
}

func (p *postService) GetPosts(searchTerm string) ([]model.Post, error) {
	var posts []model.Post
	var err error
	if searchTerm == "" {
		log.Println("basic search")
		posts, err = p.postRepository.Get()
	} else {
		log.Println("term search")
		posts, err = p.postRepository.GetByTerm(searchTerm)
	}
	if err != nil {
		log.Printf("Service: failed to get posts, %v", err)
	}
	return posts, err
}

func (p *postService) GetPostByID(id uint) (*model.Post, error) {
	post, err := p.postRepository.GetOne(id)
	if err != nil {
		log.Printf("Service: failed to get post, %v", err)
	}
	return post, err
}

func (p *postService) DeletePost(id uint) error {
	err := p.postRepository.Delete(id)
	if err != nil {
		log.Printf("Service: failed to delete post, %v", err)
	}
	return err
}

func (p *postService) UpdatePost(updatedPost *model.Post) error {
	updatedPost.UpdatedAt = time.Now()
	log.Println(updatedPost)
	err := p.postTagRepository.RemoveLinksByID(updatedPost.ID)
	if err != nil {
		log.Printf("Service post update: failed to update post (removed links), %v", err)
	}
	err = p.postRepository.Update(updatedPost.ID, updatedPost)
	if err != nil {
		log.Printf("Service post update: failed to update post (updatedPost), %v", err)
	}

	tagsID, err := p.tagService.CreateTags(updatedPost.Tags)
	links := make([]model.Link, 0, len(tagsID))
	for _, tagID := range tagsID {
		links = append(links, model.Link{
			PostID: updatedPost.ID,
			TagID: tagID, 
		})
	}
	if err != nil {
		log.Printf("Service post update: failed to update post (new tags), %v", err)
	}
	err = p.postTagRepository.AddLinks(links)
	if err != nil {
		log.Printf("Service post update: failed to update post (added links), %v", err)
	}
	return err
}

func (p *postService) GetPostByTag(tag string) ([]model.Post, error) {
	log.Println("not implemented")
	return []model.Post{}, nil
}
