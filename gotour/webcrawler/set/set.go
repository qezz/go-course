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
	self.set[str] = EmptyStruct{}
}

func (self SafeSet) String() string {
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
