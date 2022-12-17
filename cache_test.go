package intubate

import (
	"strconv"
	"testing"
)

func TestSafeCache(t *testing.T) {
	cache := NewCache()

	for i := 0; i < 10; i++ {
		cache.Set(strconv.Itoa(i), i)
	}

	want := 4
	got := cache.Get("4")

	if got.(int) != want {
		t.Fatalf("want %d, got %d", want, got)
	}
}
