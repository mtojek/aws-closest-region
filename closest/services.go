package closest

// Services type provides a list of supported AWS regions.
type Services struct{}

// Endpoints is a map with region name as a key and an endpoint as value.
type Endpoints map[string]string

// ForService method provides a list of endpoints for a given service.
func (e *Endpoints) ForService(serviceName string) (Endpoints, error) {
	return nil, nil
}
