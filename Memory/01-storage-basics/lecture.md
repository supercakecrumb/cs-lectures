## Part 0 · Bits & Bytes (foundations)

**Core facts**

* A **bit** is 0 or 1.  
* A **byte** is 8 bits and is the **smallest addressable unit** on mainstream CPUs (x86/x86‑64, ARM/ARM64). Each memory address names one **byte**.  
* A **word** is the CPU's natural register size (often 32 or 64 bits). It is **not** the smallest addressable unit.  
* Addresses count bytes; **types** decide how many bytes to read/write starting at that address.  

* x86/x86‑64 and ARM/ARM64 in practice: **little‑endian**. Networking uses **big‑endian** ("network byte order").  

---

### Alignment (just enough)

* Compilers prefer storing a 4‑byte `int` at an address divisible by 4, an 8‑byte `long`/`double` at a multiple of 8.  
* **x86** tolerates unaligned accesses (slower). Some architectures may fault or require special handling.  
* Structures may include **padding** to meet alignment; this affects `sizeof` and memory layout.  

---

### Signedness & ranges (two's complement)

* Unsigned 8‑bit: 0…255 (0x00…0xFF). Signed 8‑bit: −128…127.  
* Unsigned overflow wraps by definition; **signed overflow is undefined behavior** in C.  

---

### Units

* 1 **byte** = 8 bits  
* 1 **KiB** = 1024 bytes  
* 1 **MiB** = 1024 KiB  
* Use KiB/MiB/GiB for precision; KB/MB/GB are often used loosely.  

---

## Visual map (conceptual address space)

![alt text](memory.png)

``` mermaid
graph TB
  %% vertical stack (top = higher addresses)
  S[Stack (grows down)]
  H[Dynamic Data (Heap) (grows up)]
  SD[Static Data]
  L[Literals]
  I[Instructions]

  %% order top→bottom
  S --> H --> SD --> L --> I

  %% left-side permissions
  LS["writable; not executable"] --- S
  LH["writable; not executable"] --- H
  LSD["writable; not executable"] --- SD
  LL["read-only; not executable"] --- L
  LI["read-only; executable"] --- I

  %% right-side notes
  RS["managed automatically (by compiler)"] --- S
  RH["managed by programmer"] --- H
  RSD["initialized when process starts"] --- SD
  RL["initialized when process starts"] --- L
  RI["initialized when process starts"] --- I
```

> This is a **conceptual** layout; exact addresses change every run due to ASLR (address randomization). The relative roles remain the same.

---

## One-sentence definitions

* **Stack** — Automatic storage for each function call (locals, return addresses); super fast; limited; reclaimed when the function returns.  
* **Heap** — Dynamic storage you request with `malloc` and release with `free`; flexible size; slower; you own the lifetime.  
* **Static / Globals** — Variables that exist for the entire program (either file-scope globals or `static` variables inside functions).  
* **Mapped regions** — Memory areas the OS maps into your process (shared libraries, memory‑mapped files, and sometimes large allocations under the hood).  

---

## Tiny C example

```c
// globals.c (illustrative only)
#include <stdlib.h>

int g = 1;          // global → static storage (whole program lifetime)
static int sg = 2;  // file-scope static → static storage

int main(void) {
    static int sl = 3;      // function-scope static → static storage
    int local = 4;          // stack (lives until function returns)

    int *p = malloc(sizeof *p); // heap (lives until you free it)
    *p = 5;

    const char *s = "hi";  // pointer on stack; the literal lives in a mapped, read-only region

    free(p);
    return 0;
}
```

---

## Rules of thumb

* Prefer **stack** for small, short-lived data. It's automatic and fast.  
* Use the **heap** for big or flexible-sized data that must outlive the current function. Always `free` what you `malloc`.  
* Reserve **static/globals** for truly shared, long‑lived state or constants. Don't overuse them.  
* **Mapped regions** are mostly handled by the OS/runtime (shared libs). When you explicitly `mmap` a file, you treat it like an array in memory.  

---

## Summary

This overview covers the fundamental concepts of memory organization in C programs, including:
- Basic storage units (bits, bytes, words)
- Memory addressing and alignment
- Signed/unsigned representations
- The four main memory regions (stack, heap, static/globals, mapped regions)
- Practical guidelines for memory usage

Understanding these concepts is essential for writing efficient, reliable C programs and debugging memory-related issues.
