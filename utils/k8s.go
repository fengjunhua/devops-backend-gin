package utils

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func K8sClient() *kubernetes.Clientset {
	var kubeconfig *string
	if home := GetHomePath(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", filepath.Join("", ".kube", "config"), "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}else {
		log.Println("connect k8s success!")
	}
	fmt.Println(clientset)

	return clientset
}

func GetHomePath() string {
	sysType := runtime.GOOS
	fmt.Println(sysType)
	var HomePath string
	if sysType == "linux" {
		HomePath = os.Getenv("HOME")
		fmt.Println(HomePath)
	}
	if sysType == "windows" {
		HomePath = os.Getenv("USERPROFILE")
	}
    return HomePath
}