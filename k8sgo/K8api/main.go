package main

import (
	"flag"
	"fmt"
	"log"
	"os/user"
	"path/filepath"
	"sync"
	"time"

	"k8s.io/client-go/kubernetes"
	k8 "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//ClientWg for k8 client
type ClientWg struct {
	clientset *k8.Clientset
	wg        *sync.WaitGroup
}

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	var kubeconfig *string
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	if home := usr.HomeDir; home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, _ := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	// creates the clientset
	var clientset *k8.Clientset
	clientset, _ = kubernetes.NewForConfig(config)

	cwg := ClientWg{clientset, &wg}

	wg.Add(1)
	printDeployments(cwg)

	wg.Add(1)
	printPods(cwg)

	namespace := "default"
	pod := "hello-foo-59f76b8655-fbll9"
	wg.Add(1)
	checkPodinNamespace(pod, namespace, cwg)

	wg.Add(1)
	createDeployment(cwg, "demo-deployment6")

	wg.Add(1)
	deletePod("demo-deployment6", cwg)

	t := time.Now()
	elapsed := t.Sub(start)
	wg.Wait()
	fmt.Println("Time taken =>", elapsed)
}
