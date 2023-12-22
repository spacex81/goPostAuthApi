package repository

import (
	"blog/models"
	"github.com/stretchr/testify/mock"
)

type MockPostRepository struct {
	mock.Mock
}

func (mock *MockPostRepository) Save(post models.Post) error {
	//TODO implement me
	panic("implement me")
}

func (mock *MockPostRepository) FindAll(post models.Post, keyword string) (*[]models.Post, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (mock *MockPostRepository) Update(post models.Post) error {
	//TODO implement me
	panic("implement me")
}

func (mock *MockPostRepository) Find(post models.Post) (models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (mock *MockPostRepository) Delete(post models.Post) error {
	//TODO implement me
	panic("implement me")
}

func (mock *MockPostRepository) GetPostsById(id int64) (*[]models.Post, int64, error) {
	args := mock.Called(id)
	return args.Get(0).(*[]models.Post), args.Get(1).(int64), args.Error(2)
}
