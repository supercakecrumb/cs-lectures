#include <stdio.h>

static const char *keywords[] = { "if", "else", "for", "while", NULL };

int main(void) {
    for (int i = 0; keywords[i]; ++i) puts(keywords[i]);
}