#include <stdio.h>

void foo(char a) {
    a = a + 1;
    printf("foo: a = '%c' (%d), &a = %p\n", a, (void*)&a);
}

void badzim_func(char badzim) {
    badzim = badzim + 1;
    printf("badzim_func: badzim = %d, &badzim = %p\n", badzim, (void*)&badzim);

    return
}

void bar(char* a) {
    *a = *a + 1;
    printf("bar: *a = %d, a = %p\n", *a, (void*)a);
}

main
c // 0x16eea273b
printf // 0x16eea2720
badzim_func
char badzim = c; // 0x16eea273a
badzim = badzim + 1;
printf("badzim_not_func: badzim = %d, &badzim = %p\n", badzim, (void*)&badzim);


..


int main(void) {
    char c = 'A'; // 0x1000 'A'
    // char d = 'B';
    
    
    printf("main: c = %d, &c = %p\n", c, (void*)&c);
    badzim_func(c);
    
    char badzim = c;
    badzim = badzim + 1;
    printf("badzim_not_func: badzim = %d, &badzim = %p\n", badzim, (void*)&badzim);

    printf("main: c = %d, &c = %p\n", c, (void*)&c);
    bar(&c);
    printf("main: c = %d, &c = %p\n", c, (void*)&c);
    return 0;
}


// 0x16d37a733 c = 'A'
// 0x16d37a732 d = 'B'
