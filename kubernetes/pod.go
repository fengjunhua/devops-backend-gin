package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

//1.查询所有的pods
func GetAllPods(namespace string) []v1.Pod {
	podList, err := K8sClientSet.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}

	for index,node := range podList.Items {
		log.Println(index)
		log.Println(node)
	}

	return podList.Items
}