package routers

import (
	"github.com/gin-gonic/gin"
	"github/fengjunhua/devops-backend-gin/controllers"
)

func LoadLoginRouters(e *gin.Engine) {

	LoginRouters := e.Group("/")
	{
		LoginRouters.GET("/login", controllers.LoginIn)
		LoginRouters.GET("/logout", controllers.LoginOut)
		LoginRouters.GET("/reg", controllers.LoginOut)
	}
}