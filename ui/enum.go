package ui

type PromptMode int

const (
	ExecPromptMode PromptMode = iota
	ConfigPromptMode
	ChatPromptMode
	DefaultPromptMode
)

type RunMode int

const (
	CliMode RunMode = iota
	ReplMode
)
