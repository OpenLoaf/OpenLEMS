#include "include/hexlib.h"

extern "C" {

int hex_add(int a, int b) {
    return a + b;
}

const char* hex_version(void) {
    return "0.1.0-mpc";
}

} // extern "C"


