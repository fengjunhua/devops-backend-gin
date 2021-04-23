package k8s

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
	clientSet *kubernetes.Clientset

)

func init()  {
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
	clientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Println(err)
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