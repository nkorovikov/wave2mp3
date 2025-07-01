package input

import (
	"flag"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/nkorovikov/wave2mp3/internal/pkg/domain/io"
)

const (
	iArg    = "i"
	argHelp = "Input WAV file path"

	emptyStr = ""

	wavExt = ".wav"
	mp3Ext = ".mp3"
)

// Provide provides the input & output filess
func (p *Provider) Provide() (io.IOFiles, error) {
	inputFile := flag.String(iArg, emptyStr, argHelp)
	flag.Parse()

	if *inputFile == emptyStr {
		return io.IOFiles{}, fmt.Errorf("usage: wave2mp3 -i input.wav")
	}

	if !strings.HasSuffix(strings.ToLower(*inputFile), wavExt) {
		return io.IOFiles{}, fmt.Errorf("error: Input file must be a .wav file")
	}

	outputFile := strings.TrimSuffix(*inputFile, filepath.Ext(*inputFile)) + mp3Ext

	return io.IOFiles{
		InputFile:  *inputFile,
		OutputFile: outputFile,
	}, nil
}
