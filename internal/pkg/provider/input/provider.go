package input

// Provider is a input provider
type Provider struct{}

// New allocates the memory for Provider
func New() *Provider {
	return &Provider{}
}
