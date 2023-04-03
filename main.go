package main

import (
	"flag"
	"fmt"
	"log"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// k := v1.Kcluster{}
	// fmt.Println(k)
	// //https://github.com/viveksinghggits/ekspose/blob/master/main.go

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
	fmt.Println("test")
}
