//		Copyright (c) 2019 Apfel
//
//		Permission is hereby granted, free of charge, to any person obtaining a copy
//		of this software and associated documentation files (the "Software"), to deal
//		in the Software without restriction, including without limitation the rights
//		to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//		copies of the Software, and to permit persons to whom the Software is
//		furnished to do so, subject to the following conditions:
//
//		The above copyright notice and this permission notice shall be included in all
//		copies or substantial portions of the Software.
//
//		THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//		IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//		FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//		AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//		LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//		OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//		SOFTWARE.
//
//		Fun fact: I never had so much fun, copying descriptions for stuff to make lint happy was never so nice for me.

package openhmd

//#cgo LDFLAGS: -L. -lopenhmd
//#include "./OpenHMD/include/openhmd.h"
import "C"

/// Types

// Context - An opaque pointer to a context structure.
type Context C.struct_ohmd_context

// Device - An opaque pointer to a structure representing a device, such as an HMD.
type Device C.struct_ohmd_device

// StringValue - A collection of string value information types.
type StringValue C.ohmd_string_value

// FloatValue - A collection of float value information types.
type FloatValue C.ohmd_float_value

// IntValue - A collection of int value information types.
type IntValue C.ohmd_int_value

/// Functions

// Create an OpenHMD context.
func Create() *Context {
	return C.ohmd_ctx_create()
}

// Destroy an OpenHMD context.
func Destroy(context *Context) {
	C.ohmd_ctx_destroy(C.struct_ohmd_context(context))
}

// GetError - Get the last error as a human readable string.
func GetError(context *Context) *C.char {
	return C.ohmd_ctx_get_error(C.struct_ohmd_context(context))
}

// Update a context.
func Update(context *Context) {
	C.ohmd_ctx_update(C.struct_ohmd_context(context))
}

// Probe for devices.
func Probe(context *Context) C.int {
	return C.ohmd_ctx_probe(C.struct_ohmd_context(context))
}

// ListGetString - Get device description from enumeration list index.
func ListGetString(context *Context, index C.int, value StringValue) *C.char {
	return C.ohmd_list_gets(C.struct_ohmd_context(context), index, C.ohmd_string_value(value))
}

// ListOpenDevice - Lists all opened Devices.
func ListOpenDevice(context *Context, index C.int) *C.ohmd_device {
	return C.ohmd_list_open_device(C.struct_ohmd_context(context), index)
}

// CloseDevice - Close a device.
func CloseDevice(device *Device) C.int {
	return C.ohmd_close_device(C.struct_ohmd_device(device))
}

// GetFloatDevice - Get a floating point value from a device.
func GetFloatDevice(device *Device, value FloatValue, out *C.float) C.int {
	return C.ohmd_device_getf(device, C.ohmd_float_value(value), out)
}

// SetFloatDevice - Set a floating point value for a device.
func SetFloatDevice(device *Device, value FloatValue, values *C.float) C.int {
	return C.ohmd_device_setf(C.struct_ohmd_device(device), C.ohmd_float_value(value), values)
}

// GetIntDevice - Get an integer value from a device.
func GetIntDevice(device *Device, value IntValue, out *C.int) C.int {
	return C.ohmd_device_geti(C.struct_ohmd_device(device), C.ohmd_int_value(value), out)
}
