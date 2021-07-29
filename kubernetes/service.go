package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

//1.查询所有的services
func GetAllServices(namespace string) []v1.Service {
	serviceList, err := K8sClientSet.CoreV1().Services(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}

	for index,service := range serviceList.Items {
		log.Println(index)
		log.Println(service)
	}

	return serviceList.Items
}