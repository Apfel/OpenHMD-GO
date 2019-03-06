package openhmd

/*
#include <openhmd/openhmd.h>
#include "floats.h"
#cgo LDFLAGS: -L. -lopenhmd
*/
import "C"

import (
	"bytes"
	"encoding/binary"
	"unsafe"
)

func ptoaf(pointer *C.float, len uintptr) []float32 {
	var output []float32
	size := int(unsafe.Sizeof(*pointer))
	length := int(len)
	gbytes := C.GoBytes(unsafe.Pointer(pointer), (C.int)(size*length))
	buf := bytes.NewBuffer(gbytes)

	for i := 0; i < length; i++ {
		var out float32
		if err := binary.Read(buf, binary.LittleEndian, &out); err == nil {
			output[i] = out
		}
	}

	return output
}

func getfloat(d *Device, value FloatValue, length ArraySize) (StatusCode, []float32) {
	switch length {
	case ArraySizeSingle:
		var value float32
		code := StatusCode(C.GetSingleFloatValue(d.c, C.ohmd_float_value(value), C.float(value)))
		return code, []float32{value}
	case ArraySizeThree:
		var array *C.float
		code := StatusCode(C.GetFloatArrayThree(d.c, C.ohmd_float_value(value), array))
		if code != StatusCodeOkay {
			return code, nil
		}
		farray := make([]float32, 3)
		for i, v := range ptoaf(array, unsafe.Sizeof(array)) {
			farray[i] = float32(v)
		}
		return code, farray
	case ArraySizeFour:
		var array *C.float
		code := StatusCode(C.GetFloatArrayFour(d.c, C.ohmd_float_value(value), array))
		if code != StatusCodeOkay {
			return code, nil
		}
		farray := make([]float32, 4)
		for i, v := range ptoaf(array, unsafe.Sizeof(array)) {
			farray[i] = float32(v)
		}
		return code, farray
	case ArraySizeSix:
		var array *C.float
		code := StatusCode(C.GetFloatArraySix(d.c, C.ohmd_float_value(value), array))
		if code != StatusCodeOkay {
			return code, nil
		}
		farray := make([]float32, 6)
		for i, v := range ptoaf(array, unsafe.Sizeof(array)) {
			farray[i] = float32(v)
		}
		return code, farray
	case ArraySizeTen:
		var array *C.float
		code := StatusCode(C.GetFloatArrayTen(d.c, C.ohmd_float_value(value), array))
		if code != StatusCodeOkay {
			return code, nil
		}
		farray := make([]float32, 10)
		for i, v := range ptoaf(array, unsafe.Sizeof(array)) {
			farray[i] = float32(v)
		}
		return code, farray
	case ArraySizeSixteen:
		var array *C.float
		code := StatusCode(C.GetFloatArraySixteen(d.c, C.ohmd_float_value(value), array))
		if code != StatusCodeOkay {
			return code, nil
		}
		farray := make([]float32, 16)
		for i, v := range ptoaf(array, unsafe.Sizeof(array)) {
			farray[i] = float32(v)
		}
		return code, farray
	default:
		return StatusCodeUnsupported, nil
	}
}

/*func (d *Device) setfloat(value FloatValue, input []float32) StatusCode {
	switch ArraySize(len(input)) {
	case ArraySizeSingle:
		return StatusCode(C.SetSingleFloatValue(d.c, C.ohmd_float_value(value), C.float(input[0])))
	case ArraySizeThree:
		var array [3]C.float
		for i, v := range input {
			array[i] = C.float(v)
		}
		return StatusCode(C.SetFloatArrayThree(d.c, C.ohmd_float_value(value), array))
	case ArraySizeFour:
		var array [4]C.float
		for i, v := range input {
			array[i] = C.float(v)
		}
		return StatusCode(C.SetFloatArrayFour(d.c, C.ohmd_float_value(value), array))
	case ArraySizeSix:
		var array [6]C.float
		for i, v := range input {
			array[i] = C.float(v)
		}
		return StatusCode(C.SetFloatArraySix(d.c, C.ohmd_float_value(value), array))
	case ArraySizeTen:
		var array [10]C.float
		for i, v := range input {
			array[i] = C.float(v)
		}
		return StatusCode(C.SetFloatArrayTen(d.c, C.ohmd_float_value(value), array))
	case ArraySizeSixteen:
		var array [16]C.float
		for i, v := range input {
			array[i] = C.float(v)
		}
		return StatusCode(C.SetFloatArraySixteen(d.c, C.ohmd_float_value(value), array))
	default:
		return StatusCodeUnsupported
	}
}*/
