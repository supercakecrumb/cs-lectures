#include <stdio.h>    // printf
#include <stdlib.h>   // malloc, realloc, free
#include <string.h>   // strcpy, strlen

int main(void) {

    // char s[3]; ==
    char *s = stcpy(s, "Hello");

    // int len = 0;
    // for(len = 0; s[i] != 0; len++) {}
    // printf("len of \"Hello\" = %d\n", len);
    // char *s;
    // s = malloc(len * sizeof(char))
    // for(int i = 0; i < len; i++) {}

    strcpy(s, "Hello");
    printf("%s\n", s);

    return 0;
}