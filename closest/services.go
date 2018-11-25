package closest

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws/endpoints"
	log "github.com/sirupsen/logrus"
)

const (
	defaultServiceName = "dynamodb"
	pingURLSuffix      = "/ping"
)

var errServiceNotAvailableInAnyRegion = errors.New("service is not available in any region")

// Services type provides a list of supported AWS regions.
type Services struct{}

// Endpoints is a map with region name as a key and an endpoint as value.
type Endpoints map[string]string

// EndpointsForService method provides a list of endpoints for a given service.
// US-Gov partition will be skipped.
func (s *Services) EndpointsForService(serviceName string) (Endpoints, error) {
	serviceName = s.serviceNameOrDefault(serviceName)

	var anyRegionExists bool
	serviceEndpoints := Endpoints{}
	for _, partition := range endpoints.DefaultPartitions() {
		if partition.ID() == endpoints.AwsUsGovPartition().ID() {
			log.Info(`Partition "us-gov" will be skipped.`)
			continue
		}

		serviceRegions, exists := endpoints.RegionsForService(endpoints.DefaultPartitions(),
			partition.ID(), serviceName)
		log.Infof(`Service "%s" is available in %d regions in "%s" partition.`, serviceName,
			len(serviceRegions), partition.ID())
		anyRegionExists = anyRegionExists || exists

		if exists {
			for regionName := range serviceRegions {
				serviceEndpoint, err := partition.EndpointFor(serviceName, regionName)
				if err != nil {
					return nil, err
				}
				serviceEndpoints[regionName] = serviceEndpoint.URL + pingURLSuffix
			}
		}
	}

	if !anyRegionExists {
		return nil, errServiceNotAvailableInAnyRegion
	}

	if log.IsLevelEnabled(log.InfoLevel) {
		log.Infoln("Service is accessing via following endpoints:")
		for regionName, endpoint := range serviceEndpoints {
			log.Infof("  %s: %s\n", regionName, endpoint)
		}
	}

	return serviceEndpoints, nil
}

func (s *Services) serviceNameOrDefault(serviceName string) string {
	if serviceName == "" {
		log.Infof("Service name hasn't been provided. Use default service name: %s", defaultServiceName)
		return defaultServiceName
	}
	return serviceName
}
