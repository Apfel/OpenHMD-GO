// Note: This is just a handler to make my life easier.

#include "handler.h"
#include <openhmd.h>

int getfloatsingle(ohmd_device* dev, ohmd_float_value type, float out) {
    return ohmd_device_getf(dev, type, &out);
}

int getfloatmulti(ohmd_device* dev, ohmd_float_value type, float out[]) {
    return ohmd_device_getf(dev, type, &out);
}


int setfloatsingle(ohmd_device* dev, ohmd_float_value type, float in) {
    return ohmd_device_setf(dev, type, &in);
}

int setfloatmulti(ohmd_device* dev, ohmd_float_value type, float in[]) {
    return ohmd_device_setf(dev, type, &in);
}


int getintsingle(ohmd_device* dev, ohmd_int_value type, int out) {
    return ohmd_device_geti(dev, type, &out);
}

int getintmulti(ohmd_device* dev, ohmd_int_value type, int out[]) {
    return ohmd_device_geti(dev, type, &out);
}


int setintsingle(ohmd_device* dev, ohmd_int_value type, int in) {
    return ohmd_device_seti(dev, type, &in);
}

int setintmulti(ohmd_device* dev, ohmd_int_value type, int in[]) {
    return ohmd_device_seti(dev, type, &in);
}

/* This is just too broken
int getstring(ohmd_string_description description, char* out) {
    const char** outt = "";
    int code = ohmd_gets(description, out);

    if (strncpy(outt, &out) != 0) return OHMD_S_UNKNOWN_ERROR;
    return code;
}*/