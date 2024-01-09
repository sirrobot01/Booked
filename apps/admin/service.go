package admin

import (
	"booked/models"
	"gorm.io/gorm"
)

type Service interface {
	GetUserById(id int) (*models.User, error)
	GetUsers() ([]*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
}

func NewService(db *gorm.DB) Service {
	return &service{db}
}

type service struct {
	db *gorm.DB
}

func (s *service) GetUserById(id int) (*models.User, error) {
	return nil, nil
}

func (s *service) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	s.db.Where("username = ?", username).First(&user)
	return &user, nil
}

func (s *service) GetUsers() ([]*models.User, error) {
	var users []*models.User
	s.db.Find(&users)
	return users, nil
}
