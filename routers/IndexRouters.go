package routers

import (
	"github.com/gin-gonic/gin"
	"github/fengjunhua/k8smanager/controllers"
)

func LoadIndexRouters(e *gin.Engine) {

	IndexRouters := e.Group("/")
	{
		IndexRouters.GET("/index.html", controllers.Index)
	}
}