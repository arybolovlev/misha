package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	nodev1 "k8s.io/api/node/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// CREATE
	o := &nodev1.RuntimeClass{
		ObjectMeta: metav1.ObjectMeta{
			// Name: "runtime-class",
			GenerateName: "runtime-class-",
		},
		Handler: "runc",
	}
	rc, err := clientset.NodeV1().RuntimeClasses().Create(context.TODO(), o, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Failed to create a new Runtime Class object: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Created Runtime Class:", rc)

	// Read
	rcRead, err := clientset.NodeV1().RuntimeClasses().Get(context.TODO(), rc.Name, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Failed to read the Runtime Class object: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Read Runtime Class:", rcRead)

	// Delete
	err = clientset.NodeV1().RuntimeClasses().Delete(context.TODO(), rc.Name, metav1.DeleteOptions{})
	if err != nil {
		fmt.Printf("Failed to delete the Runtime Class object: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Deleted Runtime Class")
}
