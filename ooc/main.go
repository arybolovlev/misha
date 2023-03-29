/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
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
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
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
			Name: "new",
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
	rc, err = clientset.NodeV1().RuntimeClasses().Get(context.TODO(), o.GetObjectMeta().GetName(), metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Failed to read the Runtime Class object: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Read Runtime Class:", rc)

	// Delete
	err = clientset.NodeV1().RuntimeClasses().Delete(context.TODO(), o.GetObjectMeta().GetName(), metav1.DeleteOptions{})
	if err != nil {
		fmt.Printf("Failed to delete the Runtime Class object: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Deleted Runtime Class")
}
