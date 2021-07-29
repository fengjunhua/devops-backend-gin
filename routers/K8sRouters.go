package routers

import (
	"github.com/gin-gonic/gin"
	"github/fengjunhua/devops-backend-gin/controllers/k8s"

)


func LoadK8sRouters(e *gin.Engine) {

	K8sNamespaceRouters := e.Group("/api/v1/admin/namespace")
	{
		K8sNamespaceRouters.GET("/")

	}
	K8sNodeRouters := e.Group("/api/v1/k8s/node")
	{
		K8sNodeRouters.GET("/list", k8s.GetAllNodes)
		K8sNodeRouters.GET("/node/:nodeName", k8s.GetNode)
	}

	K8sDeploymentRouters := e.Group("/api/v1/k8s/deployment")
	{
		K8sDeploymentRouters.GET("/list", k8s.GetAllDeployments)
		K8sDeploymentRouters.GET("/namespace/:namespace/deploymentName/:deploymentName", k8s.GetDeployment)
	}
	K8sServiceRouters := e.Group("/api/v1/admin/service")
	{
		K8sServiceRouters.GET("/pods", k8s.GetAllPods)

	}
	K8sPodRouters := e.Group("/api/v1/admin/pod")
	{
		K8sPodRouters.GET("/pods", k8s.GetAllServices)

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