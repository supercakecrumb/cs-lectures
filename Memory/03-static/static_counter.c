#include <stdio.h>

void hit(void) {
    static int count = 0; // initialized once, persists across calls
    ++count;
    printf("hit count = %d\n", count);
}

int main(void) {
    hit(); hit(); hit(); // prints 1, then 2, then 3
}