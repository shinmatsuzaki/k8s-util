package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	// GCPのauth moduleを使いたいのでimport
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {

	// kubeconfigのPATHの組み立て
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	// flagの内容をpraseし、^で定義の変数に格納
	flag.Parse()

	// BuildConfigFromFlagsでconfigにkubeconfigの内容を格納
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// NewForConfig() でconfigを元にClientsetを生成
	// https://godoc.org/k8s.io/client-go/kubernetes#NewForConfig
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// kind: podの一覧を nodes に格納
	// https://godoc.org/k8s.io/client-go/kubernetes/typed/core/v1

	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get pods:", err)
	}

	for i, node := range pods.Items {
		fmt.Printf("[%d] %s\n", i, node.GetName())
	}
}
