package config

const (
	openai_key       = "OPENAI_KEY"
	openai_model     = "OPENAI_MODEL"
	openai_proxy     = "OPENAI_PROXY"
	open_temperature = "OPENAI_TEMPERATURE"
	open_max_tokens  = "OPENAI_MAX_TOKENS"
)

type AiConfig struct {
	key         string
	model       string
	proxy       string
	temperature string
	maxToken    string
}
