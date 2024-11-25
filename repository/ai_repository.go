package repository

import (
	"github.com/GDG-on-Campus-KHU/SC7_BE/db"
	"log"
)

// UpdatePostAI : 게시글 AI 데이터 업데이트
func UpdatePostAI(postID int, category string, accuracy float64) error {
	query := `
		UPDATE posts
		SET category = ?, accuracy = ?
		WHERE id = ?
	`
	_, err := db.DB.Exec(query, category, accuracy, postID)
	if err != nil {
		log.Printf("Failed to update post with AI data: %v", err)
	}
	return err
}
