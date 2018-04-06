package set

import (
	"sync"
)

type EmptyStruct struct{}

// `set` structure simulation
type SafeSet struct {
	set map[string]EmptyStruct
	mux sync.Mutex
}

func NewSafeSet() SafeSet {
	return SafeSet{set: make(map[string]EmptyStruct)}
}

func (set *SafeSet) Contains(str string) bool {
	set.mux.Lock()
	defer set.mux.Unlock()

	_, ok := set.set[str]
	return ok
}

func (set *SafeSet) Add(str string) {
	set.mux.Lock()
	defer set.mux.Unlock()

	set.set[str] = EmptyStruct{}
}

// Similar to `sync.Map.LoadOrStore`, but without actual load
// https://godoc.org/sync#Map.LoadOrStore
func (set *SafeSet) IsPresentOrAdd(str string) bool {
	set.mux.Lock()
	defer set.mux.Unlock()

	_, present := set.set[str]
	if !present {
		set.set[str] = EmptyStruct{}
	}

	return present
}

func (set SafeSet) String() string {
	set.mux.Lock()
	defer set.mux.Unlock()

	var ret string

	// string.Builder since go-1.10
	ret += "[ "
	for k := range set.set {
		ret += k
		ret += " "
	}
	ret += "]"

	return ret
}
