#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void) {
    // String on the heap
    const char *src = "dynamic";
    size_t n = strlen(src) + 1;     // +1 for terminator
    char *s = malloc(n);
    memcpy(s, src, n);

    // Array on the heap
    size_t len = 5;
    int *a = malloc(len * sizeof *a);
    for (size_t i = 0; i < len; ++i) a[i] = (int)(10*i);

    printf("s='%s' at %p\n", s, (void*)s);
    printf("a at %p: ", (void*)a);
    for (size_t i = 0; i < len; ++i) printf("%d ", a[i]); puts("");

    free(a); free(s);
}