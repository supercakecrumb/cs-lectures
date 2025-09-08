#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void) {
    char *p = malloc(8);
    strcpy(p, "this is too long"); // ❌ overflow (writes past 8)

    int *q = malloc(sizeof *q);
    free(q);
    // printf("%d\n", *q);        // ❌ use-after-free (commented out to run safely)

    // free(q);                    // ❌ double free (don’t do this)

    free(p);
}