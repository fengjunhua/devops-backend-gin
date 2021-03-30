package routers

import (
	"github.com/gin-gonic/gin"
	"github/fengjunhua/k8smanager/controllers"
)

func LoadK8sRouters(e *gin.Engine) {

	K8sNodeRouters := e.Group("/api/v1/admin/node")
	{
		K8sNodeRouters.GET("/list",controllers.K8sListNode)
		K8sNodeRouters.GET("/node/:nodeName",controllers.K8sGetNode)
	}

	K8sNamespaceRouters := e.Group("/api/v1/admin/namespace")
	{
		K8sNamespaceRouters.GET("/")

	}
	K8sServiceRouters := e.Group("/api/v1/admin/service")
	{
		K8sServiceRouters.GET("/pods",controllers.K8sGetPods)

	}
	K8sPodRouters := e.Group("/api/v1/admin/pod")
	{
		K8sPodRouters.GET("/pods",controllers.K8sGetPods)

	}


}
/*
func getNodes(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{"nodes":8})

}
func getNamespaces(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"namespaces":8})
}
func getServices(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"services":8})
}
func getPods(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"pods":8})
}
 */