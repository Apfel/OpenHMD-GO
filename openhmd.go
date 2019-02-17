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

package openhmd

//#cgo LDFLAGS: -L. -lopenhmd
//#include "OpenHMD/include/openhmd.h"
import "C"

/// Types

// Context - An opaque pointer to a context structure.
type Context C.struct_ohmd_context

// Device - An opaque pointer to a structure representing a device, such as an HMD.
type Device C.struct_ohmd_device

// IntValue - A collection of int value information types.
type IntValue C.ohmd_int_value

// FloatValue - A collection of float value information types.
type FloatValue C.ohmd_float_value

// StringValue - A collection of string value information types.
type StringValue C.ohmd_string_value

// Float - A C-based Float.
type Float C.float

// Int - A C-based Integer.
type Int C.int

// String - A C-based char* in a string
type String C.char

// Create an OpenHMD context.
func Create() *Context {
	return C.ohmd_ctx_create()
}

/// Functions

// Destroy an OpenHMD context.
func Destroy(context *Context) {
	C.ohmd_ctx_destroy(context.(C.struct_ohmd_context))
}

// GetError - Get the last error as a human readable string.
func GetError(context *Context) *String {
	return C.ohmd_ctx_get_error(context.(C.struct_ohmd_context))
}

// Update a context.
func Update(context *Context) {
	C.ohmd_ctx_update(context.(C.struct_ohmd_context))
}

// Probe for devices.
func Probe(context *Context) Int {
	return C.ohmd_ctx_probe(context.(C.struct_ohmd_context))
}

// ListGetString - Get device description from enumeration list index.
func ListGetString(context *Context, index Int, value StringValue) *String {
	return C.ohmd_list_gets(context.(C.struct_ohmd_context), index.(C.int), value.(C.ohmd_string_value))
}

// ListOpenDevice - Lists all opened Devices.
func ListOpenDevice(context *Context, index Int) *Device {
	return C.ohmd_list_open_device(context.(C.struct_ohmd_context), index.(C.int))
}

// CloseDevice - Close a device.
func CloseDevice(device *Device) Int {
	return C.ohmd_close_device(device.(C.struct_ohmd_device))
}

// GetFloatDevice - Get a floating point value from a device.
func GetFloatDevice(device *Device, value FloatValue, out *Float) Int {
	return C.ohmd_device_getf(device.(C.struct_ohmd_device), value.(C.ohmd_float_value), out.(C.float))
}

// SetFloatDevice - Set a floating point value for a device.
func SetFloatDevice(device *Device, value FloatValue, values *Float) Int {
	return C.ohmd_device_setf(device.(C.struct_ohmd_device), value.(C.ohmd_float_value), values.(C.float))
}

// GetIntDevice - Get an integer value from a device.
func GetIntDevice(device *Device, value IntValue, out *Int) Int {
	return C.ohmd_device_geti(device.(C.struct_ohmd_device), value.(C.ohmd_int_value), out.(C.int))
}
