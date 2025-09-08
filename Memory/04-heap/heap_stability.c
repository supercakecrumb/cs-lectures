#include <stdio.h>
#include <stdlib.h>

static int *g; // points to heap memory

static void show(const char *who) {
    int stack_local = 0; // on this frame
    printf("%s: &stack_local=%p  g=%p -> %d\n", who, (void*)&stack_local, (void*)g, *g);
}

static void f1(void) { show("f1"); }
static void f2(void) { show("f2"); }

int main(void) {
    g = malloc(sizeof *g); *g = 1234;
    show("main"); f1(); f2();
    free(g); g = NULL;
}