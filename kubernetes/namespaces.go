package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

//1.查询所有的namespaces
func GetAllNamespaces() []v1.Namespace {

	namespaceList, err := K8sClientSet.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}

	for index,namespace := range namespaceList.Items{
		log.Println(index)
		log.Println(namespace)
	}

	return namespaceList.Items
}