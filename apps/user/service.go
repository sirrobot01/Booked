package user

import (
	"booked/common"
	"booked/models"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

type Service interface {
	GetUserById(id int) (*models.User, error)
	GetUsers() ([]*models.User, error)
	CreateUser(user UserRegister) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	Authenticate(user *UserLogin) (map[string]string, error)
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

func (s *service) CreateUser(user UserRegister) (*models.User, error) {
	var u models.User
	u = *common.CreateUser(user.Username, user.FirstName, user.LastName, user.Email, user.Password, false, true)
	if err := s.db.Create(&u).Error; err != nil {
		return &u, err
	}
	return &u, nil
}

func (s *service) Authenticate(user *UserLogin) (map[string]string, error) {
	var u models.User
	err := s.db.Where("username = ?", user.Username, user.Password).First(&u).Error
	if err != nil {
		return nil, err
	}
	// Check password
	if !common.CheckPassword(user.Password, u.Password) {
		return nil, errors.New("invalid password")
	}
	// Generate JWT token
	token, err := common.GenerateToken(strconv.Itoa(int(u.ID)), u.Username, u.IsAdmin)
	if err != nil {
		return nil, err
	}
	response := map[string]string{
		"token":     token,
		"username":  u.Username,
		"firstName": u.FirstName,
		"lastName":  u.LastName,
		"email":     u.Email,
	}
	return response, nil
}

func (s *service) UpdateUser(user *models.User) (*models.User, error) {
	return nil, nil
}
