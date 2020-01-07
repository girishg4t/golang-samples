package main

import (
	"fmt"

	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8type "k8s.io/client-go/kubernetes/typed/apps/v1"
)

func printPods(cwg ClientWg) {
	var pods *corev1.PodList
	go func() {
		// access the API to list pods
		pods, _ = cwg.clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		for _, val := range pods.Items {
			if val.Status.Phase == corev1.PodRunning {
				fmt.Println("Pods are =>", val.GetName(), val.Status.Phase)
			}
		}
		defer cwg.wg.Done()
	}()
}

func printDeployments(cwg ClientWg) {

	var deployments *appv1.DeploymentList

	go func() {
		deployments, _ = cwg.clientset.AppsV1().Deployments(metav1.NamespaceDefault).List(metav1.ListOptions{})
		fmt.Printf("There are %d deployments in the development namespace\n", len(deployments.Items))

		for _, val := range deployments.Items {
			fmt.Println("Deployments are =>", val.GetName())
		}
		defer cwg.wg.Done()
	}()
}

func checkPodinNamespace(pod, ns string, cwg ClientWg) {
	go func() {
		_, err := cwg.clientset.CoreV1().Pods(ns).Get(pod, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod %s in namespace %s not found\n", pod, ns)
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %s in namespace %s: %v\n",
				pod, ns, statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found pod %s in namespace %s\n", pod, ns)
		}
		defer cwg.wg.Done()
	}()
}

//CreateDeployment for creating the deployment
func createDeployment(cwg ClientWg, name string) {
	go func() {
		deploymentsClient := cwg.clientset.AppsV1().Deployments(metav1.NamespaceDefault)
		found := checkNodeAlreadyPresent(deploymentsClient, name)
		if found {
			defer cwg.wg.Done()
			return			
		}
		deployment := &appv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
			Spec: appv1.DeploymentSpec{
				Replicas: int32Ptr(1),
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"app": "demo",
					},
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"app": "demo",
						},
					},
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							{
								Name:  "web",
								Image: "nginx:1.12",
								Ports: []corev1.ContainerPort{
									{
										Name:          "http",
										Protocol:      corev1.ProtocolTCP,
										ContainerPort: 80,
									},
								},
							},
						},
					},
				},
			},
		}

		// Create Deployment
		fmt.Println("Creating deployment...")
		result, err := deploymentsClient.Create(deployment)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
		defer cwg.wg.Done()
	}()
}

func deletePod(name string, cwg ClientWg) {
	go func() {
		deploymentsClient := cwg.clientset.AppsV1().Deployments(metav1.NamespaceDefault)
		found := checkNodeAlreadyPresent(deploymentsClient, name)
		if !found {
			defer cwg.wg.Done()
			return			
		}
		options := &metav1.DeleteOptions{}
		err := deploymentsClient.Delete(name, options)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Deleted deployment %q.\n", name)
		defer cwg.wg.Done()
	}()
}

func checkNodeAlreadyPresent(deploymentsClient interface{}, name string) bool {
	if v, ok := deploymentsClient.(k8type.DeploymentInterface); ok {
		existingDeployment, _ := v.List(metav1.ListOptions{})
		for _, el := range existingDeployment.Items {
			if el.GetName() == name {
				fmt.Println("Deployment demo-deployment already present")
				return true
			}
		}
	}	
	return false
}

func int32Ptr(i int32) *int32 { return &i }
