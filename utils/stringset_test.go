package utils

import (
	"testing"
)

func TestStringSet(t *testing.T) {
	//var set *StringSet
	set := NewStringSet()
	set.Add("shangsan")
	set.Add("lisi")
	set.Add("shangsan")

	x := len(set.ToArray())
	if x == 2 {
		t.Logf("success, %v", set.ToArray())
	} else {
		t.Errorf("error")
	}
}

func TestStringSetInterSection(t *testing.T) {
	var a, b *StringSet
	a = NewStringSet()
	b = NewStringSet()
	a.Add("dog")
	a.Add("cat")
	a.Add("fish")
	a.Add("cat")
	a.Add("tortoise")

	b.Add("deer")
	b.Add("dog")
	b.Add("cat")
	b.Add("fish")

	x := a.Intersection(b)

	if len(x.InternalMap) == 3 {
		t.Logf("success, %v", x.ToArray())
	} else {
		t.Errorf("error")
	}
}
