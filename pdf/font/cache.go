// cache.go provides a thread-safe LRU font cache to avoid re-parsing fonts repeatedly during rendering.

package font

import (
	"container/list"
	"fmt"
	"sync"
)

// DefaultCacheSize is the default maximum number of fonts to cache.
const DefaultCacheSize = 100

// FontCache is a thread-safe LRU cache for parsed fonts.
// It prevents repeated parsing of the same font files during rendering.
type FontCache struct {
	mu       sync.RWMutex
	fonts    map[string]*cacheEntry
	lruList  *list.List // For LRU ordering (front = most recently used)
	maxSize  int        // Max number of fonts to cache
	fallback *FallbackChain
}

// cacheEntry represents a cached font with its LRU position.
type cacheEntry struct {
	font    *Font
	key     string
	lruElem *list.Element
}

// NewFontCache creates a new font cache with the specified maximum size.
// If maxSize is 0 or negative, DefaultCacheSize (100) is used.
func NewFontCache(maxSize int) *FontCache {
	effectiveSize := maxSize
	if effectiveSize <= 0 {
		effectiveSize = DefaultCacheSize
	}

	return &FontCache{
		fonts:   make(map[string]*cacheEntry),
		lruList: list.New(),
		maxSize: effectiveSize,
	}
}

// cacheKey generates a unique cache key for a font based on family and style.
func cacheKey(
	family string,
	style FontStyle,
) string {
	return fmt.Sprintf("%s:%d", family, style)
}

// pathKey generates a cache key for a font loaded from a file path.
func pathKey(path string) string {
	return fmt.Sprintf("path:%s", path)
}

// Get retrieves a font from the cache by family name and style.
// Returns (font, true) if found, (nil, false) if not found.
// Updates LRU ordering on access.
func (c *FontCache) Get(
	family string,
	style FontStyle,
) (*Font, bool) {
	key := cacheKey(family, style)

	return c.getByKey(key)
}

// getByKey retrieves a font from the cache by its cache key.
func (c *FontCache) getByKey(
	key string,
) (*Font, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.fonts[key]
	if !ok {
		return nil, false
	}

	// Move to front of LRU list (most recently used)
	c.lruList.MoveToFront(entry.lruElem)

	return entry.font, true
}

// Put adds a font to the cache.
// The cache key is derived from the font's Family and Style.
// If the cache is at capacity, the least recently used font is evicted.
func (c *FontCache) Put(font *Font) {
	if font == nil {
		return
	}
	key := cacheKey(font.Family, font.Style)
	c.putWithKey(key, font)
}

// putWithKey adds a font to the cache with a specific key.
func (c *FontCache) putWithKey(
	key string,
	font *Font,
) {
	if font == nil {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	// Check if font already exists
	if entry, ok := c.fonts[key]; ok {
		// Update the font and move to front
		entry.font = font
		c.lruList.MoveToFront(entry.lruElem)

		return
	}

	// Evict least recently used if at capacity
	for c.lruList.Len() >= c.maxSize {
		c.evictLRU()
	}

	// Add new entry
	entry := &cacheEntry{
		font: font,
		key:  key,
	}
	entry.lruElem = c.lruList.PushFront(entry)
	c.fonts[key] = entry
}

// evictLRU removes the least recently used entry from the cache.
// Must be called with the lock held.
func (c *FontCache) evictLRU() {
	back := c.lruList.Back()
	if back == nil {
		return
	}

	entry, ok := back.Value.(*cacheEntry)
	if !ok {
		return
	}
	c.lruList.Remove(back)
	delete(c.fonts, entry.key)
}

// GetOrLoad returns a cached font if available, otherwise calls the loader function,
// caches the result, and returns it.
// This is the recommended way to ensure fonts are cached properly.
func (c *FontCache) GetOrLoad(
	family string,
	style FontStyle,
	loader func() (*Font, error),
) (*Font, error) {
	// Try to get from cache first (read lock)
	if font, ok := c.Get(family, style); ok {
		return font, nil
	}

	// Not in cache, load the font
	font, err := loader()
	if err != nil {
		return nil, err
	}

	// Ensure the font has the correct family and style for caching
	if font.Family == "" {
		font.Family = family
	}
	if font.Style != style &&
		font.Family == family {
		// Override style if it was explicitly requested
		font.Style = style
	}

	// Cache the font
	c.Put(font)

	return font, nil
}

// LoadFile loads a font from a file path, caching by the file path.
// Uses ParseFontFile internally for parsing.
// If the font is already cached by path, returns the cached version.
func (c *FontCache) LoadFile(
	path string,
) (*Font, error) {
	key := pathKey(path)

	// Try to get from cache first
	if font, ok := c.getByKey(key); ok {
		return font, nil
	}

	// Load the font from file
	font, err := ParseFontFile(path)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to load font from %s: %w",
			path,
			err,
		)
	}

	// Cache by path
	c.putWithKey(key, font)

	// Also cache by family+style for lookup by name
	familyKey := cacheKey(font.Family, font.Style)
	c.putWithKey(familyKey, font)

	return font, nil
}

// Clear removes all cached fonts.
func (c *FontCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.fonts = make(map[string]*cacheEntry)
	c.lruList.Init()
}

// Size returns the number of fonts currently cached.
func (c *FontCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.fonts)
}

// SetFallbackChain sets the fallback chain for font substitution.
// This is used by task 1.9 for automatic font fallback.
func (c *FontCache) SetFallbackChain(
	chain *FallbackChain,
) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.fallback = chain
}

// GetFallbackChain returns the current fallback chain.
func (c *FontCache) GetFallbackChain() *FallbackChain {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.fallback
}

// MaxSize returns the maximum number of fonts the cache can hold.
func (c *FontCache) MaxSize() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.maxSize
}

// SetMaxSize changes the maximum cache size.
// If the new size is smaller than the current number of cached fonts,
// the least recently used fonts are evicted until the cache is within limits.
func (c *FontCache) SetMaxSize(maxSize int) {
	effectiveSize := maxSize
	if effectiveSize <= 0 {
		effectiveSize = DefaultCacheSize
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.maxSize = effectiveSize

	// Evict until within limits
	for c.lruList.Len() > c.maxSize {
		c.evictLRU()
	}
}

// Keys returns a slice of all cache keys (for debugging/testing).
func (c *FontCache) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := make([]string, 0, len(c.fonts))
	for key := range c.fonts {
		keys = append(keys, key)
	}

	return keys
}
