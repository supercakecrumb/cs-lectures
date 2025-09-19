# Часть 2 · Стек — Подробно

**Что это такое**: хранилище для локальных переменных, сохранённых регистров и адреса возврата для каждого вызова функции. Каждый поток имеет свой собственный стек. В большинстве систем стек **растёт в сторону младших адресов**.

---

## Визуализация: фреймы в стеке

```mermaid
graph TB
  HA[Старшие адреса]
  LA[Младшие адреса]

  subgraph Stack
    direction BT
    F2[Фрейм f2() — локальные, сохранённые регистры]
    F1[Фрейм f1()]
    FM[Фрейм main()]
  end

  HA --- Stack
  Stack --- LA
```

> Новые вызовы добавляют новый **фрейм**; возврат из функции удаляет этот фрейм.

---

## Демо A · Фреймы и адреса смещаются вниз

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

**Сборка и запуск**

```bash
gcc -O0 -g -Wall -Wextra -fno-omit-frame-pointer stack_frames.c -o stack_frames
./stack_frames
```

Обычно вы увидите, что адреса **уменьшаются** по мере увеличения глубины вызовов (Linux/x86-64). Значения меняются от запуска к запуску из-за ASLR.

---

## Демо B · Просмотр сырых байтов в стеке (безопасно)

```c
// stack_bytes.c
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
    for (int i = 0; i < 32; i++) a[i] = 0xA0 + i; // паттерн A0..BF
    for (int i = 0; i < 32; i++) b[i] = 0xB0 + i; // паттерн B0..CF

    printf("&a=%p  &b=%p\n", (void*)a, (void*)b);
    puts("a bytes:"); hex(a, sizeof a);
    puts("b bytes:"); hex(b, sizeof b);
}
```

**Сборка и запуск**

```bash
gcc -O0 -g -Wall -Wextra stack_bytes.c -o stack_bytes
./stack_bytes
```

Часто `&a` будет **выше**, чем `&b`, потому что `b` был выделен позже, а стек растёт **вниз**. Hex-дамп показывает точное содержимое памяти.

---

## Демо C · Классическая ошибка времени жизни (возврат адреса в стеке)

```c
// stack_bug.c
#include <stdio.h>

int *bad_ptr(void) {
    int local = 7;           // существует только внутри bad_ptr
    return &local;           // ОШИБКА: возврат адреса локальной переменной
}

int main(void) {
    int *p = bad_ptr();
    printf("*p = %d\n", *p); // неопределённое поведение
}
```

**Сборка с санитайзерами** (для обнаружения ошибки):

```bash
gcc -O0 -g -fsanitize=address,undefined -fno-omit-frame-pointer stack_bug.c -o stack_bug
./stack_bug
```

AddressSanitizer сообщит об **использовании после возврата** (use-after-return). Вывод: данные в стеке уничтожаются при возврате из функции.

---

## Краткие факты

* Типичный лимит стека на поток в Linux составляет несколько МиБ (например, 8 МиБ); глубокая рекурсия или большие локальные переменные могут вызвать переполнение.
* Каждый поток имеет **свой** стек. Куча разделяется между потоками.
* Локальные переменные, объявленные как `const`, всё равно находятся в стеке (они доступны только для чтения по типу, а не по месту хранения).