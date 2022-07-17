package user

import (
	"framework/database"
	"framework/internal/app/services"
	"framework/internal/models"
)

func (s *services.Service) First() (*models.User, error) {
	db := database.NewQueryBuilder()
}
