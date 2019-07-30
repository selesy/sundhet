package sundhet

import (
	"github.com/blang/semver"
	corev1 "k8s.io/api/core/v1"
	networkingv1beta1 "k8s.io/api/networking/v1beta1"
)

type Kind int

const (
	Batch Kind = iota
	Service
	Stream
	UI
)

var kinds = []struct {
	name  string
	chart string
}{
	{"Batch", ""},
	{"Service", "eio-swe-service"},
	{"Stream", ""},
	{"UI", "angular-client"},
}

func (k Kind) Chart() string {
	return kinds[k].chart
}

func (k Kind) Name() string {
	return kinds[k].name
}

type Instance struct {
	Ingress networkingv1beta1.Ingress
	Service corev1.Service
	Version semver.Version
}

type Application struct {
	Name       string
	Kind       Kind
	Develop    Instance
	Acceptance Instance
	Production Instance
	Features   []Instance
}
