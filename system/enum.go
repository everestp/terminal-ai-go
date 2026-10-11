package system

type OperatingSystem int

const (
	UnknownOperatingSystem OperatingSystem = iota
	LinuxOperatingSystem
	MacOperatingSystem
	WindowOperatingSystem
)
