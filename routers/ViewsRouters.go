
package routers

import (
	"github.com/gin-gonic/gin"
	"github/fengjunhua/k8smanager/controllers"
)

func LoadViewsRouters(e *gin.Engine) {

	ViewsRouters := e.Group("/views")
	{
		ViewsRouters.GET("/users", controllers.UsersList)
	}
}