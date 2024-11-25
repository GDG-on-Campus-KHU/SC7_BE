package service

import (
	"github.com/GDG-on-Campus-KHU/SC7_BE/model"
	"github.com/GDG-on-Campus-KHU/SC7_BE/repository"
)

func UpdatePostWithAI(post *model.Post) error {
	return repository.UpdatePostAI(post)
}
