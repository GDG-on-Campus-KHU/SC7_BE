package handler

import (
	"github.com/GDG-on-Campus-KHU/SC7_BE/model"
	"github.com/GDG-on-Campus-KHU/SC7_BE/service"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// CreatePost 핸들러: 게시글 생성
func CreatePost(c *gin.Context) {
	var post model.Post

	// JSON 데이터 바인딩
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 이미지 파일 업로드
	file, err := c.FormFile("image")
	if err == nil {
		savePath := filepath.Join("uploads", file.Filename)
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}
		post.ImagePath = savePath
	}

	// 게시글 생성
	if err := service.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	// AI 서버로 비동기 요청 전송
	go service.SendToAI(&post)

	// 클라이언트에 성공 응답 반환
	c.JSON(http.StatusCreated, post)
}

// GetPost 핸들러: 게시글 조회
func GetPost(c *gin.Context) {
	id := c.Param("id")
	post, err := service.GetPost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch post"})
		return
	}

	c.JSON(http.StatusOK, post)
}
