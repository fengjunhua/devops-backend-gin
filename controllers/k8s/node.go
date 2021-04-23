package k8s

import (
	"fmt"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

//1.对node节点的增删改查操作
func ListNode(c *gin.Context) {

	nodes, err := clientSet.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(nodes)
	c.JSON(http.StatusOK,nodes)
	for _,nds := range nodes.Items {
		fmt.Printf("NodeName: %s\n", nds.Name)
	}

}
func GetNode(c *gin.Context) {
	//nodeName := c.Request.URL.Query().Get("nodeName")
	nodeName := c.Param("nodeName")
	fmt.Println(nodeName)
    nodeInfo, err := clientSet.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
    if err != nil {
        panic(err)
    }
	c.JSON(http.StatusOK, nodeInfo)
    fmt.Println(nodeInfo)

}


/*
2.对namespace进行增删改查相关的操作
*/

func K8sGetPods(*gin.Context) {

}

