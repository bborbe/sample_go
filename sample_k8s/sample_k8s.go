package main

import (
	"flag"
	"fmt"
	"path/filepath"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"strings"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	printDeployments(clientset, err)

	printIngress(clientset, err)

}
func printIngress(clientset *kubernetes.Clientset, err error) {
	ingressClient := clientset.ExtensionsV1beta1().Ingresses(apiv1.NamespaceAll)
	ingressList, err := ingressClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ingress\n")
	for _, d := range ingressList.Items {
		var hosts []string
		for _, rule := range d.Spec.Rules {
			hosts = append(hosts, rule.Host)
		}
		fmt.Printf(" * %s (%s)\n", d.Name, strings.Join(hosts, ","))
	}
}
func printDeployments(clientset *kubernetes.Clientset, err error) {
	deploymentsClient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceAll)
	deploymentList, err := deploymentsClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deployments\n")
	for _, d := range deploymentList.Items {
		fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
	}
}
