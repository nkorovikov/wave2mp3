package convlib

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

const (
	osLinux   = "linux"
	osMac     = "darwin"
	osWindows = "windows"

	extExe = ".exe"

	ffmpegName = "ffmpeg"

	accessCode = 0o755
)

var (
	//go:embed ffmpeg/ffmpeg-linux
	ffmpegLinux []byte

	//go:embed ffmpeg/ffmpeg-macos
	ffmpegMacos []byte

	//go:embed ffmpeg/ffmpeg-win.exe
	ffmpegWin []byte
)

func getFFmpegBinary() ([]byte, error) {
	switch runtime.GOOS {
	case osLinux:
		return ffmpegLinux, nil
	case osMac:
		return ffmpegMacos, nil
	case osWindows:
		return ffmpegWin, nil
	default:
		return nil, fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
}

// Provide provides the conv lib path
func (p *Provider) Provide() (string, error) {
	ffmpegData, err := getFFmpegBinary()
	if err != nil {
		return "", fmt.Errorf("FFmpeg error: %w", err)
	}

	tmpDir := os.TempDir()
	ffmpegPath := filepath.Join(tmpDir, ffmpegName)

	if runtime.GOOS == osWindows {
		ffmpegPath += extExe
	}

	err = os.WriteFile(ffmpegPath, ffmpegData, accessCode)
	if err != nil {
		return "", fmt.Errorf("failed to extract FFmpeg: %w", err)
	}

	p.ffmpegPath = ffmpegPath

	return ffmpegPath, nil
}
