package adaptor

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	for i := 0; i < 10; i++ {
		q.Push(i)
	}

	for i := 0; i < 10; i++ {
		x := q.Pop()
		if x == i {
			t.Logf("success, %v", x)
		} else {
			t.Errorf("error")
		}
	}
}
