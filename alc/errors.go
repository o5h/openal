package alc

import "fmt"

const (
	ALC_NO_ERROR        = 0
	ALC_INVALID_DEVICE  = 0xA001
	ALC_INVALID_CONTEXT = 0xA002
	ALC_INVALID_ENUM    = 0xA003
	ALC_INVALID_VALUE   = 0xA004
	ALC_OUT_OF_MEMORY   = 0xA005
)

var (
	InvalidDevice  = fmt.Errorf("No device")
	InvalidContext = fmt.Errorf("Invalid Context ID")
	InvalidEnum    = fmt.Errorf("Bad enum")
	InvalidValue   = fmt.Errorf("Bad Value")
	OutOfMemory    = fmt.Errorf("Out of memory")
)
