package admin

import (
	"booked/config/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Routes(group *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	Db = db

	group.Use(middlewares.AdminMiddlewares(db)...)

	group.GET("/users", GetUsers)

	return group
}
