package handler

import (
	"github.com/GDG-on-Campus-KHU/SC7_BE/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdatePostWithAI(c *gin.Context) {
	var aiResponse struct {
		PostID   int     `json:"post_id"`  // 게시글 ID
		Category string  `json:"category"` // AI가 반환한 카테고리
		Accuracy float64 `json:"accuracy"` // AI가 반환한 정확도
	}

	// JSON 바인딩
	if err := c.ShouldBindJSON(&aiResponse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid AI response"})
		return
	}

	// DB 갱신
	if err := repository.UpdatePostAI(aiResponse.PostID, aiResponse.Category, aiResponse.Accuracy); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post with AI data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}
