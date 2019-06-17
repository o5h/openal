package alc

// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	dll                    *windows.DLL
	captureCloseDevice     *windows.Proc
	captureOpenDevice      *windows.Proc
	captureStart           *windows.Proc
	captureStop            *windows.Proc
	closeDevice            *windows.Proc
	destroyContext         *windows.Proc
	getContextsDevice      *windows.Proc
	getCurrentContext      *windows.Proc
	getEnumValue           *windows.Proc
	getProcAddress         *windows.Proc
	getString              *windows.Proc
	openDevice             *windows.Proc
	procCaptureSamples     *windows.Proc
	procCreateContext      *windows.Proc
	processContext         *windows.Proc
	procGetError           *windows.Proc
	procGetIntegerv        *windows.Proc
	procIsExtensionPresent *windows.Proc
	procMakeContextCurrent *windows.Proc
	suspendContext         *windows.Proc
)

func init() {
	dll = windows.MustLoadDLL("OpenAL32.dll")
	captureCloseDevice = dll.MustFindProc("alcCaptureCloseDevice")
	captureOpenDevice = dll.MustFindProc("alcCaptureOpenDevice")
	captureStart = dll.MustFindProc("alcCaptureStart")
	captureStop = dll.MustFindProc("alcCaptureStop")
	closeDevice = dll.MustFindProc("alcCloseDevice")
	destroyContext = dll.MustFindProc("alcDestroyContext")
	getContextsDevice = dll.MustFindProc("alcGetContextsDevice")
	getCurrentContext = dll.MustFindProc("alcGetCurrentContext")
	getEnumValue = dll.MustFindProc("alcGetEnumValue")
	getProcAddress = dll.MustFindProc("alcGetProcAddress")
	getString = dll.MustFindProc("alcGetString")
	openDevice = dll.MustFindProc("alcOpenDevice")
	procCaptureSamples = dll.MustFindProc("alcCaptureSamples")
	procCreateContext = dll.MustFindProc("alcCreateContext")
	processContext = dll.MustFindProc("alcProcessContext")
	procGetError = dll.MustFindProc("alcGetError")
	procGetIntegerv = dll.MustFindProc("alcGetIntegerv")
	procIsExtensionPresent = dll.MustFindProc("alcIsExtensionPresent")
	procMakeContextCurrent = dll.MustFindProc("alcMakeContextCurrent")
	suspendContext = dll.MustFindProc("alcSuspendContext")
}

//ALCdevice *alcOpenDevice(	const ALCchar *devicename);
func OpenDevice(name string) Device {
	if name == "" {
		r0, _, _ := openDevice.Call(uintptr(0))
		return Device(r0)
	} else {
		cstr := unsafe.Pointer(C.CString(name))
		defer C.free(cstr)
		r0, _, _ := openDevice.Call(uintptr(cstr))
		return Device(r0)
	}
}

//ALCboolean alcCloseDevice(ALCdevice *device);
func CloseDevice(d Device) bool {
	r0, _, _ := closeDevice.Call(uintptr(d))
	return r0 != 0
}

//ALCcontext * alcCreateContext( ALCdevice *device,ALCint* attrlist);
func CreateContext(device Device, attrlist []int) Context {
	if len(attrlist) == 0 {
		r0, _, _ := procCreateContext.Call(uintptr(device), uintptr(uintptr(0)))
		return Context(r0)
	} else {
		r0, _, _ := procCreateContext.Call(
			uintptr(device),
			uintptr(unsafe.Pointer(&attrlist[0])))
		return Context(r0)
	}
}
func GetCurrentContext() Context {
	r0, _, _ := getCurrentContext.Call()
	return Context(r0)
}

//ALCdevice * alcGetContextsDevice( ALCcontext *context );
func GetContextsDevice(c Context) Device {
	r0, _, _ := getContextsDevice.Call(uintptr(c))
	return Device(r0)
}

//void alcProcessContext(ALCcontext *context);
func ProcessContext(c Context) {
	processContext.Call(
		uintptr(c))
}

//void alcSuspendContext(ALCcontext *context);
func SuspendContext(c Context) {
	suspendContext.Call(
		uintptr(c))
}

//void alcDestroyContext(ALCcontext *context);
func DestroyContext(cont Context) {
	destroyContext.Call(
		uintptr(cont))
}

func MakeContextCurrent(ctx Context) bool {
	r0, _, _ := procMakeContextCurrent.Call(uintptr(ctx))
	return r0 != 0
}

//ALCenum alcGetError( ALCdevice *device );
func GetError(device Device) error {
	r0, _, _ := procGetError.Call(uintptr(device))
	switch r0 {
	case ALC_NO_ERROR:
		return nil
	case ALC_INVALID_DEVICE:
		return InvalidDevice
	case ALC_INVALID_CONTEXT:
		return InvalidContext
	case ALC_INVALID_ENUM:
		return InvalidEnum
	case ALC_INVALID_VALUE:
		return InvalidValue
	case ALC_OUT_OF_MEMORY:
		return OutOfMemory
	default:
		return fmt.Errorf("Unknown error 0x%x", r0)
	}
}

//const ALCchar * alcGetString(	ALCdevice *device,ALenum param);
func GetString(device Device, param Enum) []string {
	r0, _, _ := getString.Call(
		uintptr(device),
		uintptr(param))
	if param == ALC_EXTENSIONS {
		return goString(r0)
	}
	return goStrings(r0)
}

func goString(ptr uintptr) []string {
	strings := make([]string, 0)
	cPtr := (*C.char)(unsafe.Pointer(ptr))
	str := C.GoString(cPtr)
	strings = append(strings, str)
	return strings
}

func goStrings(ptr uintptr) []string {
	strings := make([]string, 0)

	for {
		nextByte := *(*byte)(unsafe.Pointer(ptr))
		if nextByte == 0 {
			break
		}

		cPtr := (*C.char)(unsafe.Pointer(ptr))
		str := C.GoString(cPtr)
		strings = append(strings, str)
		lenInBytes := len([]byte(str))
		ptr = ptr + uintptr(lenInBytes) + 1

	}
	return strings
}

//ALCboolean alcIsExtensionPresent(  ALCdevice *device,   const ALCchar *extName );
func IsExtensionPresent(device Device, extName string) bool {
	cstr := unsafe.Pointer(C.CString(extName))
	defer C.free(cstr)
	r0, _, _ := procIsExtensionPresent.Call(uintptr(device), uintptr(cstr))
	return r0 != 0
}

//ALCenum alcGetEnumValue(ALCdevice *device,const ALCchar *enumName);
func GetEnumValue(d Device, enumName string) Enum {
	cstr := unsafe.Pointer(C.CString(enumName))
	defer C.free(cstr)
	r0, _, _ := getEnumValue.Call(uintptr(d), uintptr(cstr))
	return Enum(r0)
}

//void * alcGetProcAddress(ALCdevice *device,const ALCchar *funcName);
func GetProcAddress(d Device, funcName string) unsafe.Pointer {
	cstr := unsafe.Pointer(C.CString(funcName))
	defer C.free(cstr)
	r0, _, _ := getProcAddress.Call(uintptr(d), uintptr(cstr))
	return unsafe.Pointer(r0)
}

//ALCdevice * alcCaptureOpenDevice(const ALCchar *devicename,ALCuint frequency,ALCenum format,ALCsizei buffersize);
func CaptureOpenDevice(d Device, freq uint32, form Enum, buffersize int32) Device {
	r0, _, _ := captureOpenDevice.Call(
		uintptr(d),
		uintptr(freq),
		uintptr(form),
		uintptr(buffersize))
	return Device(r0)
}

//ALCboolean alcCaptureCloseDevice(ALCdevice *device);
func CaptureCloseDevice(d Device) bool {
	r0, _, _ := captureCloseDevice.Call(
		uintptr(d))
	return r0 != 0
}

//void alcCaptureStart(ALCdevice *device);
func CaptureStart(d Device) {
	captureStart.Call(uintptr(d))

}

//void alcCaptureStop(ALCdevice *device);
func CaptureStop(d Device) {
	captureStop.Call(uintptr(d))
}

//ALC_MAJOR_VERSION
//ALC_MINOR_VERSION
//ALC_ATTRIBUTES_SIZE
//ALC_ALL_ATTRIBUTES

//void alcGetIntegerv(ALCdevice *device,ALCenum param,ALCsizei size,ALCint *data);
func GetIntegerv(d Device, param Enum, data []int32) {
	procGetIntegerv.Call(
		uintptr(d),
		uintptr(param),
		uintptr(len(data)),
		uintptr(unsafe.Pointer(&data[0])))
}

//void alcCaptureSamples(	ALCdevice *device,	ALCvoid *buffer,	ALCsizei samples   );
func CaptureSamples(device Device, buffer []byte, samples int32) {
	procCaptureSamples.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(samples))
}
