package config

import (
	"booked/apps/admin"
	"booked/apps/user"
	"booked/config/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CommonMiddlewares(db)...)

	user.Routes(r.Group("/user"), db)
	admin.Routes(r.Group("/admin"), db)
	// payment.Routes(r.Group("/payment"))
	return r

}
