package closest

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServices_EndpointsForService_EmptyServiceName(t *testing.T) {
	// given
	serviceName := ""
	services := new(Services)

	// when
	endpoints, err := services.EndpointsForService(serviceName)

	// then
	assert.Nil(t, err, "no errors should be returned")
	assert.True(t, len(endpoints) > 0, "at least one endpoint should be returned")
	assert.Contains(t, endpoints, "us-west-2")
	assert.NotContains(t, endpoints, "us-west-999")
}

func TestServices_EndpointsForService_ServiceNamePassed(t *testing.T) {
	// given
	serviceName := "polly"
	services := new(Services)

	// when
	endpoints, err := services.EndpointsForService(serviceName)

	// then
	assert.Nil(t, err, "no errors should be returned")
	assert.True(t, len(endpoints) > 0, "at least one endpoint should be returned")
	assert.Contains(t, endpoints, "us-west-2")
	assert.NotContains(t, endpoints, "us-west-999")

	for _, endpoint := range endpoints {
		assert.True(t, strings.HasSuffix(endpoint, pingURLSuffix), "should contain ping suffix")
	}
}

func TestServices_EndpointsForService_UnknownServiceName(t *testing.T) {
	// given
	serviceName := "unknown"
	services := new(Services)

	// when
	endpoints, err := services.EndpointsForService(serviceName)

	// then
	assert.NotNil(t, err, "an error should be returned")
	assert.Nil(t, endpoints, "no endpoints should be returned")
}

func TestServices_EndpointsForService_WithChinaPartition(t *testing.T) {
	// given
	serviceName := "dynamodb"
	services := new(Services)

	// when
	endpoints, err := services.EndpointsForService(serviceName)

	// then
	assert.Nil(t, err, "no errors should be returned")
	assert.True(t, len(endpoints) > 0, "at least one endpoint should be returned")
	assert.Contains(t, endpoints, "us-west-2")
	assert.Contains(t, endpoints, "cn-north-1")
	assert.Contains(t, endpoints, "cn-northwest-1")
	assert.NotContains(t, endpoints, "us-west-999")

	for _, endpoint := range endpoints {
		assert.True(t, strings.HasSuffix(endpoint, pingURLSuffix), "should contain ping suffix")
	}
}
