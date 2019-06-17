package alc

import (
	"unsafe"
)

//-- ALC types
//
//-- | 8-bit boolean
//type ALCboolean = CChar
//
//-- | Character
//type ALCchar = CChar
//
//-- | Signed 8-bit 2\'s complement integer
//type ALCbyte = CSChar
//
//-- | Unsigned 8-bit integer
//type ALCubyte = CUChar
//
//-- | Signed 16-bit 2\'s complement integer
//type ALCshort = CShort
//
//-- | Unsigned 16-bit integer
//type ALCushort = CUShort
//
//-- | Signed 32-bit 2\'s complement integer
//type ALCint = CInt
//
//-- | Unsigned 32-bit integer
//type ALCuint = CUInt
//
//-- | Non-negatitve 32-bit binary integer size
//type ALCsizei = CInt
//
//-- | Enumerated 32-bit value
//type ALCenum = CInt
//
//-- | 32-bit IEEE754 floating-point
//type ALCfloat = CFloat
//
//-- | 64-bit IEEE754 floating-point
//type ALCdouble = CDouble
//

type Boolean byte
type Char []byte
type Byte int8
type Ubyte uint8

type Short int16
type Ushort uint16
type Int int32
type Uint uint32

type SizeI uint32
type Enum int32
type Float float32
type Double float64
type Device unsafe.Pointer
type Context unsafe.Pointer
