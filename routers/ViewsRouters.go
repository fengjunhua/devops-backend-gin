
package routers

import (
	"github.com/gin-gonic/gin"
	"github/fengjunhua/devops-backend-gin/controllers"
)

func LoadViewsRouters(e *gin.Engine) {

	ViewsRouters := e.Group("/views")
	{
		ViewsRouters.GET("/users", controllers.UsersList)
	}
}