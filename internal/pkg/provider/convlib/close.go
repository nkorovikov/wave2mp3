package convlib

import (
	"log"
	"os"
)

// Close closes the deps
func (p *Provider) Close() {
	err := os.Remove(p.ffmpegPath)
	if err != nil {
		log.Printf("cannot remove ffmpeg. Err: %v\n", err)
	}
}
