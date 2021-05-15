package utils

type IntSet struct {
	InternalMap map[int]int
}

func NewIntSet() *IntSet {
	return &IntSet{InternalMap: make(map[int]int)}
}
func (set *IntSet) Add(i int) {
	set.InternalMap[i] = 1
}

func (set *IntSet) AddAll(itemSlice []int) {
	for _, item := range itemSlice {
		set.InternalMap[item] = 1
	}
}

func (set *IntSet) Has(key int) bool {
	_, ok := set.InternalMap[key]
	return ok
}

func (set *IntSet) Remove(key int) {
	delete(set.InternalMap, key)
}

func (set *IntSet) RemoveAll(other *IntSet) {
	for _, item := range other.ToArray() {
		delete(set.InternalMap, item)
	}
}

func (set *IntSet) Size() int {
	return len(set.InternalMap)
}

func (set *IntSet) Intersection(set2 *IntSet) *IntSet {
	result := NewIntSet()
	if set.Size() > set2.Size() {
		set, set2 = set2, set
	}

	for key := range set.InternalMap {
		if _, ok := set2.InternalMap[key]; ok {
			result.Add(key)
		}
	}
	return result
}

func (set *IntSet) ToArray() []int {
	res := make([]int, 0, 5)
	for key := range set.InternalMap {
		res = append(res, key)
	}
	return res
}

func (set *IntSet) Clone() *IntSet {
	result := NewIntSet()
	result.AddAll(set.ToArray())
	return result
}
