package routers

import (
	"github.com/gin-gonic/gin"
	"github/fengjunhua/devops-backend-gin/controllers"
)

func LoadIndexRouters(e *gin.Engine) {

	IndexRouters := e.Group("/")
	{
		IndexRouters.GET("/index.html", controllers.Index)
	}
}