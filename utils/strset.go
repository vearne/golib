package utils

type StrSet struct {
	InternalMap map[string]int
}

func NewStrSet() *StrSet {
	return &StrSet{InternalMap: make(map[string]int)}
}

func (set *StrSet) Add(str string) {
	set.InternalMap[str] = 0
}

func (set *StrSet) Intersection(set2 *StrSet) (*StrSet) {
	result := NewStrSet()
	for key, _ := range set.InternalMap {
		if _, ok := set2.InternalMap[key]; ok {
			result.Add(key)
		}
	}
	return result
}

func (set *StrSet) ToSlice() []string{
	res := make([]string, 0, 5)
	for key:=range set.InternalMap{
		res = append(res, key)
	}
	return res
}