package service

import (
	"blog/api/repository"
	"blog/models"
)

type PostService struct {
	repository repository.PostRepositoryInterface
}

func NewPostService(r repository.PostRepositoryInterface) PostService {
	return PostService{
		repository: r,
	}
}

func (p PostService) Save(post models.Post) error {
	return p.repository.Save(post)
}

func (p PostService) FindAll(post models.Post, keyword string) (*[]models.Post, int64, error) {
	return p.repository.FindAll(post, keyword)
}

func (p PostService) GetPostsById(id int64) (*[]models.Post, int64, error) {
	return p.repository.GetPostsById(id)
}

func (p PostService) Update(post models.Post) error {
	return p.repository.Update(post)
}

func (p PostService) Delete(id int64) error {
	var post models.Post
	post.ID = id
	return p.repository.Delete(post)
}

func (p PostService) Find(post models.Post) (models.Post, error) {
	return p.repository.Find(post)
}

func (p PostService) UserExists(userID int64) bool {
	var user models.User
	err := p.repository.(repository.PostRepository).Db.DB.First(&user, userID).Error
	return err == nil && user.ID > 0
}
