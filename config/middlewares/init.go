package middlewares

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CommonMiddlewares(db *gorm.DB) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		CorsMiddleware(db),
	}
}

func AdminMiddlewares(db *gorm.DB) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		AuthMiddleware(db),
		AdminMiddleware(db),
	}
}
