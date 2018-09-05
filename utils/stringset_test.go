package utils

import (
	"testing"
)

func TestStringSet(t *testing.T) {

	var set *StringSet
	set = NewStringSet()
	set.Add("shangsan")
	set.Add("lisi")
	set.Add("shangsan")

	x := len(set.ToArray())
	if x == 2 {
		t.Logf("success, %v", x)
	} else {
		t.Errorf("error")
	}
}
