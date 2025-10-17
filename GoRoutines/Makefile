.PHONY: help test test-verbose test-short bench bench-verbose race coverage clean run-demos run-channels

help:
	@echo "Available targets:"
	@echo "  test          - Run all tests"
	@echo "  test-verbose  - Run tests with verbose output"
	@echo "  test-short    - Run tests in short mode"
	@echo "  bench         - Run all benchmarks"
	@echo "  bench-verbose - Run benchmarks with verbose output"
	@echo "  race          - Run tests with race detector"
	@echo "  coverage      - Run tests with coverage report"
	@echo "  clean         - Clean test cache and coverage files"
	@echo "  run-demos     - Run all demo files"
	@echo "  run-channels  - Run all channel examples"

test:
	@echo "Running tests..."
	go test ./homework -v

test-verbose:
	@echo "Running tests with verbose output..."
	go test ./homework -v -count=1

test-short:
	@echo "Running tests in short mode..."
	go test ./homework -short

bench:
	@echo "Running benchmarks..."
	go test ./homework -bench=. -benchmem

bench-verbose:
	@echo "Running benchmarks with verbose output..."
	go test ./homework -bench=. -benchmem -v

bench-parallel:
	@echo "Running parallel sum benchmarks..."
	go test ./homework -bench=BenchmarkParallelSum -benchmem

bench-worker:
	@echo "Running worker pool benchmarks..."
	go test ./homework -bench=BenchmarkWorkerPool -benchmem

race:
	@echo "Running tests with race detector..."
	go test ./homework -race

coverage:
	@echo "Running tests with coverage..."
	go test ./homework -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

clean:
	@echo "Cleaning test cache and coverage files..."
	go clean -testcache
	rm -f coverage.out coverage.html

run-demos:
	@echo "Running demo files..."
	@echo "\n=== 01-goroutines ==="
	go run demos/01-goroutines.go
	@echo "\n=== 02-goroutines-anon ==="
	go run demos/02-goroutines-anon.go
	@echo "\n=== 03-mutex ==="
	go run demos/03-mutex.go

run-channels:
	@echo "Running channel examples..."
	@echo "\n=== 01-basics ==="
	go run channels/01-basics.go
	@echo "\n=== 02-buffered ==="
	go run channels/02-buffered.go
	@echo "\n=== 03-directions ==="
	go run channels/03-directions.go
	@echo "\n=== 04-ownership ==="
	go run channels/04-ownership.go
	@echo "\n=== 05-select ==="
	go run channels/05-select.go
	@echo "\n=== 06-for-select ==="
	go run channels/06-for-select.go
	@echo "\n=== 07-range ==="
	go run channels/07-range.go

# Individual test targets
test-task1:
	go test ./homework -run TestParallelSum -v

test-task2:
	go test ./homework -run TestFetchURLs -v

test-task3:
	go test ./homework -run TestProcessPipeline -v

test-task4:
	go test ./homework -run TestWorkerPool -v

# Run specific benchmark
bench-task1:
	go test ./homework -bench=BenchmarkParallelSum -benchmem -benchtime=3s

bench-task4:
	go test ./homework -bench=BenchmarkWorkerPool -benchmem -benchtime=3s