package main

import (
	"fmt"
	"strconv"
	"time"

	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"

	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

type service struct {
	name      string
	stype     corev1.ServiceType
	port      int32
	selector  map[string]string
	namespace string
}

var cs *kubernetes.Clientset
var serviceName string
var namespace string = "mynamespace"

func main() {
	cs = getKubeHandle()

	watchlist := cache.NewListWatchFromClient(cs.AppsV1().RESTClient(), "deployments", namespace,
		fields.Everything())
	_, controller := cache.NewInformer(
		watchlist,
		&v1.Deployment{},
		time.Second*0,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    onAdd,
			DeleteFunc: onDelete,
		},
	)
	stop := make(chan struct{})
	go controller.Run(stop)
	for {
		time.Sleep(time.Second)
	}
}

func onAdd(obj interface{}) {
	dep := obj.(*v1.Deployment)
	name := dep.GetName()
	if name == "" {
		return
	}

	metadata := dep.Spec.Template.ObjectMeta
	spec := dep.Spec.Template.Spec

	isServiceReq, _ := strconv.ParseBool(metadata.Annotations["auto-create-svc"])
	if !isServiceReq {
		return
	}

	serviceName = name + "-service"
	annServiceType := metadata.Annotations["auto-create-svc-type"]
	serviceType := corev1.ServiceTypeClusterIP
	if annServiceType == "NodePort" {
		serviceType = corev1.ServiceTypeNodePort
	}
	port := int32(8080)
	if spec.Containers[0].Ports != nil && len(spec.Containers[0].Ports) > 0 {
		port = spec.Containers[0].Ports[0].ContainerPort
	}

	service := &service{namespace: namespace,
		name:  serviceName,
		port:  port,
		stype: serviceType,
		selector: map[string]string{
			"app": "demo-app",
		}}
	createService(service)
}

func onDelete(obj interface{}) {
	dep := obj.(*v1.Deployment)
	name := dep.GetName()
	if name != "" {
		ds := cs.CoreV1().Services(namespace)

		options := &metav1.DeleteOptions{}
		err := ds.Delete(serviceName, options)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Deleted service %q.\n", name)
	}
}

func createService(s *service) {
	sc := cs.CoreV1().Services(namespace)
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceName,
			Namespace: namespace,
			Labels: map[string]string{
				"app": "demo-app",
			},
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       s.port,
					TargetPort: intstr.FromInt(int(s.port)),
				},
			},
			Selector: s.selector,
			Type:     s.stype,
		},
	}

	// Create Service
	fmt.Println("Creating service...")
	result, err := sc.Create(service)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created service %q.\n", result.GetObjectMeta().GetName())
}

func int32Ptr(i int32) *int32 { return &i }
