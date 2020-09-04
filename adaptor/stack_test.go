package adaptor

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()
	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	for i := 9; i >= 0; i-- {
		x := s.Pop()
		if x == i {
			t.Logf("success, %v", x)
		} else {
			t.Errorf("error")
		}
	}
}
