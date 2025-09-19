# Part 2 · Stack — Deep Dive

**What it is**: per‑call storage for local variables, bookkeeping (saved registers), and the return address. Each thread has its own stack. On most systems the stack **grows toward lower addresses**.

---

---

## Demo A — Frames & addresses move downward

```c
// stack_frames.c
#include <stdio.h>

static void f2(int depth) {
    int x = 200 + depth;
    printf("f2: &x=%p\n", (void*)&x);
}

static void f1(int depth) {
    int y = 100 + depth;
    printf("f1: &y=%p\n", (void*)&y);
    f2(depth+1);
}

int main(void) {
    int z = 42;
    printf("main: &z=%p\n", (void*)&z);
    f1(0);
}
```

**Build & run**

```bash
gcc -O0 -g -Wall -Wextra -fno-omit-frame-pointer stack_frames.c -o stack_frames
./stack_frames
```

You’ll usually see addresses **decrease** as the call depth increases (Linux/x86‑64). Values vary run‑to‑run due to ASLR.


## Demo B — A classic lifetime bug (returning a stack address)

```c
// stack_bug.c
#include <stdio.h>

int *bad_ptr(void) {
    int local = 7;           // lives only inside bad_ptr
    return &local;           // BUG: returning address of a dead local
}

int main(void) {
    int *p = bad_ptr();
    printf("*p = %d\n", *p); // undefined behavior
}
```

**Build with sanitizers** (to catch it):

```bash
gcc -O0 -g -fsanitize=address,undefined -fno-omit-frame-pointer stack_bug.c -o stack_bug
./stack_bug
```

AddressSanitizer will flag a **use-after-return**. Takeaway: stack data dies when the function returns.

---

## Quick facts

* Typical per‑thread stack limit on Linux is a few MiB (e.g., 8 MiB); deep recursion or huge locals can overflow it.
* Each thread has its **own** stack. The heap is shared between threads.
* Locals declared `const` still live on the stack (they’re just read‑only by type, not by storage).
