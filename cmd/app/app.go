package app

import (
	"fmt"

	"github.com/nkorovikov/wave2mp3/internal/pkg/domain/io"
	"github.com/nkorovikov/wave2mp3/internal/pkg/provider/convlib"
	"github.com/nkorovikov/wave2mp3/internal/pkg/provider/input"
	"github.com/nkorovikov/wave2mp3/internal/pkg/service/cmd"
)

type App struct {
	filesProvider   *input.Provider
	convlibProvider *convlib.Provider
	cmdService      *cmd.Service
}

func New() *App {
	return &App{
		filesProvider:   nil,
		convlibProvider: nil,
		cmdService:      nil,
	}
}

func (a *App) Init() error {
	a.filesProvider = input.New()
	a.convlibProvider = convlib.New()
	a.cmdService = cmd.New()

	return nil
}

func (a *App) Run() error {
	files, err := a.filesProvider.Provide()
	if err != nil {
		return fmt.Errorf("cannot provide files: %w", err)
	}

	ffmpegPath, err := a.convlibProvider.Provide()
	if err != nil {
		return fmt.Errorf("cannot provide convlib: %w", err)
	}
	defer a.convlibProvider.Close()

	err = a.cmdService.Exec(io.ExecArgs{
		InputFile:  files.InputFile,
		OutputFile: files.OutputFile,
		FfmpegPath: ffmpegPath,
	})
	if err != nil {
		return fmt.Errorf("cannot exec cmd: %w", err)
	}

	return nil
}
