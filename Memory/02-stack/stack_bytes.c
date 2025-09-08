#include <stdio.h>
#include <stdint.h>

static void hex(const void *p, size_t n) {
    const unsigned char *b = (const unsigned char*)p;
    for (size_t i = 0; i < n; i++) {
        printf("%02X%s", b[i], (i+1)%16?" ":"\n");
    }
    if (n % 16) puts("");
}

int main(void) {
    unsigned char a[32];
    unsigned char b[32];
    for (int i = 0; i < 32; i++) a[i] = 0xA0 + i; // pattern A0..BF
    for (int i = 0; i < 32; i++) b[i] = 0xB0 + i; // pattern B0..CF

    printf("&a=%p  &b=%p\n", (void*)a, (void*)b);
    puts("a bytes:"); hex(a, sizeof a);
    puts("b bytes:"); hex(b, sizeof b);
}