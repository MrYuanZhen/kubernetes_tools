package main

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "net/http/pprof"
)


func getNodeIP(clientset *kubernetes.Clientset) []string {

	//get controller node list
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{LabelSelector:"node-role.kubernetes.io/controlplane=true"})
	if err != nil {
		panic(err)
	}

	var ipArr = make([]string, len(nodes.Items))

	for i := 0; i < len(nodes.Items); i++ {
		ipArr[i] = nodes.Items[i].Status.Addresses[0].Address
	}

	//return controller node ip
	return ipArr
}
