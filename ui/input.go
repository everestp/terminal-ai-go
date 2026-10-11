package ui

import "golang.org/x/text/message/pipeline"

type UiInput struct {
	runMode    RunMode
	promptMode PromptMode
	args       string
	pipeline   string
}
