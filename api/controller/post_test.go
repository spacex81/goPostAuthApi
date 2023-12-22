package controller

import (
	"blog/api/repository"
	"blog/api/service"
	"blog/models"
	"blog/util"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPostsById(t *testing.T) {
	// Create a mock instance
	mockRepo := new(repository.MockPostRepository)
	postService := service.NewPostService(mockRepo)
	postController := NewPostController(postService)

	// Mock data
	posts := []models.Post{
		{ID: 1, Title: "Test Title 1", Body: "Test Body 1"},
		{ID: 1, Title: "Test Title 2", Body: "Test Body 2"}}
	mockRepo.On("GetPostsById", int64(1)).Return(&posts, int64(2), nil)

	// Setup Gin and register handler
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/posts/byuser/:id", postController.GetPostsById)

	// Create a request to the endpoint
	req, _ := http.NewRequest("GET", "/posts/byuser/1", bytes.NewBuffer([]byte{}))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse and check the response body
	var response util.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, true, err == nil)
	assert.Equal(t, "Post result set", response.Message)

	data, ok := response.Data.(map[string]interface{})
	assert.True(t, ok, "Type assertion for response data failed")

	totalRowsFloat, ok := data["total_rows"].(float64) // Assert as float64
	assert.True(t, ok, "Type assertion for total_rows failed")

	totalRows := int64(totalRowsFloat) // Convert to int64 if needed
	assert.Equal(t, totalRows, int64(2))

	rows := data["rows"]

	rowsSlice, ok := data["rows"].([]interface{})
	assert.True(t, ok, "Type assertion for rows failed")

	assert.Equal(t, len(rowsSlice), len(posts))

	for i, rowInterface := range rowsSlice {
		row, ok := rowInterface.(map[string]interface{})
		assert.True(t, ok, "Type assertion for individual row failed")

		// Assert data for each row
		assert.Equal(t, row["id"], float64(posts[i].ID))
		assert.Equal(t, row["title"], posts[i].Title)
		assert.Equal(t, row["body"], posts[i].Body)
	}

	log.Printf("Response rows: %+v", rows)
}
