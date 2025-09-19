#include <stdio.h>

int main(void) {
    char letter = 97;
    printf("sizeof(char) = %d\n", sizeof(char));
    printf("address of char = %p\n", &letter);
    printf("value of letter = %c\n", letter);

    for( char i = 0; i < 127; i++) {
        printf("ASCII symbol number %d is %c\n", i, i);
    }

    return 0;
}