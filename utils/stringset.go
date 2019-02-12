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

func (set *StringSet) Remove(str string) {
	delete(set.InternalMap, str)
}

func (set *StringSet) ToArray() []string {
	res := make([]string, len(set.InternalMap))
	i := 0
	for key, _ := range set.InternalMap {
		res[i] = key
		i++
	}
	return res
}
