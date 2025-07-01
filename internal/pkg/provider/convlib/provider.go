package convlib

// Provider is a conv lib provider
type Provider struct {
	ffmpegPath string
}

// New allocates the memory for Provider
func New() *Provider {
	return &Provider{
		ffmpegPath: "",
	}
}
