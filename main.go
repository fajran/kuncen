package main

import (
	"flag"

	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/fajran/kuncen/pkg/apis/kuncen/v1alpha1"
	clientset "github.com/fajran/kuncen/pkg/client/clientset/versioned"
)

var (
	masterURL  string
	kubeconfig string
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}

func main() {
	flag.Parse()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	client, err := clientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building kuncen clientset: %s", err.Error())
	}

	opts := v1.ListOptions{}
	watch, err := client.KuncenV1alpha1().EncryptedSecrets("").Watch(opts)
	if err != nil {
		glog.Fatalf("Cannot watch resource: %s", err.Error())
	}

	glog.Info("Watching resources")
	for event := range watch.ResultChan() {
		obj, ok := event.Object.(*v1alpha1.EncryptedSecret)
		if !ok {
			continue
		}

		glog.Infof("Watch event, type:%s, namespace: %s, name: %s",
			event.Type, obj.Namespace, obj.Name)
	}

}
