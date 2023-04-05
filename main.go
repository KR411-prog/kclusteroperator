package main

import (
	"flag"
	"log"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	klient "github.com/KR411-prog/klusteroperator/pkg/client/clientset/versioned"
)

func main() {
	// k := v1.Kcluster{}
	// fmt.Println(k)
	// //https://github.com/singhggits/ekspose/blob/master/main.go

	// var kubeconfig *string
	// kubeconfig = flag.String("kubeconfig", "/home/.kube/config", "location to your kubeconfig file")
	// flag.Parse()
	// config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	// if err != nil {
	// 	// handle error
	// 	if err != nil {
	// 		log.Printf("Building config from flags %s", err.Error())
	// 	}
	// }

	// klientset,err := klient.NewForConfig(config)
	// if err != nil {
	// 	log.Printf("")
	// }
	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Printf("Building config from flags failed, %s, trying to build inclusterconfig", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			log.Printf("error %s building inclusterconfig", err.Error())
		}
	}
	// for native kubernetes resource, we have client set imported in normal way. but for custom ones, we are importing the clientset versioned as shows in import statements
	klientset, err := klient.NewForConfig(config)
	if err != nil {
		log.Printf("getting klient set %s\n", err.Error())
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("getting std client %s\n", err.Error())
	}

}
