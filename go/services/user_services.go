package services

import "webapp-go/models"
import "github.com/google/uuid"

func UserCreate(userData models.UserCreate) models.UserDetail {
	user := models.UserDetail{
		ID:   uuid.New().String(),
		Name: userData.Name,
	}
	return user
}

func UserDetail(ID string) models.UserDetail {
	user := models.UserDetail{
		ID:   uuid.New().String(),
		Name: uuid.New().String(),
	}
	return user
}
