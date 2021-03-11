package controllers

import "github.com/gin-gonic/gin"

func Index(c *gin.Context)  {
	c.HTML(200, "index.tmpl",gin.H{
		"name" : "标题",
	})
}

func UsersList(c *gin.Context) {
	//c.String(http.StatusOK,result)
	c.HTML(200, "userList.tmpl", gin.H{
		"title": "ss",
	})
}