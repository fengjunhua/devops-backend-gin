package routers

import (
	"github.com/gin-gonic/gin"
	"github/fengjunhua/devops-backend-gin/controllers"
)

func InitLinuxRouter(Router *gin.Engine) {

	LinuxRouter := Router.Group("linux")
	{
		LinuxRouter.GET("/ls", controllers.ExecLinuxCommand)
		LinuxRouter.GET("/ws", controllers.ExecLinuxWebSocket)
	}
}