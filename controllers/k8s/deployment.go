package k8s

import (
	"fmt"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"net/http"
)

type Deployment struct {
	Name string
	Replicas int32
	Images string
}

//1.对deployment节点的增删改查操作
func ListDeployments(c *gin.Context) {
	deploymentList, err := clientSet.AppsV1().Deployments("mec").List(metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(deploymentList)
	//c.JSON(http.StatusOK,deploymentList)
	var ret []*Deployment

	for _,item := range deploymentList.Items{
		ret = append(ret,&Deployment{Name:item.Name})
		fmt.Println(item.Name)
	}
	log.Println(&ret)
	c.JSON(
		http.StatusOK,
		deploymentList,
	)

}

//2.查询指定deployment的详细信息
func GetDeploymentsByName(c *gin.Context) {
	namespace := c.Param("namespace")
	deploymentName := c.Param("deploymentName")
	deployment, err := clientSet.AppsV1().Deployments(namespace).Get(deploymentName,metav1.GetOptions{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(deployment)
	c.JSON(http.StatusOK,deployment.Name)
}