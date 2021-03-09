package routers

import (
	"github.com/gin-gonic/gin"
	"github/fengjunhua/k8smanager/controllers"
)

func LoadTagRouters(e *gin.Engine)  {
	TagRouters := e.Group("/api/v1")
	{

		TagRouters.GET("/tags",controllers.Get)

	}

}