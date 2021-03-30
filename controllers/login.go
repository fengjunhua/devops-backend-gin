package controllers

import "github.com/gin-gonic/gin"

type LoginController struct{

}



func LoginIn(c *gin.Context) {

	c.HTML(200, "login.html", gin.H{
		"title": "login",
	})


}
func LoginOut(c *gin.Context) {

}

func register(c *gin.Context)  {
	
}