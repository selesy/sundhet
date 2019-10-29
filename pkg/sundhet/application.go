package sundhet

import (
	"github.com/blang/semver"
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
	{"Batch", "eio-swe-cronjob"},
	{"Service", "eio-swe-service"},
	{"Stream", "eio-swe-service"},
	{"UI", "angular-client"},
}

func (k Kind) Chart() string {
	return kinds[k].chart
}

func (k Kind) Name() string {
	return kinds[k].name
}

type DeploymentVersion struct {
	Suffix  string
	Version semver.Version
	Containers
}

type PodVersion

type Application struct {
	Name       string
	Kind       Kind
	Develop    Instance
	Acceptance Instance
	Production Instance
	Features   []Instance
}
