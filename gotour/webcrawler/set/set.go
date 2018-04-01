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

func (self *SafeSet) Contains(str string) bool {
	self.mux.Lock()
	defer self.mux.Unlock()

	_, ok := self.set[str]
	return ok
}

func (self *SafeSet) Add(str string) {
	self.mux.Lock()
	defer self.mux.Unlock()

	self.set[str] = EmptyStruct{}
}

// Similar to `sync.Map.LoadOrStore`, but without actual load
// https://godoc.org/sync#Map.LoadOrStore
func (self *SafeSet) IsPresentOrAdd(str string) bool {
	self.mux.Lock()
	defer self.mux.Unlock()

	_, present := self.set[str]
	if !present {
		self.set[str] = EmptyStruct{}
	}

	return present
}

func (self SafeSet) String() string {
	self.mux.Lock()
	defer self.mux.Unlock()

	var ret string

	// string.Builder since go-1.10
	ret += "[ "
	for k := range self.set {
		ret += k
		ret += " "
	}
	ret += "]"

	return ret
}
