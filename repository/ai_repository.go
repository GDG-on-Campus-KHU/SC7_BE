package repository

import (
	"github.com/GDG-on-Campus-KHU/SC7_BE/db"
	"github.com/GDG-on-Campus-KHU/SC7_BE/model"
	"log"
)

// UpdatePostAI : 게시글 AI 데이터 업데이트
func UpdatePostAI(post *model.Post) error {
	query := `
		UPDATE posts
		SET category = ?, accuracy = ?
		WHERE id = ?
	`

	_, err := db.DB.Exec(query, post.Category, post.Accuracy, post.ID)
	if err != nil {
		log.Printf("Failed to update post with AI data: %v", err)
	}
	return err
}
