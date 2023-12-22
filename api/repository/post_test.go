package repository

import (
	"blog/infrastructure"
	"blog/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestPostRepository_GetPostsById(t *testing.T) {
	// Setup SQL mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database", err)
	}

	// GORM setup with mock db
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	// Setup expectations
	rows := sqlmock.NewRows([]string{"id", "title", "body", "user_id", "created_at", "updated_at"}).
		AddRow(1, "Test Title 1", "Test Body 1", 1, nil, nil).
		AddRow(2, "Test Title 2", "Test Body 2", 1, nil, nil)

	mock.ExpectQuery("^SELECT \\* FROM `post` WHERE user_id = \\? ORDER BY created_at desc").
		WithArgs(1).
		WillReturnRows(rows)

	countRows := sqlmock.NewRows([]string{"count"}).AddRow(2)
	mock.ExpectQuery("^SELECT count\\(\\*\\) FROM `post` WHERE user_id = \\?").
		WithArgs(1).
		WillReturnRows(countRows)

	// Create repository instance
	infraDb := infrastructure.Database{DB: gormDB}
	postRepository := NewPostRepository(infraDb)

	// Run SQL
	posts, totalRows, err := postRepository.GetPostsById(1)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, posts)
	//assert.Equal(t, int64(2), totalRows)
	assert.Equal(t, int64(2), totalRows)
	// let's add the test so that we can check each data of each row

	expectedPosts := []models.Post{
		{ID: 1, Title: "Test Title 1", Body: "Test Body 1"},
		{ID: 2, Title: "Test Title 2", Body: "Test Body 2"},
	}
	assert.Equal(t, len(expectedPosts), len(*posts))

	for i, post := range *posts {
		assert.Equal(t, expectedPosts[i].ID, post.ID)
		assert.Equal(t, expectedPosts[i].Title, post.Title)
		assert.Equal(t, expectedPosts[i].Body, post.Body)
	}
}
