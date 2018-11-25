package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mtojek/aws-closest-region/closest"
)

func main() {
	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "verbose mode")
	flag.Parse()
	serviceName := flag.Arg(0)

	endpoints := new(closest.Endpoints)
	serviceEndpoints, err := endpoints.ForService(serviceName)
	if err != nil {
		log.Fatal(err)
	}

	regions := new(closest.Regions)
	closest, err := regions.FindClosest(serviceEndpoints, verbose)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(closest)
}
