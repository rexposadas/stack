package main

import (
	"fmt"
	"testing"
)

func BenchmarkCacheGet(b *testing.B) {
	cache, _ := NewCache(1000, &Connection{})

	// Pre-populate cache
	for i := 0; i < 500; i++ {
		cache.Put(fmt.Sprintf("key-%d", i), i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", i%1000)
		_, _ = cache.Get(key)
	}
}

func BenchmarkCachePut(b *testing.B) {
	cache, _ := NewCache(1000, &Connection{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", i)
		cache.Put(key, i%100)
	}
}

func BenchmarkCacheEviction(b *testing.B) {
	cache, _ := NewCache(100, &Connection{})

	// Fill cache to capacity
	for i := 0; i < 100; i++ {
		cache.Put(fmt.Sprintf("key-%d", i), i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// This will trigger eviction
		cache.Put(fmt.Sprintf("new-key-%d", i), i%100)
	}
}

func BenchmarkCacheConcurrentReads(b *testing.B) {
	cache, _ := NewCache(1000, &Connection{})

	// Pre-populate cache
	for i := 0; i < 1000; i++ {
		cache.Put(fmt.Sprintf("key-%d", i), i)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := fmt.Sprintf("key-%d", i%1000)
			_, _ = cache.Get(key)
			i++
		}
	})
}

func BenchmarkCacheConcurrentWrites(b *testing.B) {
	cache, _ := NewCache(1000, &Connection{})

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := fmt.Sprintf("key-%d", i)
			cache.Put(key, i%100)
			i++
		}
	})
}

func BenchmarkCacheConcurrentMixed(b *testing.B) {
	cache, _ := NewCache(1000, &Connection{})

	// Pre-populate cache
	for i := 0; i < 500; i++ {
		cache.Put(fmt.Sprintf("key-%d", i), i)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%2 == 0 {
				key := fmt.Sprintf("key-%d", i%1000)
				_, _ = cache.Get(key)
			} else {
				key := fmt.Sprintf("key-%d", i)
				cache.Put(key, i%100)
			}
			i++
		}
	})
}
