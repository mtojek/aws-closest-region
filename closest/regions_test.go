package closest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegions_FindClosest_ValidEndpoints(t *testing.T) {
	// given
	regions := new(Regions)
	endpoints := Endpoints{}
	endpoints["slower_endpoint"] = "http://httpbin.org/delay/3"
	endpoints["slow_endpoint"] = "http://httpbin.org/delay/1"
	endpoints["slowest_endpoint"] = "http://httpbin.org/delay/5"

	// when
	region, err := regions.FindClosest(endpoints)

	// then
	assert.Equal(t, "slow_endpoint", region, "faster region should be reported")
	assert.Nil(t, err, "no error should be returned")
}

func TestRegions_FindClosest_BothValidInvalidEndpoints(t *testing.T) {
	// given
	regions := new(Regions)
	endpoints := Endpoints{}
	endpoints["slow_endpoint"] = "http://httpbin.org/delay/1"
	endpoints["wrong_endpoint"] = "http://localhost:57585"

	// when
	region, err := regions.FindClosest(endpoints)

	// then
	assert.Equal(t, "slow_endpoint", region, "faster region should be reported")
	assert.Nil(t, err, "no error should be returned")
}

func TestRegions_FindClosest_InvalidEndpoints(t *testing.T) {
	// given
	regions := new(Regions)
	endpoints := Endpoints{}
	endpoints["wrong_endpoint"] = "http://localhost:69695"

	// when
	region, err := regions.FindClosest(endpoints)

	// then
	assert.NotNil(t, err, "an error should be returned")
	assert.Empty(t, region, "no endpoint should be returned")
}
