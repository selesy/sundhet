package sundhet

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func Something(kontext string) {
	// //cfg, err := rest.InClusterConfig()
	// var kubeconfig *string
	// home := "/home/swm16"
	// kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// flag.Parse()

	// // use the current context in kubeconfig
	// cfg, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// log.Info("Kfg: ", cfg)

	_, cs, err := GetKubeClient(kontext)
	//cs, err := kubernetes.NewForConfig(cfg)
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

	services, err := cs.CoreV1().Services("eio-swe").List(opts)
	if err != nil {
		log.Error(err.Error())
	}
	for _, i := range services.Items {
		log.Info("Service --- Name: ", i.Name, ", Chart: ", i.GetObjectMeta().GetLabels()["chart"])
	}

	cronjobs, err := cs.BatchV1beta1().CronJobs("eio-swe").List(opts)
	if err != nil {
		log.Error(err.Error())
	}
	for _, i := range cronjobs.Items {
		log.Info("CronJobs --- Name: ", i.Name, ", Chart: ", i.GetObjectMeta().GetLabels()["chart"])
	}

	pods, err := cs.CoreV1().Pods("eio-swe").List(opts)
	if err != nil {
		log.Error(err.Error())
	}
	for _, i := range pods.Items {
		log.Info("Pods --- Name: ", i.Name, ", Chart: ", i.GetObjectMeta().GetLabels()["chart"])
	}

}

// GetKubeClient creates a Kubernetes config and client for a given kubeconfig context.
func GetKubeClient(context string) (*rest.Config, kubernetes.Interface, error) {
	config, err := configForContext(context)
	if err != nil {
		return nil, nil, err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, fmt.Errorf("could not get Kubernetes client: %s", err)
	}
	return config, client, nil
}

// configForContext creates a Kubernetes REST client configuration for a given kubeconfig context.
func configForContext(context string) (*rest.Config, error) {
	config, err := getConfig(context).ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("could not get Kubernetes config for context %q: %s", context, err)
	}
	return config, nil
}

// getConfig returns a Kubernetes client config for a given context.
func getConfig(context string) clientcmd.ClientConfig {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	rules.DefaultClientConfig = &clientcmd.DefaultClientConfig

	overrides := &clientcmd.ConfigOverrides{ClusterDefaults: clientcmd.ClusterDefaults}

	if context != "" {
		overrides.CurrentContext = context
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, overrides)
}
