package closest

import (
	"errors"
	"time"

	log "github.com/sirupsen/logrus"
)

var errEndpointsUnavailable = errors.New("service endpoints are unavailable")

// Regions type selects available endpoints end finds the closest one.
type Regions struct{}

// Latencies type stores average latencies measured while accessing service endpoints.
type Latencies map[string]time.Duration

// FindClosest method finds the closest AWS region to the caller.
func (r *Regions) FindClosest(endpoints Endpoints) (string, error) {
	latencies := Latencies{}
	for regionName, endpoint := range endpoints {
		latency, err := r.measureLatency(endpoint)
		if err == nil {
			latencies[regionName] = latency
		}
	}

	if len(latencies) == 0 {
		return "", errEndpointsUnavailable
	}
	return r.regionWithLowestLatency(latencies), nil
}

func (r *Regions) measureLatency(endpoint string) (time.Duration, error) {
	return 0, nil
}

func (r *Regions) regionWithLowestLatency(latencies Latencies) string {
	var theRegion string
	var theLatency = time.Hour

	for regionName, latency := range latencies {
		if latency < theLatency {
			theRegion = regionName
			theLatency = latency
		}
	}

	log.Infof(`Lowest latency was measured while accessing endpoint in the region "%s": %v`, theRegion, theLatency)
	return theRegion
}
