package utils

type IntSet struct {
	InternalMap map[int]int
}

func NewIntSet() *IntSet {
	return &IntSet{InternalMap: make(map[int]int)}
}
func (set *IntSet) Add(i int) {
	set.InternalMap[i] = i
}

func (set *IntSet) Intersection(set2 *IntSet) (*IntSet) {
	result := NewIntSet()
	for key, _ := range set.InternalMap {
		if _, ok := set2.InternalMap[key]; ok {
			result.Add(key)
		}
	}
	return result
}
