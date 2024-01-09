package user

import (
	"booked/config/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Routes(group *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	Db = db
	group.POST("/login", Login)
	group.POST("/register", Register)
	group.GET("/profile", middlewares.AuthMiddleware(db), Profile)
	return group
}
