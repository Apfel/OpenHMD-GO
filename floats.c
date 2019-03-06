#include <openhmd/openhmd.h>
#include "floats.h"

// GetFloat

int GetSingleFloatValue(ohmd_device* dev, ohmd_float_value value, float out) {
    float values[1];
    int code = ohmd_device_getf(dev, value, values);
    out = values[0];
    return code;
}

int GetFloatArrayThree(ohmd_device* dev, ohmd_float_value value, float out[3]) {
    return ohmd_device_getf(dev, value, out);
}

int GetFloatArrayFour(ohmd_device* dev, ohmd_float_value value, float out[4]) {
    return ohmd_device_getf(dev, value, out);
}

int GetFloatArraySix(ohmd_device* dev, ohmd_float_value value, float out[6]) {
    return ohmd_device_getf(dev, value, out);
}

int GetFloatArrayTen(ohmd_device* dev, ohmd_float_value value, float out[10]) {
    return ohmd_device_getf(dev, value, out);
}

int GetFloatArraySixteen(ohmd_device* dev, ohmd_float_value value, float out[16]) {
    return ohmd_device_getf(dev, value, out);
}

// SetFloat

int SetSingleFloatValue(ohmd_device* dev, ohmd_float_value value, float in) {
    float values[1];
    int code = ohmd_device_getf(dev, value, values);
    in = values[0];
    return code;
}

int SetFloatArrayThree(ohmd_device* dev, ohmd_float_value value, float in[3]) {
    return ohmd_device_getf(dev, value, in);
}

int SetFloatArrayFour(ohmd_device* dev, ohmd_float_value value, float in[4]) {
    return ohmd_device_getf(dev, value, in);
}

int SetFloatArraySix(ohmd_device* dev, ohmd_float_value value, float in[6]) {
    return ohmd_device_getf(dev, value, in);
}

int SetFloatArrayTen(ohmd_device* dev, ohmd_float_value value, float in[10]) {
    return ohmd_device_getf(dev, value, in);
}

int SetFloatArraySixteen(ohmd_device* dev, ohmd_float_value value, float in[16]) {
    return ohmd_device_getf(dev, value, in);
}