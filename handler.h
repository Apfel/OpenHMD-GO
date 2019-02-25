#ifndef HANDLER_H
#define HANDLER_H

#include <openhmd.h>

int getfloatsingle(ohmd_device* dev, ohmd_float_value type, float out);
int getfloatmulti(ohmd_device* dev, ohmd_float_value type, float* out);

int setfloatsingle(ohmd_device* dev, ohmd_float_value type, float in);
int setfloatmulti(ohmd_device* dev, ohmd_float_value type, float* in);

int getintsingle(ohmd_device* dev, ohmd_int_value type, int out);
int getintmulti(ohmd_device* dev, ohmd_int_value type, int* out);

int setintsingle(ohmd_device* dev, ohmd_int_value type, int in);
int setintmulti(ohmd_device* dev, ohmd_int_value type, int* in);

int getstring(ohmd_string_description description, char* out);

#endif