package routers

import (
	"github.com/gin-gonic/gin"
	"github/fengjunhua/devops-backend-gin/controllers"
)

func InitSshRouter(Router *gin.RouterGroup) {
	SshRouter := Router.Group("ssh")
	{
		SshRouter.GET("/ls", controllers.ExecLinuxCommand)
		SshRouter.GET("/ws", controllers.ExecLinuxWebSocket)

	}
}

