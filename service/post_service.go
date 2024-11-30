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

type AIResponseBody struct {
	ID       int     `json:"id"`
	Category string  `json:"prediction"`
	Accuracy float64 `json:"highest_probability"`
}

// SendToAI: AI 서버로 게시글 데이터를 전송
func SendToAI(imagePath string, text string, id int) {
	// AI 서버의 URL
	url := "http://192.168.22.237:8000/predict" // AI 서버 주소

	requestBody := AIRequestBody{
		ImagePath: imagePath,
		Text:      text,
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

	log.Printf("AI 서버로 POST 요청을 보냈습니다. 응답: %v\n", resp)

	// 응답 데이터 읽기
	var responseBody AIResponseBody
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		log.Printf("Failed to decode AI response: %v", err)
		return
	}

	// 응답 로그 출력
	if resp.StatusCode == http.StatusOK {
		log.Printf("Successfully sent post to AI server. Response: %+v\n", responseBody)

		// Post 업데이트
		post := model.Post{
			ID:       id,
			Category: &responseBody.Category,
			Accuracy: &responseBody.Accuracy,
		}

		err = repository.UpdatePostAI(&post)
		if err != nil {
			log.Printf("Failed to update post in database: %v", err)
		}
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
