#include <stdio.h>

int *bad_ptr(int a) {
    int local = a;           // lives only inside bad_ptr
    return &local;           // BUG: returning address of a dead local
}

int *bad_ptr2(int a) {
    int local = a + 1;         // lives only inside bad_ptr
    return &local;           // BUG: returning address of a dead local
}

int main(void) {
    int *p = bad_ptr(7);
    // int *q = bad_ptr2(8);
    printf("*p = %d\n", *p); // undefined behavior
    bad_ptr(100);
    printf("%d\n", *p); // undefined behavior
    printf("%d\n", *p); // undefined behavior

    printf("p = %p\n", p);
    // printf("q = %p\n", q);
}