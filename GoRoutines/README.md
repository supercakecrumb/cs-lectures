# Go Concurrency Course / Курс по конкурентности в Go

[English](#english) | [Русский](#русский)

---

## English

### 📚 Course Overview

This is a comprehensive Go concurrency course covering goroutines, channels, and concurrent programming patterns. The course includes practical examples, homework assignments, and automated testing.

### 🏗️ Project Structure

```
GoRoutines/
├── README.md           # This file
├── Makefile           # Build automation and testing
├── go.mod             # Go module definition
├── demos/             # Basic goroutine examples
│   ├── 01-goroutines.go
│   ├── 02-goroutines-anon.go
│   └── 03-mutex.go
├── channels/          # Channel examples and patterns
│   ├── 01-basics.go
│   ├── 02-buffered.go
│   ├── 03-directions.go
│   ├── 04-ownership.go
│   ├── 05-select.go
│   ├── 06-for-select.go
│   ├── 07-range.go
│   └── README.md
└── homework/          # Assignments and tests
    ├── task1_parallel_sum.go
    ├── task2_http_fetch.go
    ├── task3_pipeline.go
    ├── task4_worker_pool.go
    ├── task5_rate_limiter.go
    ├── task6_fan_out_in.go
    ├── task7_timeout.go
    ├── task8_semaphore.go
    ├── errors.go
    ├── helpers.go
    └── tasks_test.go
```

### 🚀 Getting Started

1. **Prerequisites**: Go 1.19+ installed
2. **Clone and navigate**: `cd GoRoutines`
3. **Run examples**: `make run-demos` or `make run-channels`
4. **Test homework**: `make test`

### 🛠️ Makefile Commands

The project uses a [Makefile](https://www.gnu.org/software/make/manual/make.html) for automation. Make is a build automation tool that runs commands based on targets defined in the Makefile.

#### 📋 Available Commands

| Command | Description |
|---------|-------------|
| `make help` | Show all available commands |
| `make test` | Run all tests (expect failures for unimplemented functions) |
| `make test-verbose` | Run tests with detailed output |
| `make test-short` | Run tests in short mode |
| `make race` | Run tests with race condition detection |
| `make bench` | Run performance benchmarks |
| `make coverage` | Generate test coverage report (creates `coverage.html`) |
| `make clean` | Clean test cache and coverage files |
| `make run-demos` | Execute all demo files |
| `make run-channels` | Execute all channel examples |

#### 🎯 Individual Task Testing

| Command | Description |
|---------|-------------|
| `make test-task1` | Test only Task 1 (ParallelSum) |
| `make test-task2` | Test only Task 2 (FetchURLs) |
| `make test-task3` | Test only Task 3 (ProcessPipeline) |
| `make test-task4` | Test only Task 4 (WorkerPool) |

#### ⚡ Performance Testing

| Command | Description |
|---------|-------------|
| `make bench-task1` | Benchmark Task 1 (3 second runs) |
| `make bench-task4` | Benchmark Task 4 (3 second runs) |

### 📖 How Makefile Works

A Makefile contains **targets** (commands) and **recipes** (shell commands to execute). When you run `make target`, it executes the associated commands.

**Basic syntax:**
```makefile
target: dependencies
	command1
	command2
```

**Key features used in this project:**
- `.PHONY`: Declares targets that don't create files
- `@echo`: Prints messages (@ suppresses command echo)
- `&&`: Chains commands (stops on first failure)
- Variables: `./homework` specifies the test directory

**Example from our Makefile:**
```makefile
test:
	@echo "Running tests..."
	go test ./homework -v
```

Learn more: [GNU Make Manual](https://www.gnu.org/software/make/manual/make.html)

### 📝 Homework Tasks

The course includes 8 practical assignments:

1. **Task 1**: Parallel Sum - Calculate sum using multiple goroutines
2. **Task 2**: HTTP Fetcher - Concurrent HTTP requests with timeouts
3. **Task 3**: Pipeline - 3-stage processing pipeline (generate → square → filter)
4. **Task 4**: Worker Pool - Fixed number of workers processing jobs
5. **Task 5**: Rate Limiter - Process items with rate limiting
6. **Task 6**: Fan-Out/Fan-In - Distribute work and collect results
7. **Task 7**: Timeout Pattern - Processing with timeout constraints
8. **Task 8**: Semaphore - Limit concurrent operations

### ✅ Testing Strategy

- **Expected Behavior**: Tests will fail initially (functions return 0/nil)
- **Fast Execution**: Tests complete in ~0.2 seconds with timeout protection
- **Comprehensive Coverage**: Edge cases, boundary conditions, error scenarios
- **Race Detection**: Use `make race` to detect concurrency issues
- **Benchmarking**: Performance testing for optimization

### 🔧 Development Workflow

1. **Study examples**: `make run-demos && make run-channels`
2. **Implement functions**: Edit files in `homework/`
3. **Test your work**: `make test-task1` (replace with task number)
4. **Check for races**: `make race`
5. **Benchmark performance**: `make bench-task1`
6. **Generate coverage**: `make coverage`

---

## Русский

### 📚 Обзор курса

Это комплексный курс по конкурентности в Go, охватывающий горутины, каналы и паттерны конкурентного программирования. Курс включает практические примеры, домашние задания и автоматизированное тестирование.

### 🏗️ Структура проекта

```
GoRoutines/
├── README.md           # Этот файл
├── Makefile           # Автоматизация сборки и тестирования
├── go.mod             # Определение Go модуля
├── demos/             # Базовые примеры горутин
│   ├── 01-goroutines.go
│   ├── 02-goroutines-anon.go
│   └── 03-mutex.go
├── channels/          # Примеры каналов и паттернов
│   ├── 01-basics.go
│   ├── 02-buffered.go
│   ├── 03-directions.go
│   ├── 04-ownership.go
│   ├── 05-select.go
│   ├── 06-for-select.go
│   ├── 07-range.go
│   └── README.md
└── homework/          # Задания и тесты
    ├── task1_parallel_sum.go
    ├── task2_http_fetch.go
    ├── task3_pipeline.go
    ├── task4_worker_pool.go
    ├── task5_rate_limiter.go
    ├── task6_fan_out_in.go
    ├── task7_timeout.go
    ├── task8_semaphore.go
    ├── errors.go
    ├── helpers.go
    └── tasks_test.go
```

### 🚀 Начало работы

1. **Требования**: Go 1.19+ установлен
2. **Клонирование и переход**: `cd GoRoutines`
3. **Запуск примеров**: `make run-demos` или `make run-channels`
4. **Тестирование заданий**: `make test`

### 🛠️ Команды Makefile

Проект использует [Makefile](https://www.gnu.org/software/make/manual/make.html) для автоматизации. Make — это инструмент автоматизации сборки, который выполняет команды на основе целей, определенных в Makefile.

#### 📋 Доступные команды

| Команда | Описание |
|---------|----------|
| `make help` | Показать все доступные команды |
| `make test` | Запустить все тесты (ожидаются ошибки для нереализованных функций) |
| `make test-verbose` | Запустить тесты с подробным выводом |
| `make test-short` | Запустить тесты в кратком режиме |
| `make race` | Запустить тесты с обнаружением состояний гонки |
| `make bench` | Запустить бенчмарки производительности |
| `make coverage` | Сгенерировать отчет о покрытии тестами (создает `coverage.html`) |
| `make clean` | Очистить кеш тестов и файлы покрытия |
| `make run-demos` | Выполнить все демо-файлы |
| `make run-channels` | Выполнить все примеры каналов |

#### 🎯 Тестирование отдельных заданий

| Команда | Описание |
|---------|----------|
| `make test-task1` | Тестировать только Задание 1 (ParallelSum) |
| `make test-task2` | Тестировать только Задание 2 (FetchURLs) |
| `make test-task3` | Тестировать только Задание 3 (ProcessPipeline) |
| `make test-task4` | Тестировать только Задание 4 (WorkerPool) |

#### ⚡ Тестирование производительности

| Команда | Описание |
|---------|----------|
| `make bench-task1` | Бенчмарк Задания 1 (3-секундные прогоны) |
| `make bench-task4` | Бенчмарк Задания 4 (3-секундные прогоны) |

### 📖 Как работает Makefile

Makefile содержит **цели** (команды) и **рецепты** (shell-команды для выполнения). Когда вы запускаете `make target`, выполняются связанные команды.

**Базовый синтаксис:**
```makefile
target: dependencies
	command1
	command2
```

**Ключевые возможности, используемые в проекте:**
- `.PHONY`: Объявляет цели, которые не создают файлы
- `@echo`: Выводит сообщения (@ подавляет вывод команды)
- `&&`: Связывает команды (останавливается при первой ошибке)
- Переменные: `./homework` указывает директорию для тестов

**Пример из нашего Makefile:**
```makefile
test:
	@echo "Running tests..."
	go test ./homework -v
```

Подробнее: [Руководство GNU Make](https://www.gnu.org/software/make/manual/make.html)

### 📝 Домашние задания

Курс включает 8 практических заданий:

1. **Задание 1**: Параллельная сумма - Вычисление суммы с использованием нескольких горутин
2. **Задание 2**: HTTP-загрузчик - Конкурентные HTTP-запросы с таймаутами
3. **Задание 3**: Конвейер - 3-этапный конвейер обработки (генерация → возведение в квадрат → фильтрация)
4. **Задание 4**: Пул воркеров - Фиксированное количество воркеров, обрабатывающих задачи
5. **Задание 5**: Ограничитель скорости - Обработка элементов с ограничением скорости
6. **Задание 6**: Fan-Out/Fan-In - Распределение работы и сбор результатов
7. **Задание 7**: Паттерн таймаута - Обработка с ограничениями по времени
8. **Задание 8**: Семафор - Ограничение конкурентных операций

### ✅ Стратегия тестирования

- **Ожидаемое поведение**: Тесты изначально будут падать (функции возвращают 0/nil)
- **Быстрое выполнение**: Тесты завершаются за ~0.2 секунды с защитой от таймаутов
- **Комплексное покрытие**: Граничные случаи, пограничные условия, сценарии ошибок
- **Обнаружение гонок**: Используйте `make race` для обнаружения проблем конкурентности
- **Бенчмаркинг**: Тестирование производительности для оптимизации

### 🔧 Рабочий процесс разработки

1. **Изучите примеры**: `make run-demos && make run-channels`
2. **Реализуйте функции**: Редактируйте файлы в `homework/`
3. **Тестируйте работу**: `make test-task1` (замените на номер задания)
4. **Проверьте гонки**: `make race`
5. **Бенчмарк производительности**: `make bench-task1`
6. **Сгенерируйте покрытие**: `make coverage`

---

### 📚 Additional Resources / Дополнительные ресурсы

- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Effective Go - Concurrency](https://go.dev/doc/effective_go#concurrency)
- [Go Memory Model](https://go.dev/ref/mem)
- [GNU Make Manual](https://www.gnu.org/software/make/manual/make.html)

### 🤝 Contributing / Участие в разработке

Feel free to submit issues and pull requests to improve the course materials.

Не стесняйтесь отправлять issues и pull requests для улучшения материалов курса.