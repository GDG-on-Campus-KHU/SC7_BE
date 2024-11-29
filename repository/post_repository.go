package repository

import (
	"encoding/json"
	"fmt"
	"github.com/GDG-on-Campus-KHU/SC7_BE/db"
	"github.com/GDG-on-Campus-KHU/SC7_BE/model"
	"log"
	"time"
)

func CreatePost(post *model.Post) (int64, error) {
	locationJSON, err := json.Marshal(post.Location)
	if err != nil {
		log.Printf("Failed to marshal location: %v", err)
		return 0, err
	}

	query := `
		INSERT INTO posts (user_id, text, image_path, location)
		VALUES (?, ?, ?, ?)
	`
	result, err := db.DB.Exec(query, post.UserID, post.Text, post.ImagePath, locationJSON)
	if err != nil {
		log.Printf("Failed to insert post: %v\n", err)
	}

	return result.LastInsertId()
}

func GetALLPosts() ([]model.FilteredPost, error) {
	query := `
		SELECT id, user_id, text, image_path, location, category, accuracy
		FROM posts
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var posts []model.FilteredPost

	for rows.Next() {
		post := model.Post{}
		var locationJSON string // JSON 문자열로 Location을 받음
		err := rows.Scan(&post.ID, &post.UserID, &post.Text, &post.ImagePath, &locationJSON, &post.Category, &post.Accuracy)
		if err != nil {
			log.Printf("failed to scan row: %v\n", err)
			continue
		}

		// Category가 없는 경우 대기
		if post.Category == nil {
			for i := 0; i < 5; i++ {
				time.Sleep(100 * time.Millisecond)
				// 게시글 상태 다시 확인, AI로 부터 데이터를 받아와야해서
				row := db.DB.QueryRow(query+" WHERE id = ?", post.ID)
				err := row.Scan(&post.ID, &post.UserID, &post.Text, &post.ImagePath, &locationJSON, &post.Category, &post.Accuracy)
				if err != nil {
					log.Printf("failed to recheck post category: %v\n", err)
					break
				}
				if post.Category != nil {
					break
				}
			}
		}

		// Location 데이터를 Unmarshal하여 [2]int로 변환
		var locationArray [2]float64
		if err := json.Unmarshal([]byte(locationJSON), &locationArray); err != nil {
			log.Printf("failed to unmarshal location JSON: %v\n", err)
			continue
		}

		// FilteredPost 구성
		filteredPost := model.FilteredPost{
			ID:        post.ID,
			UserID:    post.UserID,
			Text:      post.Text,
			ImagePath: post.ImagePath,
			Location:  locationArray,
			Category:  post.Category,
			Accuracy:  post.Accuracy,
		}

		posts = append(posts, filteredPost)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return posts, nil
}

func GetPostsByUserID(userId string) ([]model.Post, error) {
	query := `
		SELECT *
		FROM posts
		WHERE user_id = ?
	`

	rows, err := db.DB.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var posts []model.Post

	for rows.Next() {

		post := model.Post{}
		var locationJSON string
		err := rows.Scan(&post.ID, &post.UserID, &post.Text, &post.ImagePath, &locationJSON, &post.Category, &post.Accuracy, &post.CreatedAt)
		if err != nil {
			log.Printf("failed to scan row: %v\n", err)
			continue
		}

		var locationArray [2]float64
		if err := json.Unmarshal([]byte(locationJSON), &locationArray); err != nil {
			log.Printf("failed to unmarshal location JSON: %v\n", err)
			continue
		}
		post.Location = locationArray

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return posts, nil
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
