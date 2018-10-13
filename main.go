package main

import (
	"flag"
	"time"

	"github.com/golang/glog"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	clientset "github.com/fajran/kuncen/pkg/client/clientset/versioned"
	informers "github.com/fajran/kuncen/pkg/client/informers/externalversions"
	"github.com/fajran/kuncen/pkg/controller"
	"github.com/fajran/kuncen/pkg/signals"
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

	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	client, err := clientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building kuncen clientset: %s", err.Error())
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	informerFactory := informers.NewSharedInformerFactory(client, time.Second*30)

	ctrl := controller.NewController(kubeClient, client,
		kubeInformerFactory.Core().V1().Secrets(),
		informerFactory.Kuncen().V1alpha1().EncryptedSecrets())

	go kubeInformerFactory.Start(stopCh)
	go informerFactory.Start(stopCh)

	glog.Infof("Running the controller..")

	err = ctrl.Run(2, stopCh)
	if err != nil {
		glog.Fatalf("Error running controller: %s", err.Error())
	}
}
