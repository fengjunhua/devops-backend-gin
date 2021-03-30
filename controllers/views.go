package controllers

import "github.com/gin-gonic/gin"


func UsersList(c *gin.Context) {
	//c.String(http.StatusOK,result)
	c.HTML(200, "userList.tmpl", gin.H{
		"title": "ss",
	})
}