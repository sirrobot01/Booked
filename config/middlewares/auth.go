package middlewares

import (
	"booked/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Missing authorization header"})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, Bearer)
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Missing token"})
			return
		}
		token, err := jwt.ParseWithClaims(tokenString, &common.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("SECRET"), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}
		claims, ok := token.Claims.(*common.JwtClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}
		//Validate
		if err := claims.Valid(); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token expired"})
			return
		}
		if err := common.ValidateToken(claims, db); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}
		c.Set("user", token)
		c.Set("username", claims.Username)
		c.Set("user_id", claims.UserID)
		c.Set("admin", claims.Admin)
		c.Next()
	}
}

func AdminMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		admin := c.GetBool("admin")
		if !admin {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		c.Next()
	}
}
