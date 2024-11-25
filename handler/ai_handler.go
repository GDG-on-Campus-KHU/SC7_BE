package handler

import (
	"github.com/GDG-on-Campus-KHU/SC7_BE/model"
	"github.com/GDG-on-Campus-KHU/SC7_BE/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdatePostWithAI(c *gin.Context) {
	var post model.Post

	// JSON 바인딩
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid AI response"})
		return
	}

	// 서비스 호출
	if err := service.UpdatePostWithAI(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post with AI data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}
