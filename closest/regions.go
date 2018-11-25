package closest

import (
	"errors"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const measureRepeats = 5

var (
	errAllEndpointsUnavailable            = errors.New("all service endpoints are unavailable")
	errRegionalServiceEndpointUnavailable = errors.New("regional service endpoint is unavailable")

	httpClient = http.Client{Timeout: 30 * time.Second}
)

// Regions type selects available endpoints end finds the closest one.
type Regions struct{}

// Latencies type stores average latencies measured while accessing service endpoints.
type Latencies map[string]time.Duration

// FindClosest method finds the closest AWS region to the caller.
func (r *Regions) FindClosest(endpoints Endpoints) (string, error) {
	log.Info("Average latencies:")

	latencies := Latencies{}
	for regionName, endpoint := range endpoints {
		latency, err := r.measureLatency(endpoint)
		if err == nil {
			log.Infof("  %s: %v", regionName, latency)
			latencies[regionName] = latency
		} else {
			log.Infof("  %s: invalid measure", regionName)
		}
	}

	if len(latencies) == 0 {
		return "", errAllEndpointsUnavailable
	}

	theRegion, theLatency := r.regionWithLowestLatency(latencies)
	log.Infof(`Lowest latency was measured while accessing endpoint in the region "%s": %v`,
		theRegion, theLatency)
	return theRegion, nil
}

func (r *Regions) measureLatency(endpoint string) (time.Duration, error) {
	var c int64
	var sum int64

	for i := 0; i < measureRepeats; i++ {
		startTime := time.Now()
		response, err := httpClient.Get(endpoint)
		if err == nil {
			sum += int64(time.Now().Sub(startTime))
			c++
		} else {
			log.Errorf(`Error while accessing endpoint "%s": %v`, endpoint, err)
		}

		if response != nil && response.Body != nil {
			response.Body.Close()
		}
	}

	if c == 0 {
		return -1, errRegionalServiceEndpointUnavailable
	}
	return time.Duration(sum / c), nil
}

func (r *Regions) regionWithLowestLatency(latencies Latencies) (string, time.Duration) {
	var theRegion string
	var theLatency = time.Hour

	for regionName, latency := range latencies {
		if latency < theLatency {
			theRegion = regionName
			theLatency = latency
		}
	}
	return theRegion, theLatency
}
