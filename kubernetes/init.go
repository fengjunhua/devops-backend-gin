package kubernetes

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

var(
	kubeConfig *string
	K8sClientSet *kubernetes.Clientset
)

func InitK8sClient()  {
	if home := GetHomePath(); home != "" {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeConfig file")
	} else {
		kubeConfig = flag.String("kubeConfig", filepath.Join("", ".kube", "config"), "absolute path to the kubeConfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		log.Println(err)
	}

	var err2 error
	K8sClientSet, err2 = kubernetes.NewForConfig(config)
	if err2 != nil {
		log.Println(err2)
	}else {
		log.Println("connect admin success!")
	}

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