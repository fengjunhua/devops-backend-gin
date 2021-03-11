package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	//c.String(http.StatusOK,result)
	c.HTML(200, "userList.tmpl", gin.H{
		"title": "ss",
	})
}

func GetUsersJSON(c *gin.Context)  {
	c.JSON(
		200,
		gin.H{
				"name": "网站",
				"num": 3,
				"info": "Android",

	})
}