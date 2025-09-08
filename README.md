# Memory Management in C - Lecture Series

A comprehensive educational resource covering fundamental concepts of memory organization and management in C programming. This lecture series provides detailed explanations, visual diagrams, and practical code examples to help understand how memory works in modern computer systems.

## Overview

This repository contains a four-part lecture series that explores different aspects of memory in C programs:

1. **Storage Basics** - Fundamental concepts of bits, bytes, and memory organization
2. **Stack Memory** - Deep dive into stack frames and automatic storage
3. **Static Memory** - Understanding static storage duration and linkage
4. **Heap Memory** - Dynamic memory allocation and management

Each lecture includes:
- Conceptual explanations with visual diagrams
- Practical C code examples
- Build and run instructions
- Common pitfalls and best practices

## Lectures (English)

### Part 0: Storage Basics
[📖 Read Lecture](Memory/01-storage-basics/lecture.md)

Covers fundamental concepts including:
- Bits, bytes, and memory addressing
- Alignment and endianness
- Overview of memory regions (stack, heap, static, mapped)
- Basic memory layout in C programs

### Part 2: Stack - Deep Dive
[📖 Read Lecture](Memory/02-stack/lecture.md)

Explores stack memory in detail:
- Stack frames and function calls
- Stack growth direction
- Local variable storage
- Common stack-related bugs

### Part 3: Static Memory
[📖 Read Lecture](Memory/03-static/lecture.md)

Understanding static storage:
- Static storage duration
- File-scope vs function-scope static
- Internal linkage
- Use cases and best practices

### Part 4: Heap - Practical Basics
[📖 Read Lecture](Memory/04-heap/lecture.md)

Dynamic memory management:
- malloc/free fundamentals
- Memory ownership patterns
- Common heap errors
- realloc behavior

## Code Examples

The repository includes numerous C code examples demonstrating each concept:
- [`globals.c`](Memory/01-storage-basics/globals.c) - Basic memory regions example
- [`stack_frames.c`](Memory/02-stack/stack_frames.c) - Stack frame visualization
- [`stack_bytes.c`](Memory/02-stack/stack_bytes.c) - Raw stack memory inspection
- [`stack_bug.c`](Memory/02-stack/stack_bug.c) - Common stack lifetime bug
- [`static_counter.c`](Memory/03-static/static_counter.c) - Function-local static example
- [`hidden.c`](Memory/03-static/hidden.c) / [`main.c`](Memory/03-static/main.c) - File-scope static linkage
- [`static_table.c`](Memory/03-static/static_table.c) - Static lookup tables
- [`static_buf.c`](Memory/03-static/static_buf.c) - Returning static buffer
- [`heap_stability.c`](Memory/04-heap/heap_stability.c) - Heap vs stack address stability
- [`heap_data.c`](Memory/04-heap/heap_data.c) - Basic heap allocation
- [`heap_mistakes.c`](Memory/04-heap/heap_mistakes.c) - Common heap errors
- [`heap_realloc.c`](Memory/04-heap/heap_realloc.c) - realloc behavior

## Go Examples

Additional examples in Go are provided in the [`Memory/go`](Memory/go) directory to demonstrate similar concepts in a garbage-collected language.

---

# Управление памятью в C - Серия лекций

Комплексный образовательный ресурс, охватывающий фундаментальные концепции организации и управления памятью в программировании на C. Эта серия лекций предоставляет подробные объяснения, визуальные диаграммы и практические примеры кода для понимания работы памяти в современных компьютерных системах.

## Обзор

Этот репозиторий содержит серию из четырёх лекций, исследующих различные аспекты памяти в программах на C:

1. **Основы хранения** - Фундаментальные концепции битов, байтов и организации памяти
2. **Память стека** - Подробное изучение фреймов стека и автоматического хранения
3. **Статическая память** - Понимание статической продолжительности хранения и линковки
4. **Память кучи** - Динамическое выделение и управление памятью

Каждая лекция включает:
- Концептуальные объяснения с визуальными диаграммами
- Практические примеры кода на C
- Инструкции по сборке и запуску
- Распространённые ошибки и лучшие практики

## Лекции (Русский)

### Часть 0: Основы хранения
[📖 Читать лекцию](Memory/01-storage-basics/lecture-ru.md)

Охватывает фундаментальные концепции:
- Биты, байты и адресация памяти
- Выравнивание и порядок байтов
- Обзор областей памяти (стек, куча, статическая, отображённая)
- Базовая организация памяти в программах на C

### Часть 2: Стек - Подробно
[📖 Читать лекцию](Memory/02-stack/lecture-ru.md)

Детальное изучение памяти стека:
- Фреймы стека и вызовы функций
- Направление роста стека
- Хранение локальных переменных
- Распространённые ошибки, связанные со стеком

### Часть 3: Статическая память
[📖 Читать лекцию](Memory/03-static/lecture-ru.md)

Понимание статического хранения:
- Статическая продолжительность хранения
- static в области файла vs в области функции
- Внутренняя линковка
- Случаи использования и лучшие практики

### Часть 4: Куча - Практические основы
[📖 Читать лекцию](Memory/04-heap/lecture-ru.md)

Управление динамической памятью:
- Основы malloc/free
- Паттерны владения памятью
- Распространённые ошибки кучи
- Поведение realloc

## Примеры кода

Репозиторий включает множество примеров кода на C, демонстрирующих каждую концепцию:
- [`globals.c`](Memory/01-storage-basics/globals.c) - Базовый пример областей памяти
- [`stack_frames.c`](Memory/02-stack/stack_frames.c) - Визуализация фреймов стека
- [`stack_bytes.c`](Memory/02-stack/stack_bytes.c) - Инспекция сырой памяти стека
- [`stack_bug.c`](Memory/02-stack/stack_bug.c) - Распространённая ошибка времени жизни в стеке
- [`static_counter.c`](Memory/03-static/static_counter.c) - Пример static в функции
- [`hidden.c`](Memory/03-static/hidden.c) / [`main.c`](Memory/03-static/main.c) - Линковка static в области файла
- [`static_table.c`](Memory/03-static/static_table.c) - Статические таблицы поиска
- [`static_buf.c`](Memory/03-static/static_buf.c) - Возврат статического буфера
- [`heap_stability.c`](Memory/04-heap/heap_stability.c) - Стабильность адресов кучи vs стека
- [`heap_data.c`](Memory/04-heap/heap_data.c) - Базовое выделение в куче
- [`heap_mistakes.c`](Memory/04-heap/heap_mistakes.c) - Распространённые ошибки кучи
- [`heap_realloc.c`](Memory/04-heap/heap_realloc.c) - Поведение realloc

## Примеры на Go

Дополнительные примеры на Go предоставлены в директории [`Memory/go`](Memory/go) для демонстрации похожих концепций в языке со сборкой мусора.