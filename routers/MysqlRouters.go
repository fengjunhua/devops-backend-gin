package routers

import (
	"github.com/gin-gonic/gin"
	"github/fengjunhua/k8smanager/database/mysql"
)

func LoadMysqlRouters(e *gin.Engine) {

	K8sNodeRouters := e.Group("/api/v1/mysql")
	{
		K8sNodeRouters.GET("/get", mysql.GetUserById)
		//K8sNodeRouters.GET("/post", controllers.K8sGetNode)
	}
}
