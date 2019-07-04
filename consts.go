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

//#include <openhmd/openhmd.h>
//#cgo LDFLAGS: -lopenhmd
import "C"

// StringSize defines the maximum length of a string, including termination, in OpenHMD.
const StringSize = C.OHMD_STR_SIZE

const (
	// StatusCodeOkay returns that the operation went fine.
	StatusCodeOkay StatusCode = 0

	// StatusCodeUnknownError defines a return value that hasn't been specifically checked for.
	StatusCodeUnknownError StatusCode = -1

	// StatusCodeInvalidParameter gets returned if invalid parameters have been given.
	// Note that this is also used if too many entries are supplied, so GetFloat and GetInt, SetFloat and SetInt only return or accept 16 values.
	StatusCodeInvalidParameter StatusCode = -2

	// StatusCodeUnsupported defines either an unsupported version or a
	StatusCodeUnsupported StatusCode = -3

	// StatusCodeInvalidOperation defines a denied or invalid operation.
	StatusCodeInvalidOperation StatusCode = -4

	// StatusCodeUserReserved can be used for user-specific errors.
	StatusCodeUserReserved StatusCode = -16384
)

// I'm not sure if iota can still be used to define an "enum" if comments have been added in between

const (
	// StringValueVendor returns the name of the product's vendor.
	StringValueVendor StringValue = 1

	//StringValueProduct returns the name of the product itself.
	StringValueProduct StringValue = 2

	// StringValuePath returns the internal path of the device.
	StringValuePath StringValue = 3
)

// A collection of string descriptions, returning Shader-related data.
const (
	StringDescriptionGlslDisortionVertSrc StringDescription = iota
	StringDescriptionGlslDisortionFragSRC
	StringDescriptionGsls330DisortionVertSrc
	StringDescriptionGsls330DisortionFragSrc
	StringDescriptionGslsEsDisortionVertSrc
	StringDescriptionGslsEsDisortionFragSrc
)

const (
	// ControlHintGeneric - Generic button pressed.
	ControlHintGeneric ControlHint = 0

	// ControlHintTrigger - Trigger pushed.
	ControlHintTrigger ControlHint = 1

	// ControlHintTriggerClick - Trigger "clicked" - defines that the Trigger has been pushed all the way in.
	ControlHintTriggerClick ControlHint = 2

	// ControlHintSqueeze - Grip button pressed.
	ControlHintSqueeze ControlHint = 3

	// ControlHintMenu - Menu button pressed.
	ControlHintMenu ControlHint = 4

	// ControlHintHome - Home button pressed.
	ControlHintHome ControlHint = 5

	// ControlHintAnalogX - Horizontal stick movement.
	ControlHintAnalogX ControlHint = 6

	// ControlHintAnalogY - Vertical stick movement.
	ControlHintAnalogY ControlHint = 7

	// ControlHintAnalogPress - Stick pressed.
	ControlHintAnalogPress ControlHint = 8

	// ControlHintButtonA - Button A pressed.
	ControlHintButtonA ControlHint = 9

	// ControlHintButtonB - Button B pressed.
	ControlHintButtonB ControlHint = 10

	// ControlHintButtonX - Button X pressed.
	ControlHintButtonX ControlHint = 11

	// ControlHintButtonY - Button Y pressed.
	ControlHintButtonY ControlHint = 12

	// ControlHintVolumePlus - Volume up button pressed.
	ControlHintVolumePlus ControlHint = 13

	// ControlHintVolumeMinus - Volume down button pressed.
	ControlHintVolumeMinus ControlHint = 14

	// ControlHintMicMute - Microphone mute button pressed.
	ControlHintMicMute ControlHint = 15
)

const (
	// ControlTypeDigital defines digital controls, like a button.
	ControlTypeDigital ControlType = 0

	// ControlTypeAnalog defines analog controls, like a analog stick.
	ControlTypeAnalog ControlType = 1
)

const (
	// FloatValueRotationQuat defines the absolute rotation of the device, in space, as a quaternion.
	// Valid for GetFloat.
	// Returns 4 values: X, Y, Z, W.
	FloatValueRotationQuat FloatValue = 1

	// FloatValueLeftEyeGlModelViewMatrix defines a "ready to use" OpenGL style 4x4 matrix with a modelview matrix for the left eye of the HMD.
	// Valid for GetFloat.
	// Returns 16 values.
	FloatValueLeftEyeGlModelViewMatrix FloatValue = 2

	// FloatValueRightEyeGlModelViewMatrix defines a "ready to use" OpenGL style 4x4 matrix with a modelview matrix for the right eye of the HMD.
	// Valid for GetFloat.
	// Returns 16 values.
	FloatValueRightEyeGlModelViewMatrix FloatValue = 3

	// FloatValueLeftEyeGlProjectionMatrix defines a "ready to use" OpenGL style 4x4 matrix with a projection matrix for the left eye of the HMD.
	// Valid for GetFloat.
	// Returns 16 values.
	FloatValueLeftEyeGlProjectionMatrix FloatValue = 4

	// FloatValueRightEyeGlProjectionMatrix defines a "ready to use" OpenGL style 4x4 matrix with a projection matrix for the right eye of the HMD.
	// Valid for GetFloat.
	// Returns 16 values.
	FloatValueRightEyeGlProjectionMatrix FloatValue = 5

	// FloatValuePositionVector defines a 3D vector representing the absolute position of the device, in space.
	// Valid for GetFloat.
	// Returns 3 values: X, Y, Z.
	FloatValuePositionVector FloatValue = 6

	// FloatValueScreenHorizontalSize defines the width of the device's screen in metres.
	// Valid for GetFloat.
	// Returns 1 value.
	FloatValueScreenHorizontalSize FloatValue = 7

	// FloatValueScreenVerticalSize defines the height of the device's screen in metres.
	// Valid for GetFloat.
	// Returns 1 value.
	FloatValueScreenVerticalSize FloatValue = 8

	// FloatValueLensHorizontalSeparation defines the separation of the device's lenses in metres.
	// Valid for GetFloat.
	// Returns 1 value.
	FloatValueLensHorizontalSeparation FloatValue = 9

	// FloatValueLensVerticalPosition defines the vertical position of the device's lenses in metres.
	// Valid for GetFloat.
	// Returns 1 value.
	FloatValueLensVerticalPosition FloatValue = 10

	// FloatValueLeftEyeFOV defines the field of view for the left eye in degrees.
	// Valid for GetFloat.
	// Returns 1 value.
	FloatValueLeftEyeFOV FloatValue = 11

	// FloatValueLeftEyeAspectRatio defines the aspect ratio of the screen for the left eye.
	// Valid for GetFloat.
	// Returns 1 value.
	FloatValueLeftEyeAspectRatio FloatValue = 12

	// FloatValueRightEyeFOV defines the field of view for the right eye in degrees.
	// Valid for GetFloat.
	// Returns 1 value.
	FloatValueRightEyeFOV FloatValue = 13

	// FloatValueRightEyeAspectRatio defines the aspect ratio of the screen for the right eye.
	// Valid for GetFloat.
	// Returns 1 value.
	FloatValueRightEyeAspectRatio FloatValue = 14

	// FloatValueEyeIPD defines the the interpupillary distance of the user's eyes in metres.
	// Valid for GetFloat and SetFloat.
	// Returns/Accepts 1 value.
	FloatValueEyeIPD FloatValue = 15

	// FloatValueProjectionZFar defines how far the projection matrix can be drawn on the screen.
	// Valid for GetFloat and SetFloat.
	// Returns/Accepts 1 value.
	FloatValueProjectionZFar FloatValue = 16

	// FloatValueProjectionZNear defines how near the projection matrix can be drawn on the screen. This can be, for example, be used for close clipping distance.
	// Valid for GetFloat and SetFloat.
	// Returns/Accepts 1 value.
	FloatValueProjectionZNear FloatValue = 17

	// FloatValueDistortionK defines the device-specific distortion value.
	// Valid for GetFloat.
	// Returns 6 values.
	FloatValueDistortionK FloatValue = 18

	// FloatValueExternalSensorFusion defines the use of performing sensor fusion on values from the external sensors.
	// Valid for SetFloat.
	// Returns 10 values: Date + Time, X, Y, Z (gyrometer), X, Y, Z (accelerometer), X, Y, Z (magnetometer).
	FloatValueExternalSensorFusion FloatValue = 19

	// FloatValueUniversalDistortionK defines the universal shader distortion coefficients, based on the Panorama Tools model.
	// Valid for GetFloat.
	// Returns 4 values: A, B, C, D.
	FloatValueUniversalDistortionK FloatValue = 20

	// FloatValueUniversalAberrationK defines the universal shader aberration coefficients, based on Post-Warp scaling.
	// Valid for GetFloat.
	// Returns 3 values: R, G, B.
	FloatValueUniversalAberrationK FloatValue = 21

	// FloatValueControlsState defines the state of the device's controls.
	// Valid for GetFloat.
	// Returns between 1 to 16 values, based on IntValueControlsCount.
	FloatValueControlsState FloatValue = 22
)

const (
	// IntValueScreenHorizontalResolution defines the horizontal resolution of the device's screen.
	// Valid for GetInt.
	// Returns 1 value.
	IntValueScreenHorizontalResolution IntValue = 0

	// IntValueScreenVerticalResolution defines the vertical resolution of the device's screen.
	// Valid for GetInt.
	// Returns 1 value.
	IntValueScreenVerticalResolution IntValue = 1

	// IntValueDeviceClass returns the device's class, refer to DeviceClass for more information.
	// Valid For GetInt and ListGetInt.
	// Returns 1 value.
	IntValueDeviceClass IntValue = 2

	// IntValueDeviceFlags returns the device's flags, refer to DeviceFlags for more information.
	// Valid For GetInt and ListGetInt.
	// Returns 1 value.
	IntValueDeviceFlags IntValue = 3

	// IntValueControlsCount returns the count of controls for the device.
	// Valid for GetInt.
	// Returns 1 value.
	IntValueControlsCount IntValue = 4

	// IntValueControlsHints returns which buttons are supported.
	// Valid for GetInt.
	// Returns between 1 to 16 values, based on IntValueControlsCount.
	IntValueControlsHints IntValue = 5

	// IntValueControlsTypes returns whether a control entry is analog or digital.
	// Valid for GetInt.
	// Returns between 1 to 16 values, based on IntValueControlsCount.
	IntValueControlsTypes IntValue = 6
)

const (
	// DataValueDriverData defines the use of setting specific data for the internal drivers.
	DataValueDriverData DataValue = 0

	// DataValueDriverProperties defines the use of setting properties of a device internally.
	DataValueDriverProperties DataValue = 1
)

const (
	// IntSettingsIDsAutomaticUpdate allows OpenHMD to create background threads for automatic updating.
	// If this is set to 0, Update needs to be called at least 10 times per second.
	IntSettingsIDsAutomaticUpdate IntSettings = iota
)

const (
	// DeviceClassHMD defines a HMD.
	DeviceClassHMD DeviceClass = 0

	// DeviceClassController defines a motion controller.
	DeviceClassController DeviceClass = 1

	// DeviceClassGenericTracker defines a simple tracker.
	DeviceClassGenericTracker DeviceClass = 2
)

const (
	// DeviceFlagsNullDevice defines a dummy device.
	DeviceFlagsNullDevice DeviceFlags = 1

	// DeviceFlagsPositionalTracking defines a position-tracking device.
	DeviceFlagsPositionalTracking DeviceFlags = 2

	// DeviceFlagsRotationalTracking defines a rotation-tracking device.
	DeviceFlagsRotationalTracking DeviceFlags = 4

	// DeviceFlagsLeftController defines a left-sided motion controller.
	DeviceFlagsLeftController DeviceFlags = 8

	// DeviceFlagsRightController defines a right-sided motion controller.
	DeviceFlagsRightController DeviceFlags = 16
)
