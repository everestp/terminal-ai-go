package ui

import (
	"github.com/everestp/terminal-ai-go/ai"
	"github.com/everestp/terminal-ai-go/config"
	"github.com/everestp/terminal-ai-go/history"
)

type UiStates struct {
	error       error
	runMode     RunMode
	promptMode  PromptMode
	configuring bool
	querying    bool
	confirming  bool
	executing   bool
	args        string
	pipe        string
	buffer      string
	commad      string
}

type UiDimensions struct {
	width int

	height int
}
type UiComponents struct {
	prompt   *Prompt
	renderer *Renderer
	spinner  *Spinner
}

type Ui struct {
	state     UiStates
	dimension UiDimensions
	componets UiComponents
	config    *config.Config
	engine    *ai.Engine
	hisotry   *history.Histroy
}
