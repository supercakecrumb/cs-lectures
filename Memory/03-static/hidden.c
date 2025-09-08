static int hidden = 42;         // internal linkage: not visible outside this file
int public_value(void) { return hidden; }