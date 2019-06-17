package al

// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	dll                  *windows.DLL
	buffer3f             *windows.Proc
	buffer3i             *windows.Proc
	bufferData           *windows.Proc
	bufferf              *windows.Proc
	bufferfv             *windows.Proc
	bufferi              *windows.Proc
	bufferiv             *windows.Proc
	deleteBuffers        *windows.Proc
	deleteSources        *windows.Proc
	disable              *windows.Proc
	distanceModel        *windows.Proc
	dopplerFactor        *windows.Proc
	dopplerVelocity      *windows.Proc
	enable               *windows.Proc
	genBuffers           *windows.Proc
	genSources           *windows.Proc
	getBoolean           *windows.Proc
	getBooleanv          *windows.Proc
	getBuffer3f          *windows.Proc
	getBuffer3i          *windows.Proc
	getBufferf           *windows.Proc
	getBufferfv          *windows.Proc
	getBufferi           *windows.Proc
	getBufferiv          *windows.Proc
	getDouble            *windows.Proc
	getDoublev           *windows.Proc
	getEnumValue         *windows.Proc
	getError             *windows.Proc
	getFloat             *windows.Proc
	getFloatv            *windows.Proc
	getInteger           *windows.Proc
	getIntegerv          *windows.Proc
	getListener3f        *windows.Proc
	getListener3i        *windows.Proc
	getListenerf         *windows.Proc
	getListenerfv        *windows.Proc
	getListeneri         *windows.Proc
	getListeneriv        *windows.Proc
	getProcAddress       *windows.Proc
	getSource3f          *windows.Proc
	getSource3i          *windows.Proc
	getSourcef           *windows.Proc
	getSourcefv          *windows.Proc
	getSourcei           *windows.Proc
	getSourceiv          *windows.Proc
	getString            *windows.Proc
	isBuffer             *windows.Proc
	isEnabled            *windows.Proc
	isExtensionPresent   *windows.Proc
	isSource             *windows.Proc
	listener3f           *windows.Proc
	listener3i           *windows.Proc
	listenerf            *windows.Proc
	listenerfv           *windows.Proc
	listeneri            *windows.Proc
	listeneriv           *windows.Proc
	procSourcei          *windows.Proc
	procSourcePlay       *windows.Proc
	source3f             *windows.Proc
	sourcef              *windows.Proc
	sourcefv             *windows.Proc
	sourceiv             *windows.Proc
	sourcePause          *windows.Proc
	sourcePausev         *windows.Proc
	sourcePlayv          *windows.Proc
	sourceQueueBuffers   *windows.Proc
	sourceRewind         *windows.Proc
	sourceRewindv        *windows.Proc
	sourceStop           *windows.Proc
	sourceStopv          *windows.Proc
	sourceUnqueueBuffers *windows.Proc
	speedOfSound         *windows.Proc
)

func init() {
	dll = windows.MustLoadDLL("OpenAL32.dll")
	buffer3f = dll.MustFindProc("alBuffer3f")
	buffer3i = dll.MustFindProc("alBuffer3i")
	bufferData = dll.MustFindProc("alBufferData")
	bufferf = dll.MustFindProc("alBufferf")
	bufferfv = dll.MustFindProc("alBufferfv")
	bufferi = dll.MustFindProc("alBufferi")
	bufferiv = dll.MustFindProc("alBufferiv")
	deleteBuffers = dll.MustFindProc("alDeleteBuffers")
	deleteSources = dll.MustFindProc("alDeleteSources")
	disable = dll.MustFindProc("alDisable")
	distanceModel = dll.MustFindProc("alDistanceModel")
	dopplerFactor = dll.MustFindProc("alDopplerFactor")
	dopplerVelocity = dll.MustFindProc("alDopplerVelocity")
	enable = dll.MustFindProc("alEnable")
	genBuffers = dll.MustFindProc("alGenBuffers")
	genSources = dll.MustFindProc("alGenSources")
	getBoolean = dll.MustFindProc("alGetBoolean")
	getBooleanv = dll.MustFindProc("alGetBooleanv")
	getBuffer3f = dll.MustFindProc("alGetBuffer3f")
	getBuffer3i = dll.MustFindProc("alGetBuffer3i")
	getBufferf = dll.MustFindProc("alGetBufferf")
	getBufferfv = dll.MustFindProc("alGetBufferfv")
	getBufferi = dll.MustFindProc("alGetBufferi")
	getBufferiv = dll.MustFindProc("alGetBufferiv")
	getDouble = dll.MustFindProc("alGetDouble")
	getDoublev = dll.MustFindProc("alGetDoublev")
	getEnumValue = dll.MustFindProc("alGetEnumValue")
	getError = dll.MustFindProc("alGetError")
	getFloat = dll.MustFindProc("alGetFloat")
	getFloatv = dll.MustFindProc("alGetFloatv")
	getInteger = dll.MustFindProc("alGetInteger")
	getIntegerv = dll.MustFindProc("alGetIntegerv")
	getListener3f = dll.MustFindProc("alGetListener3f")
	getListener3i = dll.MustFindProc("alGetListener3i")
	getListenerf = dll.MustFindProc("alGetListenerf")
	getListenerfv = dll.MustFindProc("alGetListenerfv")
	getListeneri = dll.MustFindProc("alGetListeneri")
	getListeneriv = dll.MustFindProc("alGetListeneriv")
	getProcAddress = dll.MustFindProc("alGetProcAddress")
	getSource3f = dll.MustFindProc("alGetSource3f")
	getSource3i = dll.MustFindProc("alGetSource3i")
	getSourcef = dll.MustFindProc("alGetSourcef")
	getSourcefv = dll.MustFindProc("alGetSourcefv")
	getSourcei = dll.MustFindProc("alGetSourcei")
	getSourceiv = dll.MustFindProc("alGetSourceiv")
	getString = dll.MustFindProc("alGetString")
	isBuffer = dll.MustFindProc("alIsBuffer")
	isEnabled = dll.MustFindProc("alIsEnabled")
	isExtensionPresent = dll.MustFindProc("alIsExtensionPresent")
	isSource = dll.MustFindProc("alIsSource")
	listener3f = dll.MustFindProc("alListener3f")
	listener3i = dll.MustFindProc("alListener3i")
	listenerf = dll.MustFindProc("alListenerf")
	listenerfv = dll.MustFindProc("alListenerfv")
	listeneri = dll.MustFindProc("alListeneri")
	listeneriv = dll.MustFindProc("alListeneriv")
	procSourcei = dll.MustFindProc("alSourcei")
	procSourcePlay = dll.MustFindProc("alSourcePlay")
	source3f = dll.MustFindProc("alSource3f")
	sourcef = dll.MustFindProc("alSourcef")
	sourcefv = dll.MustFindProc("alSourcefv")
	sourceiv = dll.MustFindProc("alSourceiv")
	sourcePause = dll.MustFindProc("alSourcePause")
	sourcePausev = dll.MustFindProc("alSourcePausev")
	sourcePlayv = dll.MustFindProc("alSourcePlayv")
	sourceQueueBuffers = dll.MustFindProc("alSourceQueueBuffers")
	sourceRewind = dll.MustFindProc("alSourceRewind")
	sourceRewindv = dll.MustFindProc("alSourceRewindv")
	sourceStop = dll.MustFindProc("alSourceStop")
	sourceStopv = dll.MustFindProc("alSourceStopv")
	sourceUnqueueBuffers = dll.MustFindProc("alSourceUnqueueBuffers")
	speedOfSound = dll.MustFindProc("alSpeedOfSound")
}

func GetError() error {
	r0, _, _ := getError.Call()
	switch Enum(r0) {
	case ALC_NO_ERROR:
		return nil
	case ALC_INVALID_CONTEXT:
		return InvalidContext
	case ALC_INVALID_DEVICE:
		return InvalidDevice
	case ALC_INVALID_ENUM:
		return InvalidEnum
	case ALC_INVALID_OPERATION:
		return InvalidOperation
	case ALC_INVALID_VALUE:
		return InvalidValue
	case ALC_OUT_OF_MEMORY:
		return OutOfMemory
	}
	panic(fmt.Sprintf("Unknown error %v", r0))
}

//void alGenBuffers(  ALsizei n,   ALuint *buffers );
func GenBuffers(buffers []uint32) {
	genBuffers.Call(
		uintptr(len(buffers)),
		uintptr(unsafe.Pointer(&buffers[0])))
}

//void alSourceQueueBuffers(ALuint source,ALsizei n,ALuint* buffers);
func SourceQueueBuffers(source uint32, n uint32, buffers []uint32) {
	sourceQueueBuffers.Call(
		uintptr(source),
		uintptr(n),
		uintptr(unsafe.Pointer(&buffers[0])))
}

//void alBufferf(ALuint buffer,ALenum param,ALfloat value);
func Bufferf(buffer []uint32, enu Enum, v1 float32) {
	bufferf.Call(
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(enu),
		uintptr(v1))
}

//void alBufferfv(ALuint buffer,ALenum param,ALfloat *values);
func Bufferfv(buffer []uint32, enu Enum, floatptr *float32) {
	bufferfv.Call(
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(enu),
		uintptr(unsafe.Pointer(floatptr)))
}

//void alBufferi(ALuint buffer,ALenum param,ALint value);
func Bufferi(buffer []uint32, enu Enum, v1 *int32) {
	bufferi.Call(
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(enu),
		uintptr(unsafe.Pointer(v1)))
}

//void alBuffer3f(ALuint buffer,ALenum param,ALfloat v1,ALfloat v2,ALfloat v3);
func Buffer3f(buff []uint32, enu Enum, float1 float32, float2 float32, float3 float32) {
	buffer3f.Call(
		uintptr(unsafe.Pointer(&buff[0])),
		uintptr(enu),
		uintptr(float1),
		uintptr(float2),
		uintptr(float3))
}

//void alBufferiv(ALuint buffer,ALenum param,ALint *values);
func Bufferiv(buffer []int32, enu Enum, val *int32) {
	bufferiv.Call(
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(enu),
		uintptr(unsafe.Pointer(val)))
}

//void alBuffer3i(ALuint buffer,ALenum param,ALint v1,ALint v2,ALint v3);
func Buffer3i(buffer []uint32, enu Enum, par1 int32, par2 int32, par3 int32) {
	buffer3i.Call(
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(enu),
		uintptr(par1),
		uintptr(par2),
		uintptr(par3))
}

//void alDeleteBuffers(ALsizei n,ALuint *buffers);
func DeleteBuffers(buffers []uint32) {
	deleteBuffers.Call(
		uintptr(len(buffers)),
		uintptr(unsafe.Pointer(&buffers[0])))
}

//void alGetBufferi(ALuint buffer,ALenum pname,ALint *value);
func GetBufferi(buffer uint32, param Enum, value *int32) {
	getBufferi.Call(
		uintptr(buffer),
		uintptr(param),
		uintptr(unsafe.Pointer(value)))
}

//void alGetBufferiv(ALuint buffer,ALenum pname,ALint *values);
func GetBufferiv(buffer uint32, param Enum, values *int32) {
	getBufferiv.Call(
		uintptr(buffer),
		uintptr(param),
		uintptr(unsafe.Pointer(values)))
}

// Only For Extentions
//void alGetBufferf(ALuint buffer,ALenum pname,ALfloat *value);
func GetBufferf(buffer uint32, param Enum, value unsafe.Pointer) {
	getBufferf.Call(
		uintptr(buffer),
		uintptr(param),
		uintptr(value))
}

//Only For Extentions
//void alGetBufferfv(ALuint buffer,ALenum pname,ALfloat *values);
func GetBufferfv(buffer uint32, param Enum, values unsafe.Pointer) {
	getBufferfv.Call(
		uintptr(buffer),
		uintptr(param),
		uintptr(values))
}

//ALboolean alIsBuffer(ALuint buffer);
func IsBuffer(buffer uint32) bool {
	result, _, _ := isBuffer.Call(uintptr(buffer))
	return result != 0
}

//void alGetBuffer3f(ALuint buffer,ALenum pname,ALfloat *v1,ALfloat *v2,ALfloat *v3);
func GetBuffer3f(buffer []uint32, param Enum, v1 float32, v2 float32, v3 float32) {
	getBuffer3f.Call(
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(param),
		uintptr(v1),
		uintptr(v2),
		uintptr(v3))
}

//void alGetBuffer3i(ALuint buffer,ALenum pname,ALint *v1,ALint *v2,ALint *v3);
func GetBuffer3i(buffer []uint32, param Enum, i1 int32, i2 int32, i3 int32) {
	getBuffer3i.Call(
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(param),
		uintptr(i1),
		uintptr(i2),
		uintptr(i3))
}

//ALenum alGetEnumValue(const ALchar *ename);
func GetEnumValue(name string) Enum {
	cstr := unsafe.Pointer(C.CString(name))
	defer C.free(cstr)
	r0, _, _ := getEnumValue.Call(uintptr(cstr))
	return Enum(r0)
}

//void alGenSources(ALsizei n,ALuint *sources);
func GenSources(source []uint32) {
	genSources.Call(uintptr(len(source)), uintptr(unsafe.Pointer(&source[0])))
}

//void alDeleteSources(ALsizei n,ALuint *sources);
func DeleteSources(source []uint32) {
	deleteSources.Call(
		uintptr(len(source)),
		uintptr(unsafe.Pointer(&source[0])))
}

//param
// AL_FORMAT_MONO8
// AL_FORMAT_MONO16
// AL_FORMAT_STEREO8
// AL_FORMAT_STEREO16
//void alBufferData(ALuint buffer,ALenum format,const ALvoid *data,ALsizei size,ALsizei freq);
func BufferData(buffername uint32, format Enum, data []byte, freq uint32) {
	bufferData.Call(
		uintptr(buffername),
		uintptr(format),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
		uintptr(freq))
}

//void alSourcei(  ALuint source,  ALenum param,  ALint value );
func Sourcei(source uint32, param Enum, value int32) {
	procSourcei.Call(uintptr(source), uintptr(param), uintptr(value))
}

func SourcePlay(source uint32) {
	procSourcePlay.Call(uintptr(source))
}

//void alSourcePause(ALuint source);
func SourcePause(source uint32) {
	sourcePause.Call(uintptr(source))
}

//boolean alIsSource(ALuint source);
func IsSource(source uint32) bool {
	r0, _, _ := isSource.Call(
		uintptr(source))
	return r0 != 0
}

//void alEnable(ALenum capability);
func Enable(cap Enum) {
	enable.Call(uintptr(cap))
}

//void alDisable(ALenum capability);
func Disable(cap Enum) {
	disable.Call(uintptr(cap))
}

//ALboolean alIsEnabled(ALenum capability);
func IsEnabled(enu Enum) bool {
	r0, _, _ := isEnabled.Call(
		uintptr(enu))
	return r0 != 0
}

//void * alGetProcAddress(const ALchar *fname);
func GetProcAddress(fname string) unsafe.Pointer {
	cstr := unsafe.Pointer(C.CString(fname))
	defer C.free(cstr)
	r0, _, _ := getProcAddress.Call(uintptr(cstr))
	return unsafe.Pointer(r0)
}

/*AL_INVERSE_DISTANCE
AL_INVERSE_DISTANCE_CLAMPED
AL_LINEAR_DISTANCE
AL_LINEAR_DISTANCE_CLAMPED
AL_EXPONENT_DISTANCE
AL_EXPONENT_DISTANCE_CLAMPED
AL_NONE
*/
//void alDistanceModel(ALenum value);
func DistanceModel(enu Enum) {
	distanceModel.Call(
		uintptr(enu))
}

//void alDopplerFactor(ALfloat value);
func DopplerFactor(val float32) {
	dopplerFactor.Call(
		uintptr(val))
}

//alDopplerVelocity
func DopplerVelocity(f float32) {
	dopplerVelocity.Call(uintptr(f))
}

//ALboolean alGetBoolean(ALenum param);
func GetBoolean(enu Enum) bool {
	r0, _, _ := getBoolean.Call(uintptr(enu))
	return r0 != 0
}

//void alGetBooleanv(ALenum param,ALboolean *data);
func GetBooleanv(enu Enum, ptr unsafe.Pointer) {
	getBooleanv.Call(uintptr(enu), uintptr(ptr))
}

//Aldouble alGetDouble(ALenum param);
func GetDouble(param Enum) float64 {
	r0, _, _ := getDouble.Call(uintptr(param))
	return float64(r0)
}

//void alGetDoublev(ALenum param,ALdouble *data);
func GetDoublev(param Enum, data unsafe.Pointer) {
	getDoublev.Call(uintptr(param), uintptr(data))
}

//ALfloat alGetFloat(ALenum param);
func GetFloat(param Enum) float32 {
	r0, _, _ := getFloat.Call(uintptr(param))
	return float32(r0)
}

//void alGetFloatv(ALenum param,ALfloat *data);
func GetFloatv(param Enum, data unsafe.Pointer) {
	getFloatv.Call(uintptr(param), uintptr(data))
}

//Alint alGetInteger(ALenum param);
func GetInteger(param Enum) int32 {
	r0, _, _ := getInteger.Call(uintptr(param))
	return int32(r0)
}

//void alGetIntegerv(ALenum param,ALint *data);
func GetIntegerv(param Enum, data unsafe.Pointer) {
	getIntegerv.Call(uintptr(param), uintptr(data))
}

//void alGetListener3f(ALenum param,ALfloat *v1,ALfloat *v2,ALfloat *v3);
func GetListener3f(param Enum, v1, v2, v3 float32) {
	getListener3f.Call(
		uintptr(param),
		uintptr(v1),
		uintptr(v2),
		uintptr(v3))
}

//void alGetListener3i(ALenum param,ALint *v1,ALint *v2,ALint *v3);
func GetListener3i(param Enum, v1, v2, v3 int32) {
	getListener3i.Call(
		uintptr(param),
		uintptr(v1),
		uintptr(v2),
		uintptr(v3))
}

//void alGetListenerf(ALenum param,ALfloat *value);
func GetListenerf(param Enum, v1 float32) {
	getListenerf.Call(
		uintptr(param),
		uintptr(v1))
}

//void alGetListenerfv(ALenum param,ALfloat *values);
func GetListenerfv(param Enum, values unsafe.Pointer) {
	getListenerfv.Call(
		uintptr(param),
		uintptr(values))
}

//For Extentions
//void alGetListeneri(ALenum param,ALint *value);
func GetListeneri(param Enum, value unsafe.Pointer) {
	getListeneri.Call(
		uintptr(param),
		uintptr(value))
}

//void alGetListeneriv(ALenum param,ALint *values);
func GetListeneriv(param Enum, values unsafe.Pointer) {
	getListeneriv.Call(
		uintptr(param),
		uintptr(values))
}

//void alSourcePlayv(ALsizei n,ALuint *sources);
func SourcePlayv(num uint32, sources []uint32) {
	sourcePlayv.Call(
		uintptr(num),
		uintptr(unsafe.Pointer(&sources[0])))
}

//void alSourcePausev(ALsizei n,ALuint *sources);
func SourcePausev(num uint32, sources []uint32) {
	sourcePausev.Call(
		uintptr(num),
		uintptr(unsafe.Pointer(&sources[0])))
}

//void alSourceStop(ALuint source);
func SourceStop(source uint32) {
	sourceStop.Call(
		uintptr(source))
}

//void alSourceStopv(ALsizei n,ALuint *sources);
func SourceStopv(n uint32, sources []uint32) {
	sourceStopv.Call(
		uintptr(n),
		uintptr(unsafe.Pointer(&sources[0])))
}

//void alSourceRewind(ALuint source);
func SourceRewind(source uint32) {
	sourceRewind.Call(
		uintptr(source))
}

//void alSourceRewindv(ALsizei n,ALuint *sources);
func SourceRewindv(n uint32, sources []uint32) {
	sourceRewindv.Call(
		uintptr(n),
		uintptr(unsafe.Pointer(&sources[0])))
}

/*
param the name of the attribute to set:
 AL_POSITION
 AL_VELOCITY
 AL_DIRECTION
*/
//void alSource3f(ALuint source,ALenum param,ALfloat v1,ALfloat v2,ALfloat v3);
func Source3f(source uint32, param Enum, v1, v2, v3 float32) {
	source3f.Call(
		uintptr(source),
		uintptr(param),
		uintptr(v1),
		uintptr(v2),
		uintptr(v3))
}

/*
param the name of the attribute to set:
 AL_PITCH
 AL_GAIN
 AL_MIN_GAIN
 AL_MAX_GAIN
 AL_MAX_DISTANCE
 AL_ROLLOFF_FACTOR
 AL_CONE_OUTER_GAIN
 AL_CONE_INNER_ANGLE
 AL_CONE_OUTER_ANGLE
 AL_REFERENCE_DISTANCE
*/
//void alSourcef(ALuint source,ALenum param,ALfloat value);
func Sourcef(source uint32, param Enum, v1 float32) {
	sourcef.Call(
		uintptr(source),
		uintptr(v1))
}

//void alSourcefv(ALuint source,ALenum param,ALfloat *values);
func Sourcefv(source uint32, param Enum, values unsafe.Pointer) {
	sourcefv.Call(
		uintptr(source),
		uintptr(values))
}

//void alSourceiv(ALuint source,ALenum param,ALint *values);
func Sourceiv(source uint32, param Enum, values unsafe.Pointer) {
	sourceiv.Call(
		uintptr(source),
		uintptr(param),
		uintptr(values))
}

//void alGetSource3f(ALuint source,ALenum param,ALfloat *v1,ALfloat *v2,ALfloat *v3);
func GetSource3f(sourse uint32, param Enum, v1, v2, v3 float32) {
	getSource3f.Call(
		uintptr(sourse),
		uintptr(param),
		uintptr(v1),
		uintptr(v2),
		uintptr(v3))
}

//void alGetSource3i(ALuint source,ALenum param,ALint *v1,ALint *v2,ALint *v3);
func GetSource3i(sourse uint32, param Enum, v1, v2, v3 *int32) {
	getSource3i.Call(
		uintptr(sourse),
		uintptr(param),
		uintptr(unsafe.Pointer(v1)),
		uintptr(unsafe.Pointer(v2)),
		uintptr(unsafe.Pointer(v3)))
}

/*param the name of the attribute to retrieve:
AL_PITCH
AL_GAIN
AL_MIN_GAIN
AL_MAX_GAIN
AL_MAX_DISTANCE
AL_ROLLOFF_FACTOR
AL_CONE_OUTER_GAIN
AL_CONE_INNER_ANGLE
AL_CONE_OUTER_ANGLE
AL_REFERENCE_DISTANCE
*/
//void alGetSourcef(ALuint source,ALenum param,ALfloat *value);
func GetSourcef(source uint32, param Enum, value *float32) {
	getSourcef.Call(
		uintptr(source),
		uintptr(param),
		uintptr(unsafe.Pointer(value)))
}

//void alGetSourcefv(ALuint source,ALenum param,ALfloat *values);
func GetSourcefv(source uint32, param Enum, values unsafe.Pointer) {
	getSourcefv.Call(
		uintptr(source),
		uintptr(param),
		uintptr(values))
}

/*
pname the name of the attribute to retrieve:
 AL_SOURCE_RELATIVE
 AL_BUFFER
 AL_SOURCE_STATE
 AL_BUFFERS_QUEUED
 AL_BUFFERS_PROCESSED
*/
//void alGetSourcei(ALuint source,ALenum pname,ALint *value);
func GetSourcei(source uint32, param Enum, value *int32) {
	getSourcei.Call(
		uintptr(source),
		uintptr(param),
		uintptr(unsafe.Pointer(value)))
}

//void alGetSourceiv(ALuint source,ALenum param,ALint *values);
func GetSourceiv(source uint32, param Enum, value unsafe.Pointer) {
	getSourceiv.Call(
		uintptr(source),
		uintptr(param),
		uintptr(value))
}

//void alSourceUnqueueBuffers(ALuint source,ALsizei n,ALuint* buffers);
func SourceUnqueueBuffers(source uint32, n uint32, buffers []uint32) {
	sourceUnqueueBuffers.Call(
		uintptr(source),
		uintptr(n),
		uintptr(unsafe.Pointer(&buffers[0])))
}

/*param The property to be returned
AL_VENDOR
AL_VERSION
AL_RENDERER
AL_EXTENSIONS
*/
//const ALchar * alGetString(ALenum param);
func GetString(param Enum) string {
	r0, _, _ := getString.Call(uintptr(param))
	return C.GoString((*C.char)(unsafe.Pointer(r0)))
}

//ALboolean alIsExtensionPresent(const ALchar *extname);
func IsExtensionPresent(msg string) bool {
	cstr := unsafe.Pointer(C.CString(msg))
	defer C.free(cstr)
	r0, _, _ := isExtensionPresent.Call(uintptr(cstr))
	return r0 != 0
}

//param the name of the attribute to be set:
//AL_GAIN

//void alListenerf(ALenum param,ALfloat value);
func Listenerf(param Enum, value float32) {
	listenerf.Call(uintptr(param), uintptr(value))
}

//void alListenerfv(ALenum param,ALfloat *values);
func Listenerfv(param Enum, values unsafe.Pointer) {
	listenerfv.Call(uintptr(param), uintptr(values))
}

//void alListeneri(ALenum param,ALint value);
func Listeneri(param Enum, value int32) {
	listeneri.Call(
		uintptr(param),
		uintptr(value))
}

/*
param the name of the attribute to be set
 AL_POSITION
 AL_VELOCITY
 AL_ORIENTATION
*/
//void alListeneriv(ALenum param,ALint *values);
func Listeneriv(param Enum, values unsafe.Pointer) {
	listeneriv.Call(
		uintptr(param),
		uintptr(values))
}

/*
param the name of the attribute to set:
 AL_POSITION
 AL_VELOCITY
*/

//void alListener3i(ALenum param,ALint v1,ALint v2,ALint v3);
func Listener3i(param Enum, v1, v2, v3 int32) {
	listener3i.Call(
		uintptr(param),
		uintptr(v1),
		uintptr(v2),
		uintptr(v3))
}

//void alListener3f(ALenum param,ALfloat v1,ALfloat v2,ALfloat v3);
func Listener3f(param Enum, v1, v2, v3 float32) {
	listener3f.Call(
		uintptr(param),
		uintptr(v1),
		uintptr(v2),
		uintptr(v3))
}

//void alSpeedOfSound(ALfloat value);
func SpeedOfSound(value float32) {
	speedOfSound.Call(uintptr(value))
}
