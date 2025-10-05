#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <stddef.h>

typedef struct {
    int *data;
    size_t len;
    size_t cap;
} slice;

// sl := make([]int, 5) -> []int{0,0,0,0,0} (len = 5, cap = 5)
// sl := make([]int, 5, 10) -> []int{0,0,0,0,0,*,*,*,*,*} (len = 5, cap = 10) 
// fmt.Println(sl[5]) -> panic()
// sl := make([]int, 0, 10)
// sl := make([]type, len, cap)

// sl[0] 

// arr := malloc(10)
// for i := range len(arr):
//     arr[i] = 0

// arr := [4]int{1,2,3,4}
// arr := []int{1,2,3,4}
// arr = append(arr, 5)

slice slice_make(size_t len) {
    slice s;
    s.data = malloc(sizeof(int) * len);
    s.len = len;
    s.cap = len;
    for (int i = 0; i < len; i++) {
        s.data[i] = 0;
    }
    return s;
}

slice slice_make2(size_t len, size_t cap) {
    slice s;
    s.data = malloc(sizeof(int) * cap);
    s.len = len;
    s.cap = cap;
    for (int i = 0; i < len; i++) {
        s.data[i] = 0;
    }
    return s;
}

// append(s, 1, 2, 3, 4)
// arr := []int(1,2,3,4)
// append(s, arr...)

// sl := make([]int, 5, 10) -> []int{0,0,0,0,0,*,*,*,*,*} (len = 5, cap = 10) 
// sl = append(sl, 6)

bool append(slice *s, int badzim) {
    printf("append len = %d, cap = %d, badzim = %d\n", s->len, s->cap, badzim);
    // resize if needed
    if (s->len == s->cap) {
        printf("resizing slice\n");
        s->cap = s->cap * 2;
        int *new_data = malloc(sizeof(int) * s->cap);
        for (int i = 0; i < s->len; i++) {
            new_data[i] = s->data[i];
        }
        free(s->data);
        s->data = new_data;
        // s->cap = new_cap;
    }
    // append
    // s->data[s->len + 1] = badzim;
    // // idx = s->len + 1
    // // s->data[idx] = badzim
    // s->len = s->len + 1;

    // better append 
    s->data[s->len] = badzim;
    s->len = s->len + 1;
    printf("giving s->data[%d] value %d\n", s->len, badzim);

    return true;
}

// s := []int{0,1,2,3} len = 4 
// len(s) = 4
// s[3]
// s[len] = badzim
// len++

// get(s, 10)

// size_t == uint

// s = []int{0,1,2,3}
// get(s, 4, &out)

// out = s[idx]
bool get(slice *s, size_t idx, int *out) {
    if (idx >= s->len) {
        return false; // false - значит не вышло.
    }
    *out = s->data[idx];
    return true; // true - всё получилось!
}

// s[idx] = badzim
bool set(slice *s, size_t idx, int badzim) {
    if (idx >= s->len) {
        return false; // false - такого элемента нет.
    }
    s->data[idx] = badzim;
    return true;
}

int main(void) {

    // s := make([]int, 1, 2)
    // s = append(new_slice, 1)
    // s = append(new_slice, 2)
    // for i := range s {
    //    fmt.Println(s[i])
    // }

    slice s = slice_make2(1, 2);
    append(&s, 1);
    for (int i = 0; i < s.len; i++) {
        int out;
        bool success = get(&s, i, &out);

        printf("slice elemnt %d = %d\n", i, out);
    }
    append(&s, 2);
    for (int i = 0; i < s.len; i++) {
        int out;
        bool success = get(&s, i, &out);

        printf("slice elemnt %d = %d\n", i, out);
    }
    
    return 0;
}