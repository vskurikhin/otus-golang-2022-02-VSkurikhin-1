package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		// Write me
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}

// На логику выталкивания элементов из-за размера очереди.
func TestCacheEjection(t *testing.T) {
	cache := NewCache(3)
	setTests := []struct {
		key      string
		value    string
		expected bool
	}{
		{key: "key0", value: "test0", expected: false},
		{key: "key1", value: "test1", expected: false},
		{key: "key2", value: "test2", expected: false},
		{key: "key3", value: "test3", expected: false},
	}
	for _, tc := range setTests {
		tc := tc
		t.Run(tc.key, func(t *testing.T) {
			ok := cache.Set(Key(tc.key), tc.value)
			require.Equal(t, tc.expected, ok)
		})
	}
	getTests := []struct {
		key      string
		value    string
		expected string
		ok       bool
	}{
		{key: "key0", expected: "test0", ok: false},
		{key: "key1", expected: "test1", ok: true},
		{key: "key2", expected: "test2", ok: true},
		{key: "key3", expected: "test3", ok: true},
	}
	for _, tc := range getTests {
		tc := tc
		t.Run(tc.key, func(t *testing.T) {
			value, ok := cache.Get(Key(tc.key))
			require.Equal(t, tc.ok, ok)
			if tc.ok {
				require.Equal(t, tc.expected, value)
			} else {
				require.NotEqual(t, tc.expected, value)
			}
		})
	}
}

// На логику выталкивания давно используемых элементов.
func TestCacheEjectionLeastRecentlyUnused(t *testing.T) {
	cache := NewCache(3)
	setTests := []struct {
		key      string
		value    string
		expected bool
	}{
		{key: "key0", value: "test0", expected: false},
		{key: "key1", value: "test1", expected: false},
		{key: "key2", value: "test2", expected: false},
	}
	for _, tc := range setTests {
		tc := tc
		t.Run(tc.key, func(t *testing.T) {
			ok := cache.Set(Key(tc.key), tc.value)
			require.Equal(t, tc.expected, ok)
		})
	}

	getTests := []struct {
		key      string
		value    string
		expected string
		ok       bool
	}{
		{key: "key0", expected: "test0", ok: true},
		{key: "key2", expected: "test2", ok: true},
	}
	for _, tc := range getTests {
		tc := tc
		t.Run(tc.key, func(t *testing.T) {
			value, ok := cache.Get(Key(tc.key))
			require.Equal(t, tc.ok, ok)
			require.Equal(t, tc.expected, value)
		})
	}
	ok3 := cache.Set("key3", "test3")
	require.False(t, ok3)

	test1, ok1 := cache.Get("key1")
	require.False(t, ok1)
	require.NotEqual(t, "test1", test1)
	require.Nil(t, test1)
}

func TestCacheEmpty(t *testing.T) {
	cache := NewCache(1)
	test, ok := cache.Get("test")
	require.Nil(t, test)
	require.False(t, ok)
}

func TestCacheSet(t *testing.T) {
	cache := NewCache(1)
	ok0 := cache.Set("key", "test")
	require.False(t, ok0)
	ok1 := cache.Set("key", "test")
	require.True(t, ok1)
	test, ok2 := cache.Get("key")
	require.True(t, ok2)
	require.Equal(t, "test", test)
}
