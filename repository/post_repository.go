package repository

import (
	"encoding/json"
	"github.com/GDG-on-Campus-KHU/SC7_BE/db"
	"github.com/GDG-on-Campus-KHU/SC7_BE/model"
	"log"
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

func GetPostByID(id string) (*model.Post, error) {
	query := `
		SELECT id, user_id, text, image_path, location, category, accuracy
		FROM posts
		WHERE id = ?
	`

	post := &model.Post{}
	row := db.DB.QueryRow(query, id)

	var locationJSON string
	err := row.Scan(&post.ID, &post.UserID, &post.Text, &post.ImagePath, &locationJSON, &post.Category, &post.Accuracy)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(locationJSON), &post.Location); err != nil {
		return nil, err
	}

	return post, nil
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
