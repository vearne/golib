package utils

import (
	"testing"
)

func TestIntSetInterSection(t *testing.T) {

	var a, b *IntSet
	a = NewIntSet()
	b = NewIntSet()
	a.Add(1)
	a.Add(3)
	a.Add(5)
	a.Add(7)

	b.Add(3)
	b.Add(5)
	b.Add(7)

	x := a.Intersection(b)

	if len(x.InternalMap) == 3 {
		t.Logf("success, %v", x.ToArray())
	} else {
		t.Errorf("error")
	}
}
