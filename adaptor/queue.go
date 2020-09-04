package adaptor

import (
	"container/list"
)

type Queue struct {
	List *list.List
}

func NewQueue() *Queue {
	q := Queue{}
	q.List = list.New()
	return &q
}

func (q *Queue) Empty() bool {
	return q.List.Len() == 0
}

func (q *Queue) Size() int {
	return q.List.Len()
}

func (q *Queue) Front() interface{} {
	return q.List.Front().Value
}

func (q *Queue) Back() interface{} {
	return q.List.Back().Value
}

func (q *Queue) Push(v interface{}) {
	q.List.PushBack(v)
}

// from front
func (q *Queue) Pop() interface{} {
	e := q.List.Front()
	q.List.Remove(e)
	return e.Value
}
