package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/selesy/sundhet/pkg/sundhet"
	log "github.com/sirupsen/logrus"
)

const kubernetes = `kubernetes`

type Cfg struct {
}

func main() {
	var cfg Cfg
	err := envconfig.Process(kubernetes, &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	sundhet.Something("k8s-qa-b")
	sundhet.Something("k8s-prod-b")
}
