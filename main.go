package main

import (
	"flag"
	"fmt"

	"github.com/mtojek/aws-closest-region/closest"
	log "github.com/sirupsen/logrus"
)

func main() {
	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "verbose mode")
	flag.Parse()
	serviceName := flag.Arg(0)

	log.SetFormatter(&log.TextFormatter{
		DisableLevelTruncation: true,
		DisableTimestamp:       true,
	})
	if verbose {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

	services := new(closest.Services)
	serviceEndpoints, err := services.EndpointsForService(serviceName)
	if err != nil {
		log.Fatal(err)
	}

	regions := new(closest.Regions)
	closestEndpoint, err := regions.FindClosest(serviceEndpoints)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(closestEndpoint)
}
