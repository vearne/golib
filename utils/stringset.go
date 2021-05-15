package utils

type StringSet struct {
	InternalMap map[string]int
}

func NewStringSet() *StringSet {
	return &StringSet{InternalMap: make(map[string]int)}
}

func (set *StringSet) Add(str string) {
	set.InternalMap[str] = 1
}

func (set *StringSet) AddAll(itemSlice []string) {
	for _, item := range itemSlice {
		set.InternalMap[item] = 1
	}
}

func (set *StringSet) Has(str string) bool {
	_, ok := set.InternalMap[str]
	return ok
}

func (set *StringSet) Remove(str string) {
	delete(set.InternalMap, str)
}

func (set *StringSet) RemoveAll(other *StringSet) {
	for _, item := range other.ToArray() {
		delete(set.InternalMap, item)
	}
}

func (set *StringSet) ToArray() []string {
	res := make([]string, len(set.InternalMap))
	i := 0
	for key := range set.InternalMap {
		res[i] = key
		i++
	}
	return res
}

func (set *StringSet) Size() int {
	return len(set.InternalMap)
}

func (set *StringSet) Intersection(set2 *StringSet) *StringSet {
	result := NewStringSet()

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

func (set *StringSet) Clone() *StringSet {
	result := NewStringSet()
	result.AddAll(set.ToArray())
	return result
}
