package main

import (
	"context"
	"fmt"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespace, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		panic(err.Error())
	}

	for {
		pods, err := client.CoreV1().Pods(string(namespace)).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("There are %d pods in the namespace %s\n", len(pods.Items), string(namespace))
		}

		time.Sleep(10 * time.Second)
	}
}
