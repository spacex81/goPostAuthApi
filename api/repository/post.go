package repository

import (
	"blog/infrastructure"
	"blog/models"
)

type PostRepositoryInterface interface {
	Save(post models.Post) error
	FindAll(post models.Post, keyword string) (*[]models.Post, int64, error)
	GetPostsById(id int64) (*[]models.Post, int64, error)
	Update(post models.Post) error
	Find(post models.Post) (models.Post, error)
	Delete(post models.Post) error
}

type PostRepository struct {
	Db infrastructure.Database
}

func NewPostRepository(db infrastructure.Database) PostRepository {
	return PostRepository{
		Db: db,
	}
}

func (p PostRepository) Save(post models.Post) error {
	return p.Db.DB.Create(&post).Error
}

func (p PostRepository) FindAll(post models.Post, keyword string) (*[]models.Post, int64, error) {
	var posts []models.Post
	var totalRows int64 = 0

	queryBuilder := p.Db.DB.Order("created_at desc").Model(&models.Post{})

	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where(
			p.Db.DB.Where("post.title LIKE ? ", queryKeyword))
	}

	err := queryBuilder.Where(post).Find(&posts).Count(&totalRows).Error
	return &posts, totalRows, err
}

func (p PostRepository) GetPostsById(id int64) (*[]models.Post, int64, error) {
	var posts []models.Post
	var totalRows int64 = 0

	queryBuilder := p.Db.DB.Order("created_at desc").Model(&models.Post{})

	queryBuilder.Where("user_id = ?", id)

	err := queryBuilder.Find(&posts).Count(&totalRows).Error
	return &posts, totalRows, err
}

func (p PostRepository) Update(post models.Post) error {
	return p.Db.DB.Save(&post).Error
}

func (p PostRepository) Find(post models.Post) (models.Post, error) {
	var posts models.Post
	err := p.Db.DB.Debug().Model(&models.Post{}).Where(&post).Take(&posts).Error
	return posts, err
}

func (p PostRepository) Delete(post models.Post) error {
	return p.Db.DB.Delete(&post).Error
}
