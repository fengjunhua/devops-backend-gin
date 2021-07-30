package jwt

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type login struct {
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}

type User struct {
	UserName  string
	FirstName string
	LastName  string
}

var identityKey = "id"

func helloHandler(c *gin.Context)  {
	claims := jwt.ExtractClaims(c)
	user,_ := c.Get(identityKey)
	c.JSON(200,gin.H{
		"userID":    claims[identityKey],
		"userName":  user.(*User).UserName,
		"text":      "Hello Word.",
	})

}