#include <stdio.h>

int *bad_ptr(void) {
    int local = 7;           // lives only inside bad_ptr
    return &local;           // BUG: returning address of a dead local
}

int main(void) {
    int *p = bad_ptr();
    printf("*p = %d\n", *p); // undefined behavior
}