package cmd

// Service is a cmd service
type Service struct{}

// New allocates the memories for Service
func New() *Service {
	return &Service{}
}
