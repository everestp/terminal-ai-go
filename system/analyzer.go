package system

const APPLICATION_NAME = "terminal-assistant"

type Analysis struct {
	operatingSystem OperatingSystem

	distribution  string
	shell         string
	homeDirectory string
	username      string
	editor        string
	configFile    string
}
