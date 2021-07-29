package kubernetes

import (
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

//1.查询所有的deployment
func GetAllDeployments(namespace string) []v1.Deployment {
	deploymentList, err := K8sClientSet.AppsV1().Deployments(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}


	for index,deployment := range deploymentList.Items{
		log.Println(index)
		log.Println(deployment)
	}
	return deploymentList.Items
}

//2.查询指定deployment的详细信息
func GetDeploymentByName(namespace string,name string) *v1.Deployment {

	deployment, err := K8sClientSet.AppsV1().Deployments(namespace).Get(name,metav1.GetOptions{})
	if err != nil {
		log.Println(err)
	}
	log.Println(deployment.Name)
	log.Println(deployment.Status)
	log.Println(deployment.Spec)

	return deployment
}