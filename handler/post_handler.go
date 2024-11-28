package handler

import (
	"github.com/GDG-on-Campus-KHU/SC7_BE/model"
	"github.com/GDG-on-Campus-KHU/SC7_BE/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"sync"
)

// delete 와 get 사이의 동기화를 위한 뮤텍스
var mu sync.Mutex

/*
 create 과 get 사이의 동기화는 AI 서버로 부터 accuracy 값을 받아오는 것을 기다리는 것으로 구현
 단순히 create 이 완료될 때까지 기다렸다가 get 을 하게 되면, AI 서버로 부터 accuracy 값을 받아오기 전에 get 이 되어버리기 때문에
 mutex lock 을 사용하는 것이 아니라 AI 서버로부터 accuracy 값이 갱신될 때까지 대기하도록 구현
*/

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
	id, err := service.CreatePost(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}
	post.ID = int(id)
	// AI 서버로 비동기 요청 전송
	go service.SendToAI(&post)

	// 클라이언트에 성공 응답 반환
	c.JSON(http.StatusCreated, post)
}

// GetPost 핸들러: 게시글 조회
func GetALLPost(c *gin.Context) {
	mu.Lock()         // get 이 끝날 때까지 delete 를 막기 위해 뮤텍스 락
	defer mu.Unlock() // 작업 완료 후 해제

	posts, err := service.GetALLPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	mu.Lock()         // delete 가 끝날 때까지 get 을 막기 위해 뮤텍스 락
	defer mu.Unlock() // 작업 완료 후 해제

	err := service.DeletePost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
