#ifndef FLOATS_H
#define FLOATS_H

#include <openhmd/openhmd.h>

// GetFloat
int GetSingleFloatValue(ohmd_device* dev, ohmd_float_value value, float out);
int GetFloatArrayThree(ohmd_device* dev, ohmd_float_value value, float out[3]);
int GetFloatArrayFour(ohmd_device* dev, ohmd_float_value value, float out[4]);
int GetFloatArraySix(ohmd_device* dev, ohmd_float_value value, float out[6]);
int GetFloatArrayTen(ohmd_device* dev, ohmd_float_value value, float out[10]);
int GetFloatArraySixteen(ohmd_device* dev, ohmd_float_value value, float out[16]);

// SetFloat
int SetSingleFloatValue(ohmd_device* dev, ohmd_float_value value, float in);
int SetFloatArrayThree(ohmd_device* dev, ohmd_float_value value, float in[3]);
int SetFloatArrayFour(ohmd_device* dev, ohmd_float_value value, float in[4]);
int SetFloatArraySix(ohmd_device* dev, ohmd_float_value value, float in[6]);
int SetFloatArrayTen(ohmd_device* dev, ohmd_float_value value, float in[10]);
int SetFloatArraySixteen(ohmd_device* dev, ohmd_float_value value, float in[16]);

#endif