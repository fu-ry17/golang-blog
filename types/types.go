package types

import "time"

type BlogStore interface {
	CreateBlog(payload *CreateBlogPayload) (*Blog, error)
	GetBlogs() ([]Blog, error)
	GetBlogById(id int) (*Blog, error)
	UpdateBlog(id int, payload *CreateBlogPayload) (*Blog, error)
	DeleteBlog(id int) error
}

type Blog struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CreateBlogPayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
