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

type AIRequestBody struct {
	ImagePath string `json:"url"`
	Text      string `json:"text"`
	ID        int    `json:"id"`
}

// SendToAI: AI 서버로 게시글 데이터를 전송
func SendToAI(imagePath string, text string, id int) {
	// AI 서버의 URL
	url := "http://192.168.22.237:8000/predict" // AI 서버 주소

	requestBody := AIRequestBody{
		ImagePath: imagePath,
		Text:      text,
		ID:        id,
	}

	// JSON 데이터 생성
	postData, err := json.Marshal(requestBody)
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

	// 응답 로그 출력
	if resp.StatusCode == http.StatusOK {
		log.Printf("Successfully sent post to AI server. Response Status: %s", resp.Status)
	} else {
		log.Printf("Failed to send post to AI server. Response Status: %s", resp.Status)
	}
}

func GetALLPosts() ([]model.FilteredPost, error) {
	return repository.GetALLPosts()
}

func GetPostsByUserID(userId string) ([]model.Post, error) {
	return repository.GetPostsByUserID(userId)
}

func DeletePost(id string) error {
	return repository.DeletePostByID(id)
}
