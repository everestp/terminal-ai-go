package ui

import "github.com/charmbracelet/bubbles/textinput"

const (
	exec_icon        = "🚀"
	exec_placeholder = "Execute something..."

	config_icon        = "🔑"
	config_placeholder = "Enter your OpenAI API key..."

	chat_icon        = "🤖"
	chat_placeholder = "Ask me something..."
)

type Prompt struct {
	mode  PromptMode
	input textinput.Model
}
