#include <stdio.h>

char *buf(void) {
    static char b[32];     // lives for the entire program
    return b;              // legal to return this address
}

int main(void) {
    char *p = buf();
    for (int i = 0; i < 5; ++i) p[i] = 'A' + i;
    p[5] = '\0';
    puts(p);               // prints ABCDE
}