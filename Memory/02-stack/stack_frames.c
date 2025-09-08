#include <stdio.h>

static void f2(int depth) {
    int x = 200 + depth;
    printf("f2: &x=%p\n", (void*)&x);
}

static void f1(int depth) {
    int y = 100 + depth;
    printf("f1: &y=%p\n", (void*)&y);
    f2(depth+1);
}

int main(void) {
    int z = 42;
    printf("main: &z=%p\n", (void*)&z);
    f1(0);
}