/*
 * Memory Storage Demo
 * This program demonstrates different memory segments in C:
 *  - Static/Global memory
 *  - Stack memory 
 *  - Heap memory
 *  - Mapped regions (read-only)
 */

#include <stdlib.h>

/* Global variable - stored in static/global memory segment
 * - Allocated when program starts
 * - Exists for entire program lifetime
 * - Initialized before main() executes
 * - Shared across all functions in the file
 */
int g = 1;          

/* Static global variable - also in static memory
 * - Same lifetime as global
 * - Only visible within this file (file scope)
 */
static int sg = 2;  

int main(void) {
    /* Static local variable - still in static memory
     * - Retains value between function calls
     * - Only visible within this function
     */
    static int sl = 3;      

    /* Local variable - stored on the stack
     * - Allocated when function is called
     * - Automatically freed when function returns
     * - Limited size (typically ~8MB per thread)
     */
    int local = 4;          

    /* Heap allocation - memory requested from OS
     * - Size determined at runtime
     * - Must be manually freed with free()
     * - Accessible until explicitly freed
     * - Slower than stack but more flexible
     */
    int *p = malloc(sizeof *p); 
    *p = 5;

    /* String literal - stored in read-only mapped memory
     * - Pointer 's' is on the stack
     * - The actual string "hi" is in a read-only segment
     * - Attempting to modify it would cause a segmentation fault
     */
    const char *s = "hi";  

    /* Always free heap allocations to prevent memory leaks */
    free(p);

    return 0;
}

/*
 * Key Observations:
 * 1. Static/global vars: Persistent, shared state
 * 2. Stack vars: Automatic, function-scoped
 * 3. Heap vars: Manual management, flexible size  
 * 4. Mapped memory: Read-only, managed by OS
 */