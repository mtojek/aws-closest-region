package closest

// Regions type selects available endpoints end finds the closest one.
type Regions struct{}

// FindClosest method finds the closest AWS region to the caller.
func (r *Regions) FindClosest(endpoints Endpoints, verbose bool) (string, error) {
	return "us-west-66", nil
}
