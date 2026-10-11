package ui

import (
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const (
	exec_color    = "#F97316" // Orange — execution 🚀
	config_color  = "#A78BFA" // Purple — configuration 🔑
	chat_color    = "#38BDF8" // Blue — AI chat 🤖
	help_color    = "#94A3B8" // Slate — help
	error_color   = "#EF4444" // Red — errors
	warning_color = "#FBBF24" // Amber — warnings
	success_color = "#10B981" // Green — success
)

type Renderer struct {
	contentRenderer *glamour.TermRenderer
	successRenderer *lipgloss.Style
	warningRenderer *lipgloss.Style
	errorRenderer   *lipgloss.Style
	helpRenderer    *lipgloss.Style
}
