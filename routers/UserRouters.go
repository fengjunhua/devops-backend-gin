package routers

import (
	"github.com/gin-gonic/gin"
	"github/fengjunhua/devops-backend-gin/controllers"
)

func LoadUsersRouters(e *gin.Engine) {

	UserRouters := e.Group("/api/v1/users")
	{
		UserRouters.GET("/ls", controllers.GetUsers)
		UserRouters.GET("/user.json", controllers.GetUsersJSON)
	}
}