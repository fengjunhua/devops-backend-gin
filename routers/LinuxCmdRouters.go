package routers

import (
	"github.com/gin-gonic/gin"
	"github/fengjunhua/k8smanager/controllers"
)

func LoadLinuxCmdRouters(e *gin.Engine) {

	K8sNodeRouters := e.Group("/api/v1/linux")
	{
		K8sNodeRouters.GET("/ls", controllers.ExecLinuxCommand)
		K8sNodeRouters.GET("/ws", controllers.ExecLinuxWebSocket)
	}
}