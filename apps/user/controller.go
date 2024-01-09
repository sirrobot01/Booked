package user

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	var user UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}
	// Validate the input
	if err := user.Validate(); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	service := NewService(Db)
	// Authenticate the user
	u, err := service.Authenticate(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return the user
	c.JSON(200, u)

}
func Register(c *gin.Context) {
	var user UserRegister
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}
	// Validate the input
	if err := user.Validate(); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	service := NewService(Db)
	// Create the user
	u, err := service.CreateUser(user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return the user
	c.JSON(200, u)
}
func Profile(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Profile",
	})
}
