// Boost Software License - Version 1.0 - August 17th, 2003
//
// Permission is hereby granted, free of charge, to any person or organization
// obtaining a copy of the software and accompanying documentation covered by
// this license (the "Software") to use, reproduce, display, distribute,
// execute, and transmit the Software, and to prepare derivative works of the
// Software, and to permit third-parties to whom the Software is furnished to
// do so, all subject to the following:
//
// The copyright notices in the Software and this entire statement, including
// the above license grant, this restriction and the following disclaimer,
// must be included in all copies of the Software, in whole or in part, and
// all derivative works of the Software, unless such copies or derivative
// works are solely in the form of machine-executable object code generated by
// a source language processor.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE, TITLE AND NON-INFRINGEMENT. IN NO EVENT
// SHALL THE COPYRIGHT HOLDERS OR ANYONE DISTRIBUTING THE SOFTWARE BE LIABLE
// FOR ANY DAMAGES OR OTHER LIABILITY, WHETHER IN CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package openhmd

/*
int iarray[16];
int getint(ohmd_device* device, ohmd_int_value value) { return ohmd_device_geti(device, value, iarray); }
int getivalue(int index) { return iarray[index]; }
*/

/*
#include <openhmd/openhmd.h>
#cgo LDFLAGS: -L. -lopenhmd

float farray[16];
int getfloat(ohmd_device* device, ohmd_float_value value) { return ohmd_device_getf(device, value, farray); }
float getfvalue(int index) { return farray[index]; }

*/
import "C"

import "unsafe"

// Create makes an OpenHMD context.
// Returns nil if the context can't allocate memory.
func Create() *Context {
	ctx := C.ohmd_ctx_create()

	if ctx == nil {
		return nil
	}

	return &Context{c: ctx}
}

// Destroy removes the current OpenHMD context.
// Note: Your context will be nulled and all devices associated with the context are automatically closed.
func (c *Context) Destroy() {
	C.ohmd_ctx_destroy(c.c)
}

// GetError gets the last error as a human readable string.
func (c *Context) GetError() string {
	return C.GoString(C.ohmd_ctx_get_error(c.c))
}

// Update updates the current context
// to fetch the values for the devices handled by a context.
func (c *Context) Update() {
	C.ohmd_ctx_update(c.c)
}

// Probe for devices.
// Returns the number of devices found on the system.
func (c *Context) Probe() int {
	return int(C.ohmd_ctx_probe(c.c))
}

// GetString fetches a string from OpenHMD.
func GetString(desc StringDescription) (StatusCode, string) {
	var value string
	code := StatusCodeUnsupported //StatusCode(C.ohmd_gets(C.ohmd_string_description(desc), (**C.char)(unsafe.Pointer(&value))))
	return code, value
}

// ListGetString gets device description from enumeration list index.nt)
func (c *Context) ListGetString(deviceIndex int, value StringValue) string {
	return C.GoString(C.ohmd_list_gets(c.c, C.int(deviceIndex), C.ohmd_string_value(value)))
}

// ListGetInt gets integer value from enumeration list index.
func (c *Context) ListGetInt(deviceIndex int, value IntValue) (StatusCode, int) {
	var val C.int
	code := StatusCode(C.ohmd_list_geti(c.c, C.int(deviceIndex), C.ohmd_int_value(value), &val))
	return code, int(val)
}

// ListOpenDevice opens a device.
func (c *Context) ListOpenDevice(index int) *Device {
	return &Device{c: C.ohmd_list_open_device(c.c, C.int(index))}
}

// ListOpenDeviceSettings opens a device with additional settings provided.
func (c *Context) ListOpenDeviceSettings(index int, settings DeviceSettings) *Device {
	return &Device{c: C.ohmd_list_open_device_s(c.c, C.int(index), settings.c)}
}

// CreateSettings creates a device settings instance.nt)
func (c *Context) CreateSettings() *DeviceSettings {
	settings := C.ohmd_device_settings_create(c.c)

	if settings == nil {
		return nil
	}

	return &DeviceSettings{c: settings}
}

// Destroy destroys a device settings instance.
func (s *DeviceSettings) Destroy() {
	C.ohmd_device_settings_destroy(s.c)
}

// Close closes the current device.
func (d *Device) Close() StatusCode {
	return StatusCode(C.ohmd_close_device(d.c))
}

// GetFloat fetches (a) float value(s).
func (d *Device) GetFloat(value FloatValue, length ArraySize) (StatusCode, []float32) {
	code := StatusCode(C.getfloat(d.c, C.ohmd_float_value(value)))

	/*if code != StatusCodeOkay {
		return code, nil
	}*/

	array := make([]float32, length)
	for i := 0; i != int(length); i++ {
		array[i] = float32(C.getfvalue(C.int(i)))
	}
	return code, array
}

// SetFloat sets (a) float value(s).
func (d *Device) SetFloat(value FloatValue, input []float32) StatusCode {
	return StatusCodeOkay
}

// GetInt fechtes (a) int value(s).
func (d *Device) GetInt(value IntValue, length ArraySize) (StatusCode, []int32) {
	/*code := StatusCode(C.getint(d.c, C.ohmd_int_value(value)))

	if code != StatusCodeOkay {
		return code, nil
	}

	array := make([]int32, length)
	for i := 0; i != int(length); i++ {
		array[i] = int32(C.getivalue(C.int(i)))
	}
	return code, array*/
	return StatusCodeOkay, nil
}

// SetInt sets (a) int value(s).
func (d *Device) SetInt(value IntValue, input []int) StatusCode {
	return StatusCodeOkay
}

// SetData sets direct data for a device.
func (d *Device) SetData(value DataValue, input *interface{}) StatusCode {
	return StatusCode(C.ohmd_device_set_data(d.c, C.ohmd_data_value(value), unsafe.Pointer(input)))
}

// GetVersion fetches OpenHMD's version.
func GetVersion() (int, int, int) {
	var major, minor, patch C.int
	C.ohmd_get_version(&major, &minor, &patch)

	return int(major), int(minor), int(patch)
}

// RequireVersion checks that the library is compatible with the required version.
// Returns StatusOkay if compatible, else StatusUnsupported.
func RequireVersion(major, minor, patch int) StatusCode {
	return StatusCode(C.ohmd_require_version(C.int(major), C.int(minor), C.int(patch)))
}
