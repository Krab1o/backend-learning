package repository

import (
	"blogging-platform/internal/db"
	"blogging-platform/internal/model"
	"database/sql"
	"log"
)

type PostRepositoryPostgres struct{}

type link struct {
	postID 	uint
	tag		string
}

type PostRepository interface {
	Add(newPost *model.Post) (uint, error)
	Get() ([]model.Post, error)
	GetByTerm(searchTerm string) ([]model.Post, error)
	GetOne(id uint) (*model.Post, error)
	Delete(id uint) error
	Update(id uint, updatedPost *model.Post) error
}

func NewPostRepository() PostRepository {
	return &PostRepositoryPostgres{}
}

//TODO: use transactions

func (s *PostRepositoryPostgres) Add(newPost *model.Post) (uint, error) {
	const insertStmt = 
	`INSERT INTO posts (title, content, category, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id`	
	res := db.DB.QueryRow(insertStmt,
		newPost.Title,
		newPost.Content,
		newPost.Category,
		newPost.CreatedAt,
		newPost.UpdatedAt,
	)
	var id uint
	err := res.Scan(&id)

	if err != nil {
		log.Printf("Repo posts: %v", err)
	} else {
		log.Printf("Repo posts: post #%d inserted successfully", id)
	}
	return id, nil
}

func (s *PostRepositoryPostgres) Get() ([]model.Post, error) {
	const selectIDs = `
	SELECT p.id, t.tag
	FROM posts AS p
	LEFT JOIN posts_tags AS pt ON (p.id = pt.id_post)
	LEFT JOIN tags AS t ON (pt.id_tag = t.id)
	`
	rowsTags, err := db.DB.Query(selectIDs)
	if err != nil {
		log.Printf("Repo get posts: tags %v", err)
	} else {
		log.Println("Repo get posts: tags selected successfully")
	}
	defer rowsTags.Close()

	const selectData = `
	SELECT p.id, p.title, p.content, p.category, p.created_at, p.updated_at 
	FROM posts AS p
	`
	rowsData, err := db.DB.Query(selectData)
	if err != nil {
		log.Printf("Repo get posts: data %v", err)
	} else {
		log.Println("Repo get posts: data selected successfully")
	}
	defer rowsData.Close()

	links := []link{}
	l := link{}
	for rowsTags.Next() {
		rowsTags.Scan(
			&l.postID,
			&l.tag,
		)
		links = append(links, l)
	}
	var posts []model.Post
	var post model.Post

	for rowsData.Next() {
		rowsData.Scan(
			&post.ID, 
			&post.Title, 
			&post.Content, 
			&post.Category, 
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		post.Tags = []string{}
		for _, l := range links {
			if post.ID == l.postID {
				post.Tags = append(post.Tags, l.tag)
			}
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *PostRepositoryPostgres) GetByTerm(searchTerm string) ([]model.Post, error) {
	searchPattern := "%" + searchTerm + "%"
	const selectData = `
	SELECT p.id, p.title, p.content, p.category, p.created_at, p.updated_at, t.tag
	FROM posts AS p
	RIGHT JOIN posts_tags AS pt ON (p.id = pt.id_post)
	LEFT JOIN tags AS t ON (pt.id_tag = t.id)
	WHERE 
	p.title LIKE $1 OR
	p.content LIKE $1 OR
	p.category LIKE $1 OR
	t.tag LIKE $1
	`
	log.Println(selectData)
	rowsData, err := db.DB.Query(selectData, searchPattern)
	if err != nil {
		log.Printf("Repo get posts: HERE data %v", err)
	} else {
		log.Println("Repo get posts: data selected successfully")
	}
	defer rowsData.Close()

	var posts []model.Post
	var tag string
	var prevPost model.Post
	var post model.Post

	for rowsData.Next() {
		rowsData.Scan(
			&post.ID, 
			&post.Title, 
			&post.Content, 
			&post.Category, 
			&post.CreatedAt,
			&post.UpdatedAt,
			&tag,
		)
		if (prevPost.ID == 0) {
			prevPost = post
		}
		if post.ID != prevPost.ID { // новый пост
			posts = append(posts, prevPost)
			prevPost = post
		}
		prevPost.Tags = append(prevPost.Tags, tag)
	}
	posts = append(posts, prevPost)
	return posts, nil
}


func (s *PostRepositoryPostgres) GetOne(prevID uint) (*model.Post, error) {
	post := &model.Post{}
	const selectIDs = `
	SELECT t.tag
	FROM posts AS p
	LEFT JOIN posts_tags AS pt ON (p.id = pt.id_post)
	LEFT JOIN tags AS t ON (pt.id_tag = t.id)
	WHERE p.id = $1
	`
	rowsTags, err := db.DB.Query(selectIDs, prevID)
	if err != nil {
		log.Printf("Repo get one post: tags %v", err)
	} else {
		log.Println("Repo get one post: tags selected successfully")
	}
	const selectOneStmt = `SELECT * FROM posts WHERE id = $1`
	err = db.DB.QueryRow(selectOneStmt, prevID).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Category,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		log.Printf("Repo get one post: nothing was found")
	} else if err != nil{
		log.Printf("Repo get one post: failed to get post, %v", err)
	} else {
		log.Println("Repo get one post: selected by id successfully")
	}
	
	var tag string
	for rowsTags.Next() {
		rowsTags.Scan(
			&tag,
		)
		post.Tags = append(post.Tags, tag)
	}
	
	return post, err
}

func (s *PostRepositoryPostgres) Delete(id uint) error {
	const deleteStmt = `DELETE FROM posts WHERE id = $1`
	_, err := db.DB.Exec(deleteStmt, id)
	if err != nil {
		log.Printf("Store: %v", err)
	} else {
		log.Println("Repo delete one post: deleted successfully")
	}
	return err
}

func (s *PostRepositoryPostgres) Update(id uint, updatedPost *model.Post) error {
	const updateStmt = `UPDATE posts SET 
	title = $2,
	content = $3,
	category = $4,
	created_at = $5,
	updated_at = $6
WHERE id = $1
`
	_, err := db.DB.Exec(updateStmt, 
		id,
		updatedPost.Title,
		updatedPost.Content,
		updatedPost.Category,
		updatedPost.CreatedAt,
		updatedPost.UpdatedAt,
	)
	if err != nil {
		log.Printf("Repo upd one post: %v", err)
	} else {
		log.Println("Repo upd one post: updated successfully")
	}
	return err
}

