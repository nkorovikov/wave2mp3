package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/nkorovikov/wave2mp3/internal/pkg/domain/io"
)

const (
	iArg            = "-i"
	codecArg        = "-codec:a"
	codecArgValue   = "libmp3lame"
	bitRateArg      = "-b:a"
	bitRateArgValue = "320k"
	qualityArg      = "-q:a"
	qualityArgValue = "0"
	outputArg       = "-y"
	dangerousChars  = ";&|<>$()"
	transitionSym   = "../"
)

// Exec executes the command
func (s *Service) Exec(args io.ExecArgs) error {
	if !isSafePath(args.InputFile) || !isSafePath(args.OutputFile) || !isSafePath(args.FfmpegPath) {
		return fmt.Errorf("invalid file path")
	}

	//nolint:gosec // check the paths
	cmd := exec.Command(
		args.FfmpegPath,
		iArg, args.InputFile,
		codecArg, codecArgValue,
		bitRateArg, bitRateArgValue,
		qualityArg, qualityArgValue,
		outputArg,
		args.OutputFile,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Printf("Converting %s to %s (%s)...", args.InputFile, args.OutputFile, bitRateArgValue)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("conversion failed: %w", err)
	}

	log.Println("Conversion complete!")

	return nil
}

func isSafePath(path string) bool {
	if strings.ContainsAny(path, dangerousChars) {
		return false
	}

	cleanPath := filepath.Clean(path)
	if path != cleanPath || strings.HasPrefix(cleanPath, transitionSym) {
		return false
	}

	return true
}
