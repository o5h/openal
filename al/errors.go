package al

import "fmt"

const (
	ALC_NO_ERROR          = Enum(0)
	ALC_INVALID_CONTEXT   = Enum(40962)
	ALC_INVALID_DEVICE    = Enum(40961)
	ALC_INVALID_ENUM      = Enum(40963)
	ALC_INVALID_OPERATION = Enum(40966)
	ALC_INVALID_VALUE     = Enum(40964)
	ALC_OUT_OF_MEMORY     = Enum(40965)
)

var (
	InvalidContext   = fmt.Errorf("Bad context was passed to an OpenAL function")
	InvalidDevice    = fmt.Errorf("Bad device was passed to an OpenAL function")
	InvalidEnum      = fmt.Errorf("Unknown enum value was passed to an OpenAL function")
	InvalidOperation = fmt.Errorf("InvalidOperation")
	InvalidValue     = fmt.Errorf("One of the parameters has an invalid value")
	OutOfMemory      = fmt.Errorf("The specified device is invalid, or can not capture audio. ")
)
