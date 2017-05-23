package utils

// Set of strings
// Implements "Contains" in O(1) (slice runs in O(n))
type Set struct {
	data map[string]bool
}

func NewSet() Set {
	return Set{data: make(map[string]bool)}
}

func (set Set) Add(val string) {
	set.data[val] = true
}

func (set Set) Contains(val string) bool {
	_, ok := set.data[val]

	return ok
}

func (set Set) ToArray() []string {

	var ret []string
	for key := range set.data {
		ret = append(ret, key)
	}

	return ret
}
