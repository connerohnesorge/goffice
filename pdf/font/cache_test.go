package font

import (
	"errors"
	"sync"
	"testing"
)

func TestNewFontCache(t *testing.T) {
	tests := []struct {
		name            string
		maxSize         int
		expectedMaxSize int
	}{
		{"positive size", 50, 50},
		{
			"zero size uses default",
			0,
			DefaultCacheSize,
		},
		{
			"negative size uses default",
			-1,
			DefaultCacheSize,
		},
		{"large size", 1000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := NewFontCache(tt.maxSize)
			if cache == nil {
				t.Fatal(
					"NewFontCache returned nil",
				)
			}
			if cache.MaxSize() != tt.expectedMaxSize {
				t.Errorf(
					"MaxSize() = %d, want %d",
					cache.MaxSize(),
					tt.expectedMaxSize,
				)
			}
			if cache.Size() != 0 {
				t.Errorf(
					"Size() = %d, want 0 for new cache",
					cache.Size(),
				)
			}
		})
	}
}

func TestFontCachePutAndGet(t *testing.T) {
	cache := NewFontCache(10)

	// Create test fonts
	font1 := &Font{
		Family: "Arial",
		Style:  StyleRegular,
	}
	font2 := &Font{
		Family: "Arial",
		Style:  StyleBold,
	}
	font3 := &Font{
		Family: "Times",
		Style:  StyleRegular,
	}

	// Put fonts in cache
	cache.Put(font1)
	cache.Put(font2)
	cache.Put(font3)

	if cache.Size() != 3 {
		t.Errorf(
			"Size() = %d, want 3",
			cache.Size(),
		)
	}

	// Get fonts back
	got1, ok := cache.Get("Arial", StyleRegular)
	if !ok || got1 != font1 {
		t.Error("Get Arial Regular failed")
	}

	got2, ok := cache.Get("Arial", StyleBold)
	if !ok || got2 != font2 {
		t.Error("Get Arial Bold failed")
	}

	got3, ok := cache.Get("Times", StyleRegular)
	if !ok || got3 != font3 {
		t.Error("Get Times Regular failed")
	}

	// Try to get non-existent font
	_, ok = cache.Get("Helvetica", StyleRegular)
	if ok {
		t.Error(
			"Get should return false for non-existent font",
		)
	}
}

func TestFontCachePutNil(t *testing.T) {
	cache := NewFontCache(10)
	cache.Put(
		nil,
	) // Should not panic or add anything
	if cache.Size() != 0 {
		t.Errorf(
			"Size() = %d, want 0 after putting nil",
			cache.Size(),
		)
	}
}

func TestFontCacheLRUEviction(t *testing.T) {
	cache := NewFontCache(
		3,
	) // Small cache for testing

	// Add 3 fonts - cache should be at capacity
	font1 := &Font{
		Family: "Font1",
		Style:  StyleRegular,
	}
	font2 := &Font{
		Family: "Font2",
		Style:  StyleRegular,
	}
	font3 := &Font{
		Family: "Font3",
		Style:  StyleRegular,
	}

	cache.Put(font1)
	cache.Put(font2)
	cache.Put(font3)

	if cache.Size() != 3 {
		t.Fatalf(
			"Size() = %d, want 3",
			cache.Size(),
		)
	}

	// Add a 4th font - font1 (LRU) should be evicted
	font4 := &Font{
		Family: "Font4",
		Style:  StyleRegular,
	}
	cache.Put(font4)

	if cache.Size() != 3 {
		t.Errorf(
			"Size() = %d, want 3 after eviction",
			cache.Size(),
		)
	}

	// font1 should be evicted (it was least recently used)
	_, ok := cache.Get("Font1", StyleRegular)
	if ok {
		t.Error("Font1 should have been evicted")
	}

	// Other fonts should still be present
	_, ok = cache.Get("Font2", StyleRegular)
	if !ok {
		t.Error("Font2 should still be in cache")
	}
	_, ok = cache.Get("Font3", StyleRegular)
	if !ok {
		t.Error("Font3 should still be in cache")
	}
	_, ok = cache.Get("Font4", StyleRegular)
	if !ok {
		t.Error("Font4 should be in cache")
	}
}

func TestFontCacheLRUOrdering(t *testing.T) {
	cache := NewFontCache(3)

	font1 := &Font{
		Family: "Font1",
		Style:  StyleRegular,
	}
	font2 := &Font{
		Family: "Font2",
		Style:  StyleRegular,
	}
	font3 := &Font{
		Family: "Font3",
		Style:  StyleRegular,
	}

	cache.Put(font1)
	cache.Put(font2)
	cache.Put(font3)

	// Access font1 to make it most recently used
	cache.Get("Font1", StyleRegular)

	// Add a new font - font2 should be evicted (now LRU)
	font4 := &Font{
		Family: "Font4",
		Style:  StyleRegular,
	}
	cache.Put(font4)

	_, ok := cache.Get("Font2", StyleRegular)
	if ok {
		t.Error(
			"Font2 should have been evicted as LRU",
		)
	}

	_, ok = cache.Get("Font1", StyleRegular)
	if !ok {
		t.Error(
			"Font1 should still be in cache (was accessed)",
		)
	}
}

func TestFontCacheUpdateExisting(t *testing.T) {
	cache := NewFontCache(10)

	font1 := &Font{
		Family:  "Arial",
		Style:   StyleRegular,
		Metrics: FontMetrics{UnitsPerEm: 1000},
	}
	cache.Put(font1)

	// Update with same key but different data
	font2 := &Font{
		Family:  "Arial",
		Style:   StyleRegular,
		Metrics: FontMetrics{UnitsPerEm: 2048},
	}
	cache.Put(font2)

	// Size should still be 1
	if cache.Size() != 1 {
		t.Errorf(
			"Size() = %d, want 1 after update",
			cache.Size(),
		)
	}

	// Should get the updated font
	got, ok := cache.Get("Arial", StyleRegular)
	if !ok {
		t.Fatal("Get failed after update")
	}
	if got.Metrics.UnitsPerEm != 2048 {
		t.Error(
			"Got old font instead of updated one",
		)
	}
}

func TestFontCacheGetOrLoad(t *testing.T) {
	cache := NewFontCache(10)
	loadCount := 0

	loader := func() (*Font, error) {
		loadCount++

		return &Font{
			Family: "TestFont",
			Style:  StyleRegular,
		}, nil
	}

	// First call should load
	font1, err := cache.GetOrLoad(
		"TestFont",
		StyleRegular,
		loader,
	)
	if err != nil {
		t.Fatalf("GetOrLoad failed: %v", err)
	}
	if loadCount != 1 {
		t.Errorf(
			"loadCount = %d, want 1",
			loadCount,
		)
	}
	if font1.Family != "TestFont" {
		t.Errorf(
			"Family = %s, want TestFont",
			font1.Family,
		)
	}

	// Second call should use cache
	font2, err := cache.GetOrLoad(
		"TestFont",
		StyleRegular,
		loader,
	)
	if err != nil {
		t.Fatalf("GetOrLoad failed: %v", err)
	}
	if loadCount != 1 {
		t.Errorf(
			"loadCount = %d, want 1 (should use cache)",
			loadCount,
		)
	}
	if font2 != font1 {
		t.Error(
			"Should return same font from cache",
		)
	}
}

func TestFontCacheGetOrLoadError(t *testing.T) {
	cache := NewFontCache(10)
	expectedErr := errors.New("load error")

	loader := func() (*Font, error) {
		return nil, expectedErr
	}

	_, err := cache.GetOrLoad(
		"TestFont",
		StyleRegular,
		loader,
	)
	if err == nil {
		t.Fatal("GetOrLoad should return error")
	}
	if err != expectedErr {
		t.Errorf(
			"err = %v, want %v",
			err,
			expectedErr,
		)
	}

	// Should not cache failed loads
	if cache.Size() != 0 {
		t.Errorf(
			"Size() = %d, want 0 after failed load",
			cache.Size(),
		)
	}
}

func TestFontCacheClear(t *testing.T) {
	cache := NewFontCache(10)

	cache.Put(
		&Font{
			Family: "Font1",
			Style:  StyleRegular,
		},
	)
	cache.Put(
		&Font{
			Family: "Font2",
			Style:  StyleRegular,
		},
	)
	cache.Put(
		&Font{
			Family: "Font3",
			Style:  StyleRegular,
		},
	)

	if cache.Size() != 3 {
		t.Fatalf(
			"Size() = %d, want 3",
			cache.Size(),
		)
	}

	cache.Clear()

	if cache.Size() != 0 {
		t.Errorf(
			"Size() = %d, want 0 after Clear",
			cache.Size(),
		)
	}

	// Verify fonts are gone
	_, ok := cache.Get("Font1", StyleRegular)
	if ok {
		t.Error(
			"Font1 should be gone after Clear",
		)
	}
}

func TestFontCacheSetMaxSize(t *testing.T) {
	cache := NewFontCache(10)

	// Add 5 fonts
	for i := range 5 {
		cache.Put(
			&Font{
				Family: string(rune('A' + i)),
				Style:  StyleRegular,
			},
		)
	}

	if cache.Size() != 5 {
		t.Fatalf(
			"Size() = %d, want 5",
			cache.Size(),
		)
	}

	// Reduce max size - should evict LRU entries
	cache.SetMaxSize(3)

	if cache.MaxSize() != 3 {
		t.Errorf(
			"MaxSize() = %d, want 3",
			cache.MaxSize(),
		)
	}
	if cache.Size() != 3 {
		t.Errorf(
			"Size() = %d, want 3 after reducing max size",
			cache.Size(),
		)
	}
}

func TestFontCacheSetMaxSizeZero(t *testing.T) {
	cache := NewFontCache(10)
	cache.SetMaxSize(0)
	if cache.MaxSize() != DefaultCacheSize {
		t.Errorf(
			"MaxSize() = %d, want %d",
			cache.MaxSize(),
			DefaultCacheSize,
		)
	}
}

func TestFontCacheFallbackChain(t *testing.T) {
	cache := NewFontCache(10)

	// Initially nil
	if cache.GetFallbackChain() != nil {
		t.Error(
			"GetFallbackChain should return nil initially",
		)
	}

	// Set a fallback chain
	chain := &FallbackChain{}
	cache.SetFallbackChain(chain)

	if cache.GetFallbackChain() != chain {
		t.Error(
			"GetFallbackChain should return the set chain",
		)
	}

	// Set to nil
	cache.SetFallbackChain(nil)
	if cache.GetFallbackChain() != nil {
		t.Error(
			"GetFallbackChain should return nil after setting to nil",
		)
	}
}

func TestFontCacheKeys(t *testing.T) {
	cache := NewFontCache(10)

	cache.Put(
		&Font{
			Family: "Arial",
			Style:  StyleRegular,
		},
	)
	cache.Put(
		&Font{Family: "Arial", Style: StyleBold},
	)
	cache.Put(
		&Font{
			Family: "Times",
			Style:  StyleItalic,
		},
	)

	keys := cache.Keys()
	if len(keys) != 3 {
		t.Errorf(
			"len(Keys()) = %d, want 3",
			len(keys),
		)
	}

	// Check that expected keys are present
	keySet := make(map[string]bool)
	for _, k := range keys {
		keySet[k] = true
	}

	expectedKeys := []string{
		"Arial:0",
		"Arial:1",
		"Times:2",
	}
	for _, ek := range expectedKeys {
		if !keySet[ek] {
			t.Errorf(
				"Expected key %s not found",
				ek,
			)
		}
	}
}

func TestFontCacheConcurrency(t *testing.T) {
	cache := NewFontCache(100)
	var wg sync.WaitGroup
	numGoroutines := 10
	numOperations := 100

	// Concurrent writes
	for i := range numGoroutines {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := range numOperations {
				font := &Font{
					Family: string(
						rune(
							'A' + (id*numOperations+j)%26,
						),
					),
					Style: FontStyle(j % 4),
				}
				cache.Put(font)
			}
		}(i)
	}

	// Concurrent reads
	for i := range numGoroutines {
		wg.Add(1)
		go func(_ int) {
			defer wg.Done()
			for j := range numOperations {
				cache.Get(
					string(rune('A'+j%26)),
					FontStyle(j%4),
				)
			}
		}(i)
	}

	// Concurrent GetOrLoad
	for i := range numGoroutines {
		wg.Add(1)
		go func(_ int) {
			defer wg.Done()
			for j := range numOperations {
				family := string(rune('A' + j%26))
				style := FontStyle(j % 4)
				_, _ = cache.GetOrLoad(
					family,
					style,
					func() (*Font, error) {
						return &Font{
							Family: family,
							Style:  style,
						}, nil
					},
				)
			}
		}(i)
	}

	wg.Wait()

	// Cache should still be functional
	if cache.Size() > cache.MaxSize() {
		t.Errorf(
			"Cache size %d exceeds max size %d",
			cache.Size(),
			cache.MaxSize(),
		)
	}
}

func TestCacheKey(t *testing.T) {
	tests := []struct {
		family   string
		style    FontStyle
		expected string
	}{
		{"Arial", StyleRegular, "Arial:0"},
		{"Arial", StyleBold, "Arial:1"},
		{
			"Times New Roman",
			StyleItalic,
			"Times New Roman:2",
		},
		{"Courier", StyleBoldItalic, "Courier:3"},
		{"", StyleRegular, ":0"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			key := cacheKey(tt.family, tt.style)
			if key != tt.expected {
				t.Errorf(
					"cacheKey(%s, %v) = %s, want %s",
					tt.family,
					tt.style,
					key,
					tt.expected,
				)
			}
		})
	}
}

func TestPathKey(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{
			"/usr/share/fonts/arial.ttf",
			"path:/usr/share/fonts/arial.ttf",
		},
		{
			"C:\\Windows\\Fonts\\arial.ttf",
			"path:C:\\Windows\\Fonts\\arial.ttf",
		},
		{
			"./fonts/test.otf",
			"path:./fonts/test.otf",
		},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			key := pathKey(tt.path)
			if key != tt.expected {
				t.Errorf(
					"pathKey(%s) = %s, want %s",
					tt.path,
					key,
					tt.expected,
				)
			}
		})
	}
}

// BenchmarkFontCacheGet benchmarks cache get operations
func BenchmarkFontCacheGet(b *testing.B) {
	cache := NewFontCache(100)

	// Pre-populate cache
	for i := range 50 {
		cache.Put(&Font{
			Family: string(rune('A' + i%26)),
			Style:  FontStyle(i % 4),
		})
	}

	b.ResetTimer()
	for i := range b.N {
		cache.Get(
			string(rune('A'+i%26)),
			FontStyle(i%4),
		)
	}
}

// BenchmarkFontCachePut benchmarks cache put operations
func BenchmarkFontCachePut(b *testing.B) {
	cache := NewFontCache(100)

	fonts := make([]*Font, 100)
	for i := range 100 {
		fonts[i] = &Font{
			Family: string(rune('A' + i%26)),
			Style:  FontStyle(i % 4),
		}
	}

	b.ResetTimer()
	for i := range b.N {
		cache.Put(fonts[i%100])
	}
}

// BenchmarkFontCacheGetOrLoad benchmarks GetOrLoad operations
func BenchmarkFontCacheGetOrLoad(b *testing.B) {
	cache := NewFontCache(100)

	b.ResetTimer()
	for i := range b.N {
		family := string(rune('A' + i%26))
		style := FontStyle(i % 4)
		_, _ = cache.GetOrLoad(
			family,
			style,
			func() (*Font, error) {
				return &Font{
					Family: family,
					Style:  style,
				}, nil
			},
		)
	}
}

// BenchmarkFontCacheConcurrentAccess benchmarks concurrent access
func BenchmarkFontCacheConcurrentAccess(
	b *testing.B,
) {
	cache := NewFontCache(100)

	// Pre-populate
	for i := range 50 {
		cache.Put(&Font{
			Family: string(rune('A' + i%26)),
			Style:  FontStyle(i % 4),
		})
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%2 == 0 {
				cache.Get(
					string(rune('A'+i%26)),
					FontStyle(i%4),
				)
			} else {
				cache.Put(&Font{
					Family: string(rune('A' + i%26)),
					Style:  FontStyle(i % 4),
				})
			}
			i++
		}
	})
}
