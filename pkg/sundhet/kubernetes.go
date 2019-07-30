package sundhet

import (
	"flag"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Something() {
	//cfg, err := rest.InClusterConfig()
	var kubeconfig *string
	home := "/home/swm16"
	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	flag.Parse()

	// use the current context in kubeconfig
	cfg, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	cs, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err.Error())
	}

	var _ networkingv1beta1.IngressList

	var opts metav1.ListOptions
	nses, err := cs.CoreV1().Namespaces().List(opts)
	for _, ns := range nses.Items {
		log.Info("Namespace: ", ns)
	}

	ingresses, err := cs.ExtensionsV1beta1().Ingresses("eio-swe").List(opts)
	if err != nil {
		log.Error(err.Error())
	}
	for _, i := range ingresses.Items {
		log.Info("Ingress --- Name: ", i.Name, ", Chart: ", i.GetObjectMeta().GetLabels()["chart"])
	}

	cronjobs, err := cs.BatchV1beta1().CronJobs("eio-swe").List(opts)
	if err != nil {
		log.Error(err.Error())
	}
	for _, i := range cronjobs.Items {
		log.Info("CronJobs --- Name: ", i.Name, ", Chart: ", i.GetObjectMeta().GetLabels()["chart"])
	}

}
