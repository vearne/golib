package utils

import (
	"testing"
)

func TestInterSection1(t *testing.T) {

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
	b.Add(9)

	x := a.Intersection(b)


	if  len(x.InternalMap) == 3{
		t.Logf("success, %v", x)
	} else {
		t.Errorf("error")
	}
}