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
#include <openhmd/openhmd.h>
#cgo LDFLAGS: -lopenhmd

float getfarray[16];
int getfloat(ohmd_device* device, ohmd_float_value value) { return ohmd_device_getf(device, value, getfarray); }
float getfvalue(int index) { return getfarray[index]; }

int getiarray[16];
int getint(ohmd_device* device, ohmd_int_value value) { return ohmd_device_geti(device, value, getiarray); }
int getivalue(int index) { return getiarray[index]; }

float setfarray[16];
int setfloat(ohmd_device* device, ohmd_float_value value) { return ohmd_device_setf(device, value, setfarray); }
void setfvalue(int index, float value) { setfarray[index] = value; }

int setiarray[16];
int setint(ohmd_device* device, ohmd_int_value value) { return ohmd_device_seti(device, value, setiarray); }
void setivalue(int index, int value) { setiarray[index] = value; }
*/
import "C"
import (
	"unsafe"
)

func getError(code statusCode) error {
	var err error

	switch code {
	case statusCodeOkay:
		err = nil
		break
	case statusCodeUnknownError:
		err = ErrorUnknownError
		break
	case statusCodeInvalidParameter:
		err = ErrorInvalidParameter
		break
	case statusCodeUnsupported:
		err = ErrorUnsupported
		break
	case statusCodeInvalidOperation:
		err = ErrorInvalidOperation
		break
	}

	return err
}

// GetString fetches a string description value from OpenHMD.
// This has not been implemented yet.
func GetString(desc StringDescription) (string, error) {
	return "", ErrorInvalidOperation
}

// CreateContext makes an OpenHMD context.
// Returns nil if the context can't allocate memory.
func CreateContext() *Context {
	ctx := C.ohmd_ctx_create()

	if ctx == nil {
		return nil
	}

	return &Context{c: ctx}
}

// Destroy removes the current OpenHMD context.
// Note: Your context will be removed from memory and all devices associated with the context will be closed automatically.
func (c *Context) Destroy() {
	C.ohmd_ctx_destroy(c.c)
}

// GetError fetches the last error as a human-readable string.
func (c *Context) GetError() string {
	return C.GoString(C.ohmd_ctx_get_error(c.c))
}

// Probe for devices.
// Returns the number of devices found on the system.
func (c *Context) Probe() int {
	return int(C.ohmd_ctx_probe(c.c))
}

// Update updates the data of the current context.
// This should be called periodically, to fetch the new values for the devices handled by a context.
func (c *Context) Update() {
	C.ohmd_ctx_update(c.c)
}

// CreateSettings creates a device settings instance.
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

// SetInt sets the given value for the providen key.
func (s *DeviceSettings) SetInt(key IntSettings, value int) error {
	code := statusCode(C.ohmd_device_settings_seti(s.c, C.ohmd_int_settings(key), &(C.int(value))))

	return getError(code)
}

// ListGetString gets a given device description from the enumeration list index.
func (c *Context) ListGetString(deviceIndex int, value StringValue) string {
	return C.GoString(C.ohmd_list_gets(c.c, C.int(deviceIndex), C.ohmd_string_value(value)))
}

// ListGetInt gets an integer value from the enumeration list index.
func (c *Context) ListGetInt(deviceIndex int, value IntValue) (int, error) {
	var val C.int
	err := getError(statusCode(C.ohmd_list_geti(c.c, C.int(deviceIndex), C.ohmd_int_value(value), &val)))
	return int(val), err
}

// ListOpenDevice opens a device.
func (c *Context) ListOpenDevice(index int) *Device {
	return &Device{c: C.ohmd_list_open_device(c.c, C.int(index))}
}

// ListOpenDeviceSettings opens a device with additional settings provided.
func (c *Context) ListOpenDeviceSettings(index int, settings *DeviceSettings) *Device {
	return &Device{c: C.ohmd_list_open_device_s(c.c, C.int(index), settings.c)}
}

// Close closes the current device.
func (d *Device) Close() error {
	return getError(statusCode(C.ohmd_close_device(d.c)))
}

// GetFloat fetches (a) float value(s).
func (d *Device) GetFloat(value FloatValue, length int) ([]float32, error) {
	if length > 16 || length < 1 {
		return nil, ErrorInvalidParameter
	}

	if err := getError(statusCode(C.getfloat(d.c, C.ohmd_float_value(value)))); err != nil {
		return nil, err
	}

	array := make([]float32, length)
	for i := 0; i != int(length); i++ {
		array[i] = float32(C.getfvalue(C.int(i)))
	}

	return array, nil
}

// SetFloat sets (a) float value(s).
func (d *Device) SetFloat(value FloatValue, input []float32) error {
	if len(input) > 16 {
		return ErrorInvalidParameter
	}

	for i, v := range input {
		C.setfvalue(C.int(i), C.float(v))
	}

	return getError(statusCode(C.setfloat(d.c, C.ohmd_float_value(value))))
}

// GetInt fetches (a) int value(s).
func (d *Device) GetInt(value IntValue, length int) ([]int32, error) {
	if length > 16 || length < 1 {
		return nil, ErrorInvalidParameter
	}

	if err := getError(statusCode(C.getint(d.c, C.ohmd_int_value(value)))); err != nil {
		return nil, err
	}

	array := make([]int32, length)
	for i := 0; i != int(length); i++ {
		array[i] = int32(C.getivalue(C.int(i)))
	}

	return array, nil
}

// SetData sets direct data for a device.
// BUG: This function seems to be broken, sending anything will end up with a SIGSEGV.
func (d *Device) SetData(value DataValue, input interface{}) error {
	if err := getError(statusCode(C.ohmd_device_set_data(d.c, C.ohmd_data_value(value), unsafe.Pointer(&input)))); err != nil {
		return err
	}

	return nil
}

// GetVersion returns OpenHMD's version.
func GetVersion() (int, int, int) {
	var major, minor, patch C.int
	C.ohmd_get_version(&major, &minor, &patch)

	return int(major), int(minor), int(patch)
}

// RequireVersion checks that the library is compatible with the required version.
// Returns true if the version is met.
func RequireVersion(major, minor, patch int) bool {
	if statusCode(C.ohmd_require_version(C.int(major), C.int(minor), C.int(patch))) == statusCodeUnsupported {
		return false
	}

	return true
}

// Sleep makes OpenHMD sleep for X seconds.
func Sleep(time float64) {
	C.ohmd_sleep(C.double(time))
}
