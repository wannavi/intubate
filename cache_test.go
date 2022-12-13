package intubate

import "testing"

func TestSafeCache(t *testing.T) {
	url := "https://github.com"
	cache := SafeCache{v: make(map[string]bool)}

	t.Run("set cache to true", func(t *testing.T) {
		want := true
		cache.Set(url, want)

		got := cache.Get(url)

		if got != want {
			t.Fatalf("want %t, got %t", want, got)
		}
	})
}
