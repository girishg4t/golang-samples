package main

import (
	"flag"
	"fmt"
	"log"
	"os/user"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
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

	clientset, _ := kubernetes.NewForConfig(config)

	fmt.Println("Shared Informer app started")
	factory := informers.NewSharedInformerFactory(clientset, 0)
	informer := factory.Core().V1().Pods().Informer()
	stopper := make(chan struct{})
	defer close(stopper)
	defer runtime.HandleCrash()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAdd,
		DeleteFunc: onDelete,
	})
	go informer.Run(stopper)
	if !cache.WaitForCacheSync(stopper, informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}
	<-stopper

}

func onAdd(obj interface{}) {
	// Cast the obj as node
	node := obj.(*corev1.Pod)
	name := node.GetName()
	if name != "" {
		fmt.Println("Found with node name : ", name)
	}
}

func onDelete(obj interface{}) {
	// Cast the obj as node
	node := obj.(*corev1.Pod)
	name := node.GetName()
	if name != "" {
		fmt.Println("Deleted with pod name : ", name)
	}
}
