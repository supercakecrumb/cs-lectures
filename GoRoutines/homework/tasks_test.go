package homework

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// Task 1: ParallelSum Tests
func TestParallelSum(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		workers  int
		expected int
	}{
		{"empty slice", []int{}, 2, 0},
		{"nil slice", nil, 2, 0},
		{"single element", []int{5}, 1, 5},
		{"small slice", []int{1, 2, 3, 4, 5}, 2, 15},
		{"large slice", makeRange(1, 1000), 4, 500500},
		{"negative numbers", []int{-5, -3, 2, 4}, 2, -2},
		{"all zeros", []int{0, 0, 0, 0}, 2, 0},
		{"mixed positive negative", []int{-10, 5, -3, 8, -2, 10}, 3, 8},
		{"single worker", []int{1, 2, 3, 4, 5}, 1, 15},
		{"more workers than elements", []int{1, 2, 3}, 10, 6},
		{"zero workers", []int{1, 2, 3}, 0, 0},
		{"negative workers", []int{1, 2, 3}, -1, 0},
		{"uneven distribution", []int{1, 2, 3, 4, 5, 6, 7}, 3, 28},
		{"large numbers", []int{1000000, 2000000, 3000000}, 2, 6000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParallelSum(tt.numbers, tt.workers)
			if result != tt.expected {
				t.Errorf("ParallelSum() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestSquareSum(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		workers  int
		expected int
	}{
		{"empty", []int{}, 2, 0},
		{"nil slice", nil, 2, 0},
		{"single", []int{3}, 1, 9},
		{"multiple", []int{1, 2, 3}, 2, 14}, // 1 + 4 + 9
		{"negative", []int{-2, 3}, 2, 13},   // 4 + 9
		{"zeros", []int{0, 0, 0}, 2, 0},
		{"large numbers", []int{100, 200}, 2, 50000},      // 10000 + 40000
		{"mixed signs", []int{-5, 0, 3, -2}, 2, 38},       // 25 + 0 + 9 + 4
		{"single worker", []int{1, 2, 3, 4}, 1, 30},       // 1 + 4 + 9 + 16
		{"more workers than elements", []int{1, 2}, 5, 5}, // 1 + 4
		{"zero workers", []int{1, 2, 3}, 0, 0},
		{"negative workers", []int{1, 2, 3}, -1, 0},
		{"uneven distribution", []int{1, 2, 3, 4, 5}, 3, 55}, // 1 + 4 + 9 + 16 + 25
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SquareSum(tt.numbers, tt.workers)
			if result != tt.expected {
				t.Errorf("SquareSum() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func BenchmarkParallelSum(b *testing.B) {
	numbers := makeRange(1, 10000)

	benchmarks := []struct {
		name    string
		workers int
	}{
		{"1worker", 1},
		{"2workers", 2},
		{"4workers", 4},
		{"8workers", 8},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ParallelSum(numbers, bm.workers)
			}
		})
	}
}

// Task 2: FetchURLs Tests
func TestFetchURLs(t *testing.T) {
	tests := []struct {
		name     string
		urls     []string
		timeout  time.Duration
		wantErr  bool
		minCodes int
	}{
		{"empty urls", []string{}, 1 * time.Second, false, 0},
		{"nil urls", nil, 1 * time.Second, false, 0},
		{"single url", []string{"http://example.com"}, 1 * time.Second, false, 1},
		{"multiple urls", []string{"http://example.com", "http://google.com"}, 2 * time.Second, false, 2},
		{"with timeout", []string{"http://example.com"}, 1 * time.Millisecond, true, 0},
		{"invalid url", []string{"http://invalid.url.that.does.not.exist"}, 1 * time.Second, false, 0},
		{"mixed valid/invalid", []string{"http://example.com", "http://invalid.url.that.does.not.exist"}, 2 * time.Second, false, 1},
		{"duplicate urls", []string{"http://example.com", "http://example.com"}, 2 * time.Second, false, 2},
		{"very short timeout", []string{"http://example.com"}, 1 * time.Nanosecond, true, 0},
		{"long timeout", []string{"http://example.com"}, 10 * time.Second, false, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FetchURLs(tt.urls, tt.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchURLs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(result) < tt.minCodes {
				t.Errorf("FetchURLs() got %d results, want at least %d", len(result), tt.minCodes)
			}
		})
	}
}

func BenchmarkFetchURLs(b *testing.B) {
	urls := []string{
		"http://example.com",
		"http://google.com",
		"http://github.com",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FetchURLs(urls, 5*time.Second)
	}
}

// Task 3: ProcessPipeline Tests
func TestProcessPipeline(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{"n=0", 0, []int{}},
		{"n=1", 1, []int{}},      // 1 is odd, so no output
		{"n=2", 2, []int{4}},     // 1 (odd), 2^2=4 (even)
		{"n=5", 5, []int{4, 16}}, // squares: 1,4,9,16,25 -> even: 4,16
		{"n=10", 10, []int{4, 16, 36, 64, 100}},
		{"n=15", 15, []int{4, 16, 36, 64, 100, 144, 196}},
		{"n=20", 20, []int{4, 16, 36, 64, 100, 144, 196, 256, 324, 400}},
		{"n negative", -5, []int{}},
		{"n large", 100, makeEvenSquares(100)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := ProcessPipeline(tt.n)
			if ch == nil {
				t.Skip("ProcessPipeline not implemented yet")
				return
			}

			results := []int{}
			// Add timeout to prevent infinite blocking
			done := make(chan bool)
			go func() {
				for num := range ch {
					results = append(results, num)
				}
				done <- true
			}()

			select {
			case <-done:
				// Success
			case <-time.After(1 * time.Second):
				t.Error("ProcessPipeline() timed out - possible deadlock or infinite loop")
				return
			}

			if len(results) != len(tt.expected) {
				t.Errorf("ProcessPipeline() got %d results, want %d", len(results), len(tt.expected))
				return
			}

			for i, v := range results {
				if v != tt.expected[i] {
					t.Errorf("ProcessPipeline() result[%d] = %d, want %d", i, v, tt.expected[i])
				}
			}
		})
	}
}

// Helper function to generate expected even squares
func makeEvenSquares(n int) []int {
	result := []int{}
	for i := 1; i <= n; i++ {
		square := i * i
		if square%2 == 0 {
			result = append(result, square)
		}
	}
	return result
}

func BenchmarkProcessPipeline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := ProcessPipeline(100)
		if ch == nil {
			b.Skip("ProcessPipeline not implemented yet")
			return
		}

		// Add timeout protection for benchmark
		done := make(chan bool)
		go func() {
			for range ch {
			}
			done <- true
		}()

		select {
		case <-done:
			// Success
		case <-time.After(1 * time.Second):
			b.Fatal("ProcessPipeline() benchmark timed out")
		}
	}
}

// Task 4: WorkerPool Tests
func TestWorkerPool(t *testing.T) {
	tests := []struct {
		name       string
		jobs       []int
		numWorkers int
		expected   []int
	}{
		{"empty jobs", []int{}, 2, []int{}},
		{"nil jobs", nil, 2, []int{}},
		{"single job", []int{5}, 1, []int{10}},
		{"multiple jobs", []int{1, 2, 3, 4}, 2, []int{2, 4, 6, 8}},
		{"more workers than jobs", []int{1, 2}, 5, []int{2, 4}},
		{"zero workers", []int{1, 2, 3}, 0, []int{}},
		{"negative workers", []int{1, 2, 3}, -1, []int{}},
		{"large jobs", []int{100, 200, 300}, 2, []int{200, 400, 600}},
		{"negative jobs", []int{-1, -2, 3}, 2, []int{-2, -4, 6}},
		{"mixed jobs", []int{0, 1, -1, 5}, 3, []int{0, 2, -2, 10}},
		{"many workers few jobs", []int{1}, 100, []int{2}},
		{"many jobs few workers", makeRange(1, 50), 3, makeDoubledRange(1, 50)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := WorkerPool(tt.jobs, tt.numWorkers)

			if len(results) != len(tt.expected) {
				t.Errorf("WorkerPool() got %d results, want %d", len(results), len(tt.expected))
				return
			}

			// Results may be in different order, so sort both
			sortInts(results)
			sortInts(tt.expected)

			for i, v := range results {
				if v != tt.expected[i] {
					t.Errorf("WorkerPool() result[%d] = %d, want %d", i, v, tt.expected[i])
				}
			}
		})
	}
}

// Helper function to generate doubled range
func makeDoubledRange(min, max int) []int {
	result := make([]int, max-min+1)
	for i := range result {
		result[i] = 2 * (min + i)
	}
	return result
}

func TestWorkerPoolWithContext(t *testing.T) {
	t.Run("normal completion", func(t *testing.T) {
		ctx := context.Background()
		jobs := []int{1, 2, 3}
		results, err := WorkerPoolWithContext(ctx, jobs, 2)

		if err != nil {
			t.Errorf("WorkerPoolWithContext() unexpected error: %v", err)
		}
		if len(results) != 3 {
			t.Errorf("WorkerPoolWithContext() got %d results, want 3", len(results))
		}
	})

	t.Run("context cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		jobs := []int{1, 2, 3, 4, 5}
		_, err := WorkerPoolWithContext(ctx, jobs, 2)

		if err == nil {
			t.Error("WorkerPoolWithContext() expected error on cancelled context")
		}
	})
}

func BenchmarkWorkerPool(b *testing.B) {
	jobs := makeRange(1, 100)

	benchmarks := []struct {
		name    string
		workers int
	}{
		{"1worker", 1},
		{"2workers", 2},
		{"4workers", 4},
		{"8workers", 8},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				WorkerPool(jobs, bm.workers)
			}
		})
	}
}

// Task 5: RateLimitedProcessor Tests
func TestRateLimitedProcessor(t *testing.T) {
	t.Run("rate limiting", func(t *testing.T) {
		items := []string{"a", "b", "c", "d", "e"}
		maxPerSecond := 10

		ch := RateLimitedProcessor(items, maxPerSecond)
		if ch == nil {
			t.Skip("RateLimitedProcessor not implemented yet")
			return
		}

		start := time.Now()
		results := []string{}

		// Add timeout protection
		done := make(chan bool)
		go func() {
			for item := range ch {
				results = append(results, item)
			}
			done <- true
		}()

		select {
		case <-done:
			// Success
		case <-time.After(2 * time.Second):
			t.Error("RateLimitedProcessor() timed out")
			return
		}

		duration := time.Since(start)

		if len(results) != len(items) {
			t.Errorf("RateLimitedProcessor() got %d results, want %d", len(results), len(items))
		}

		// Should take at least (len(items)-1)/maxPerSecond seconds
		minDuration := time.Duration((len(items)-1)*1000/maxPerSecond) * time.Millisecond
		if duration < minDuration {
			t.Errorf("RateLimitedProcessor() took %v, expected at least %v", duration, minDuration)
		}
	})

	t.Run("empty items", func(t *testing.T) {
		ch := RateLimitedProcessor([]string{}, 5)
		if ch == nil {
			t.Skip("RateLimitedProcessor not implemented yet")
			return
		}

		results := []string{}
		done := make(chan bool)
		go func() {
			for item := range ch {
				results = append(results, item)
			}
			done <- true
		}()

		select {
		case <-done:
			// Success
		case <-time.After(500 * time.Millisecond):
			t.Error("RateLimitedProcessor() with empty items timed out")
			return
		}

		if len(results) != 0 {
			t.Errorf("RateLimitedProcessor() with empty items got %d results, want 0", len(results))
		}
	})

	t.Run("nil items", func(t *testing.T) {
		var items []string
		ch := RateLimitedProcessor(items, 5)
		if ch == nil {
			t.Skip("RateLimitedProcessor not implemented yet")
			return
		}

		results := []string{}
		done := make(chan bool)
		go func() {
			for item := range ch {
				results = append(results, item)
			}
			done <- true
		}()

		select {
		case <-done:
			// Success
		case <-time.After(500 * time.Millisecond):
			t.Error("RateLimitedProcessor() with nil items timed out")
			return
		}

		if len(results) != 0 {
			t.Errorf("RateLimitedProcessor() with nil items got %d results, want 0", len(results))
		}
	})

	t.Run("zero rate", func(t *testing.T) {
		items := []string{"a", "b", "c"}
		ch := RateLimitedProcessor(items, 0)
		if ch == nil {
			t.Skip("RateLimitedProcessor not implemented yet")
			return
		}

		start := time.Now()
		results := []string{}
		done := make(chan bool)
		go func() {
			for item := range ch {
				results = append(results, item)
			}
			done <- true
		}()

		select {
		case <-done:
			// Success
		case <-time.After(1 * time.Second):
			t.Error("RateLimitedProcessor() with zero rate timed out")
			return
		}

		duration := time.Since(start)

		if len(results) != len(items) {
			t.Errorf("RateLimitedProcessor() with zero rate got %d results, want %d", len(results), len(items))
		}

		// With zero rate, should process immediately (less than 100ms)
		if duration > 100*time.Millisecond {
			t.Errorf("RateLimitedProcessor() with zero rate took %v, expected to be fast", duration)
		}
	})

	t.Run("negative rate", func(t *testing.T) {
		items := []string{"a", "b", "c"}
		ch := RateLimitedProcessor(items, -1)
		if ch == nil {
			t.Skip("RateLimitedProcessor not implemented yet")
			return
		}

		start := time.Now()
		results := []string{}
		done := make(chan bool)
		go func() {
			for item := range ch {
				results = append(results, item)
			}
			done <- true
		}()

		select {
		case <-done:
			// Success
		case <-time.After(1 * time.Second):
			t.Error("RateLimitedProcessor() with negative rate timed out")
			return
		}

		duration := time.Since(start)

		if len(results) != len(items) {
			t.Errorf("RateLimitedProcessor() with negative rate got %d results, want %d", len(results), len(items))
		}

		// With negative rate, should process immediately (less than 100ms)
		if duration > 100*time.Millisecond {
			t.Errorf("RateLimitedProcessor() with negative rate took %v, expected to be fast", duration)
		}
	})

	t.Run("single item", func(t *testing.T) {
		items := []string{"single"}
		ch := RateLimitedProcessor(items, 1)
		if ch == nil {
			t.Skip("RateLimitedProcessor not implemented yet")
			return
		}

		start := time.Now()
		results := []string{}
		done := make(chan bool)
		go func() {
			for item := range ch {
				results = append(results, item)
			}
			done <- true
		}()

		select {
		case <-done:
			// Success
		case <-time.After(1 * time.Second):
			t.Error("RateLimitedProcessor() with single item timed out")
			return
		}

		duration := time.Since(start)

		if len(results) != 1 || results[0] != "single" {
			t.Errorf("RateLimitedProcessor() with single item got %v, want [single]", results)
		}

		// Single item should process quickly (less than 100ms)
		if duration > 100*time.Millisecond {
			t.Errorf("RateLimitedProcessor() with single item took %v, expected to be fast", duration)
		}
	})

	t.Run("high rate", func(t *testing.T) {
		items := []string{"a", "b", "c", "d", "e"}
		maxPerSecond := 1000 // Very high rate

		ch := RateLimitedProcessor(items, maxPerSecond)
		if ch == nil {
			t.Skip("RateLimitedProcessor not implemented yet")
			return
		}

		start := time.Now()
		results := []string{}
		done := make(chan bool)
		go func() {
			for item := range ch {
				results = append(results, item)
			}
			done <- true
		}()

		select {
		case <-done:
			// Success
		case <-time.After(1 * time.Second):
			t.Error("RateLimitedProcessor() with high rate timed out")
			return
		}

		duration := time.Since(start)

		if len(results) != len(items) {
			t.Errorf("RateLimitedProcessor() with high rate got %d results, want %d", len(results), len(items))
		}

		// With high rate, should process quickly (less than 100ms)
		if duration > 100*time.Millisecond {
			t.Errorf("RateLimitedProcessor() with high rate took %v, expected to be fast", duration)
		}
	})
}

// Task 6: FanOutFanIn Tests
func TestFanOutFanIn(t *testing.T) {
	tests := []struct {
		name       string
		numbers    []int
		numWorkers int
		expected   []int
	}{
		{"empty", []int{}, 2, []int{}},
		{"nil slice", nil, 2, []int{}},
		{"single", []int{2}, 1, []int{6}},
		{"multiple", []int{1, 2, 3, 4}, 2, []int{3, 6, 9, 12}},
		{"zero workers", []int{1, 2, 3}, 0, []int{}},
		{"negative workers", []int{1, 2, 3}, -1, []int{}},
		{"large numbers", []int{100, 200, 300}, 2, []int{300, 600, 900}},
		{"negative numbers", []int{-1, -2, 3}, 2, []int{-3, -6, 9}},
		{"mixed numbers", []int{0, 1, -1, 5}, 3, []int{0, 3, -3, 15}},
		{"many workers few numbers", []int{1}, 100, []int{3}},
		{"many numbers few workers", makeRange(1, 50), 3, makeTripledRange(1, 50)},
		{"more workers than numbers", []int{1, 2, 3}, 10, []int{3, 6, 9}},
		{"single worker many numbers", makeRange(1, 20), 1, makeTripledRange(1, 20)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := FanOutFanIn(tt.numbers, tt.numWorkers)

			if len(results) != len(tt.expected) {
				t.Errorf("FanOutFanIn() got %d results, want %d", len(results), len(tt.expected))
				return
			}

			sortInts(results)
			sortInts(tt.expected)

			for i, v := range results {
				if v != tt.expected[i] {
					t.Errorf("FanOutFanIn() result[%d] = %d, want %d", i, v, tt.expected[i])
				}
			}
		})
	}
}

// Helper function to generate tripled range
func makeTripledRange(min, max int) []int {
	result := make([]int, max-min+1)
	for i := range result {
		result[i] = 3 * (min + i)
	}
	return result
}

func BenchmarkFanOutFanIn(b *testing.B) {
	numbers := makeRange(1, 100)

	for i := 0; i < b.N; i++ {
		FanOutFanIn(numbers, 4)
	}
}

// Task 7: ProcessWithTimeout Tests
func TestProcessWithTimeout(t *testing.T) {
	t.Run("completes within timeout", func(t *testing.T) {
		data := []int{1, 2, 3}
		count, err := ProcessWithTimeout(data, 1*time.Second)

		if err != nil {
			t.Errorf("ProcessWithTimeout() unexpected error: %v", err)
		}
		if count != len(data) {
			t.Errorf("ProcessWithTimeout() processed %d items, want %d", count, len(data))
		}
	})

	t.Run("times out", func(t *testing.T) {
		data := makeRange(1, 100)
		_, err := ProcessWithTimeout(data, 1*time.Millisecond)

		if err == nil {
			t.Error("ProcessWithTimeout() expected timeout error")
		}
	})

	t.Run("empty data", func(t *testing.T) {
		data := []int{}
		count, err := ProcessWithTimeout(data, 1*time.Second)

		if err != nil {
			t.Errorf("ProcessWithTimeout() with empty data unexpected error: %v", err)
		}
		if count != 0 {
			t.Errorf("ProcessWithTimeout() with empty data processed %d items, want 0", count)
		}
	})

	t.Run("nil data", func(t *testing.T) {
		var data []int
		count, err := ProcessWithTimeout(data, 1*time.Second)

		if err != nil {
			t.Errorf("ProcessWithTimeout() with nil data unexpected error: %v", err)
		}
		if count != 0 {
			t.Errorf("ProcessWithTimeout() with nil data processed %d items, want 0", count)
		}
	})

	t.Run("very short timeout", func(t *testing.T) {
		data := makeRange(1, 1000)
		_, err := ProcessWithTimeout(data, 1*time.Nanosecond)

		if err == nil {
			t.Error("ProcessWithTimeout() with very short timeout expected error")
		}
	})

	t.Run("very long timeout", func(t *testing.T) {
		data := makeRange(1, 10)
		count, err := ProcessWithTimeout(data, 10*time.Second)

		if err != nil {
			t.Errorf("ProcessWithTimeout() with long timeout unexpected error: %v", err)
		}
		if count != len(data) {
			t.Errorf("ProcessWithTimeout() with long timeout processed %d items, want %d", count, len(data))
		}
	})

	t.Run("large data small timeout", func(t *testing.T) {
		data := makeRange(1, 10000)
		_, err := ProcessWithTimeout(data, 1*time.Microsecond)

		if err == nil {
			t.Error("ProcessWithTimeout() with large data and small timeout expected error")
		}
	})

	t.Run("single item", func(t *testing.T) {
		data := []int{42}
		count, err := ProcessWithTimeout(data, 1*time.Second)

		if err != nil {
			t.Errorf("ProcessWithTimeout() with single item unexpected error: %v", err)
		}
		if count != 1 {
			t.Errorf("ProcessWithTimeout() with single item processed %d items, want 1", count)
		}
	})
}

// Task 8: ConcurrentDownloader Tests
func TestConcurrentDownloader(t *testing.T) {
	tests := []struct {
		name          string
		urls          []string
		maxConcurrent int
		minSuccess    int
	}{
		{"empty", []string{}, 2, 0},
		{"nil urls", nil, 2, 0},
		{"single", []string{"http://example.com"}, 1, 1},
		{"multiple", []string{"http://example.com", "http://google.com", "http://github.com"}, 2, 3},
		{"zero concurrent", []string{"http://example.com"}, 0, 0},
		{"negative concurrent", []string{"http://example.com"}, -1, 0},
		{"single url many concurrent", []string{"http://example.com"}, 100, 1},
		{"many urls few concurrent", []string{
			"http://example.com", "http://google.com", "http://github.com",
			"http://stackoverflow.com", "http://reddit.com",
		}, 2, 5},
		{"duplicate urls", []string{"http://example.com", "http://example.com"}, 2, 2},
		{"invalid urls", []string{"http://invalid.url.that.does.not.exist"}, 1, 0},
		{"mixed valid/invalid", []string{
			"http://example.com", "http://invalid.url.that.does.not.exist",
			"http://google.com",
		}, 2, 2},
		{"large concurrent limit", []string{
			"http://example.com", "http://google.com", "http://github.com",
		}, 100, 3},
		{"more urls than concurrent limit", makeExampleURLs(50), 5, 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			success := ConcurrentDownloader(tt.urls, tt.maxConcurrent)
			if success < tt.minSuccess {
				t.Errorf("ConcurrentDownloader() got %d successes, want at least %d", success, tt.minSuccess)
			}
		})
	}
}

// Helper function to generate example URLs
func makeExampleURLs(count int) []string {
	urls := make([]string, count)
	for i := 0; i < count; i++ {
		urls[i] = fmt.Sprintf("http://example%d.com", i)
	}
	return urls
}

func BenchmarkConcurrentDownloader(b *testing.B) {
	urls := []string{
		"http://example.com",
		"http://google.com",
		"http://github.com",
		"http://stackoverflow.com",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcurrentDownloader(urls, 2)
	}
}

// Helper functions
func makeRange(min, max int) []int {
	result := make([]int, max-min+1)
	for i := range result {
		result[i] = min + i
	}
	return result
}

func sortInts(arr []int) {
	// Simple bubble sort for small arrays in tests
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
