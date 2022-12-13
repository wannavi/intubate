package intubate

import "sync"

type SafeCache struct {
	v   map[string]bool
	mux sync.Mutex
}

func (vh *SafeCache) Set(key string, newValue bool) {
	vh.mux.Lock()
	vh.v[key] = newValue
	vh.mux.Unlock()
}

func (vh *SafeCache) Get(key string) bool {
	vh.mux.Lock()
	defer vh.mux.Unlock()
	return vh.v[key]
}
