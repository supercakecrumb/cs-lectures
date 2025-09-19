// globals.c (illustrative only)
#include <stdlib.h>

int g = 1;          // global → static storage (whole program lifetime)
const int cg = 3;

int main(void) {
    static int sl = 3;      // function-scope static → static storage
    int local = 4;          // stack (lives until function returns)

    int *p = malloc(sizeof *p); // heap (lives until you free it)
    *p = 5;

    const char *s = "hi";  // pointer on stack; the literal lives in a mapped, read-only region

    free(p);
    return 0;
}