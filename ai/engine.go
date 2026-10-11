package ai

import "github.com/sashabaranov/go-openai"

const noexec = "[noexec]"

type Engine struct {
	mode         EngineMode
	config       *config.Config
	client       *openai.Client
	execMessages []openai.ChatCompletionMessage
	chatMessage  []openai.ChatCompletionMessage
	channel      chan EngineChatStreamOutput
	pipe         string
	running      bool
}
