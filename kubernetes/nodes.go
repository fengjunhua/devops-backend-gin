package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

//1.查询所有的nodes
func GetAllNodes() []v1.Node {
	nodeList, err := K8sClientSet.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}

	for index,node := range nodeList.Items {
		log.Println(index)
		log.Println(node)
	}

	return nodeList.Items
}


//2.对node节点的增删改查操作
func GetNode(nodeName string) {

	node, err := K8sClientSet.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
	}
	log.Println(node.Status)
	log.Println(node.Spec)

}



