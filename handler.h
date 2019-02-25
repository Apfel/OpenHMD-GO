#ifndef HANDLER_H
#define HANDLER_H

#include <openhmd.h>

int getfloat(ohmd_device* dev, ohmd_float_value type, float out);
int getfloats(ohmd_device* dev, ohmd_float_value type, float out[]);

int setfloat(ohmd_device* dev, ohmd_float_value type, float in);
int setfloats(ohmd_device* dev, ohmd_float_value type, float in[]);

int getint(ohmd_device* dev, ohmd_int_value type, int out);
int getints(ohmd_device* dev, ohmd_int_value type, int out[]);

int setint(ohmd_device* dev, ohmd_int_value type, int in);
int setints(ohmd_device* dev, ohmd_int_value type, int in[]);

//int getstring(ohmd_string_description description, char* out);

#endif