package repository

import (
	"encoding/json"
	"fmt"
	"github.com/GDG-on-Campus-KHU/SC7_BE/db"
	"github.com/GDG-on-Campus-KHU/SC7_BE/model"
	"log"
	"time"
)

func CreatePost(post *model.Post) error {
	locationJSON, err := json.Marshal(post.Location)
	if err != nil {
		log.Printf("Failed to marshal location: %v", err)
		return err
	}

	query := `
		INSERT INTO posts (user_id, text, image_path, location)
		VALUES (?, ?, ?, ?)
	`
	_, err = db.DB.Exec(query, post.UserID, post.Text, post.ImagePath, locationJSON)
	if err != nil {
		log.Printf("Failed to insert post: %v\n", err)
	}
	return err
}

func GetPostByID(id string) (*model.FilteredPost, error) {
	query := `
		SELECT id, user_id, text, image_path, location, category, accuracy
		FROM posts
		WHERE id = ?
	`
	post := &model.Post{}
	i := 0
	for {
		i++
		row := db.DB.QueryRow(query, id)

		var locationJSON string // JSON 문자열로 Location을 받음
		err := row.Scan(&post.ID, &post.UserID, &post.Text, &post.ImagePath, &locationJSON, &post.Category, &post.Accuracy)
		if err != nil {
			return nil, err
		}

		if i > 50 {
			return nil, fmt.Errorf("게시글 조회 실패: %v", err)
		}

		if post.Category == nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		// Location 데이터를 Unmarshal하여 [2]int로 변환
		var locationArray [2]int
		if err := json.Unmarshal([]byte(locationJSON), &locationArray); err != nil {
			return nil, fmt.Errorf("Failed to Location JSON Unmarshal: %v", err)
		}

		// FilteredPost 구성
		filteredPost := &model.FilteredPost{
			ID:        post.ID,
			UserID:    post.UserID,
			Text:      post.Text,
			ImagePath: post.ImagePath,
			Location:  locationArray,
			Category:  post.Category,
			Accuracy:  post.Accuracy,
		}

		return filteredPost, nil
	}
}

func DeletePostByID(id string) error {
	query := `
		DELETE FROM posts
		WHERE id = ?
	`
	_, err := db.DB.Exec(query, id)
	if err != nil {
		log.Printf("Failed to delete post: %v\n", err)
	}
	return err
}
