package main

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "net/http/pprof"
)


func getNodeIP(clientset *kubernetes.Clientset) []string {

	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	var ipArr = make([]string, len(nodes.Items))

	for i := 0; i < len(nodes.Items); i++ {
		ipArr[i] = nodes.Items[i].Status.Addresses[0].Address
	}

	return ipArr
}
