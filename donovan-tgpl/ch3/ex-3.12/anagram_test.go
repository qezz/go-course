package main

import "testing"

var anagramTests []struct {
	left  [2]string
	right bool
}

func init() {
	anagramTests = []struct {
		left  [2]string
		right bool
	}{
		{[2]string{"", ""}, true},
		{[2]string{"   ", ""}, true},
		{[2]string{"", "   "}, true},
		{[2]string{"       ", "   "}, true},
		{[2]string{"restful", "fluster"}, true},
		{[2]string{"funeral", "real fun"}, true},
		{[2]string{"adultery", "true lady"}, true},
		{[2]string{"customers", "store scum"}, true},
		{[2]string{"forty five", "over fifty"}, true},

		{[2]string{"a", ""}, false},
		{[2]string{"restfull", "fluster"}, false},
		{[2]string{"fewneral", "real fun"}, false},
		{[2]string{"adulltery", "false lady"}, false},
		{[2]string{"costumes", "store scum"}, false},
		{[2]string{"four tea five", "over fifty"}, false},
	}
}

func TestAnagrams(t *testing.T) {
	for _, row := range anagramTests {
		if AreAnagrams(row.left[0], row.left[1]) != row.right {
			t.Errorf("left:  %v %v", row.left[0], row.left[1])
			t.Errorf("right: %v\n", row.right)
			t.Fail()
		}
	}
}
