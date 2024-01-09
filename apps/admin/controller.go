package admin

import "github.com/gin-gonic/gin"

func GetUsers(c *gin.Context) {
	service := NewService(Db)
	users, err := service.GetUsers()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, users)
}
