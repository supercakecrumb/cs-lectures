#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void) {
    char *p = malloc(8);
    strcpy(p, "hi");
    printf("p=%p '%s'\n", (void*)p, p);

    char *tmp = realloc(p, 4096); // may move
    if (!tmp) { free(p); return 1; }
    p = tmp;

    printf("p(now)=%p '%s'\n", (void*)p, p);
    free(p);
}