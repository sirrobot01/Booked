package common

import (
	"booked/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type JwtClaims struct {
	Username string `json:"username"`
	UserID   string `json:"user_id"`
	Admin    bool   `json:"admin"`
	jwt.StandardClaims
}

func GenerateToken(userId string, username string, admin bool) (string, error) {
	expirationTime := time.Now().Add(100 * time.Minute)
	claims := JwtClaims{
		Username: username,
		Admin:    admin,
		UserID:   userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "booked",
			Subject:   userId,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("SECRET"))
}

func ValidateToken(claims jwt.Claims, db *gorm.DB) error {
	userId := (claims).(*JwtClaims).UserID
	var u models.User
	return db.Where("id = ?", userId).First(&u).Error
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(username string, firstName string, lastName string,
	email string, password string, isAdmin bool, isActive bool) *models.User {
	password, _ = HashPassword(password)

	return &models.User{
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		IsActive:  isActive,
		IsAdmin:   isAdmin,
	}
}
