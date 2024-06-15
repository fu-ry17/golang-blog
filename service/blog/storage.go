package blog

import (
	"fmt"
	"gorm.io/gorm"
	"simple-blog/db/models"
	"simple-blog/types"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{db: db}
}

func mapResponseToDto(blog *models.Blog) *types.Blog {
	return &types.Blog{
		Id:          blog.Id,
		Title:       blog.Title,
		Description: blog.Description,
		CreatedAt:   blog.CreatedAt,
		UpdatedAt:   blog.UpdatedAt,
	}
}

func (s *Storage) CreateBlog(payload *types.CreateBlogPayload) (*types.Blog, error) {
	blog := models.Blog{
		Title:       payload.Title,
		Description: payload.Description,
	}

	if err := s.db.Save(&blog).Error; err != nil {
		return nil, err
	}

	return mapResponseToDto(&blog), nil
}

func (s *Storage) GetBlogs() ([]types.Blog, error) {
	var blogs []models.Blog
	if err := s.db.Find(&blogs).Error; err != nil {
		return nil, err
	}

	responseBlogs := make([]types.Blog, len(blogs))
	for i, blog := range blogs {
		responseBlogs[i] = *mapResponseToDto(&blog)
	}

	return responseBlogs, nil
}

func (s *Storage) GetBlogById(id int) (*types.Blog, error) {
	var blog models.Blog
	if err := s.db.First(&blog, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("blog with id %d not found", id)
	}

	return mapResponseToDto(&blog), nil
}

func (s *Storage) UpdateBlog(id int, payload *types.CreateBlogPayload) (*types.Blog, error) {
	blog := models.Blog{
		Title:       payload.Title,
		Description: payload.Description,
	}

	if err := s.db.Model(&models.Blog{}).Where("id = ?", id).Updates(&blog).Error; err != nil {
		return nil, err
	}

	return mapResponseToDto(&blog), nil
}

func (s *Storage) DeleteBlog(id int) error {
	if err := s.db.Delete(&models.Blog{}, "id = ?", id).Error; err != nil {
		fmt.Errorf("blog with id %d not found", id)
	}

	return nil
}
