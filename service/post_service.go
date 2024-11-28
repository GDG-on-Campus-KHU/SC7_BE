package service

import (
	"bytes"
	"encoding/json"
	"github.com/GDG-on-Campus-KHU/SC7_BE/model"
	"github.com/GDG-on-Campus-KHU/SC7_BE/repository"
	"log"
	"net/http"
)

func CreatePost(post *model.Post) (int64, error) {
	return repository.CreatePost(post)
}

// SendToAI: AI 서버로 게시글 데이터를 전송
func SendToAI(post *model.Post) {
	// AI 서버의 URL
	url := "http://ai-server:8000/process" // AI 서버 주소

	// JSON 데이터 생성
	postData, err := json.Marshal(post)
	if err != nil {
		log.Printf("Failed to marshal post data for AI: %v", err)
		return
	}

	// HTTP POST 요청
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(postData))
	if err != nil {
		log.Printf("Failed to send post data to AI: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("Successfully sent post to AI server. Response Status: %s", resp.Status)
}

func GetALLPosts() ([]model.FilteredPost, error) {
	return repository.GetALLPosts()
}

func DeletePost(id string) error {
	return repository.DeletePostByID(id)
}
