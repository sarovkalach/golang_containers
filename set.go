package containers

type SetInterface interface {
	Add(x ...interface{})
	Clear()
	Count() int
	Difference(s1 SetInterface) []interface{}
	Elems() []interface{}
	Find(x interface{}) bool
	Intersection(s1 SetInterface) []interface{}
	Pop() (interface{}, bool)
	Remove(x ...interface{})
	Union(s1 SetInterface) []interface{}
}

func NewSet(size int) *Set {
	return &Set{
		data: make(map[interface{}]struct{}, size),
	}
}

type Set struct {
	data map[interface{}]struct{}
}

// Add element
func (s *Set) Add(x ...interface{}) {
	for _, v := range x {
		s.data[v] = struct{}{}
	}
}

// Clear all elems
func (s *Set) Clear() {
	s.data = make(map[interface{}]struct{}, 0)
}

// Count elems
func (s *Set) Count() int {
	return len(s.data)
}

// Difference between  s1
func (s *Set) Difference(s1 SetInterface) []interface{} {
	difference := make([]interface{}, 0, len(s.Elems())/2)
	allKeys := make(map[interface{}]struct{})

	for _, key := range s.Elems() {
		allKeys[key] = struct{}{}
	}

	for _, key := range s1.Elems() {
		allKeys[key] = struct{}{}
	}

	for key := range allKeys {
		if s.Find(key) && !s1.Find(key) {
			difference = append(difference, key)
		}
	}

	return difference
}

// Elems show data
func (s *Set) Elems() []interface{} {
	data := make([]interface{}, 0, len(s.data))
	for key := range s.data {
		data = append(data, key)
	}

	return data
}

// Intersection with s1
func (s *Set) Intersection(s1 SetInterface) []interface{} {
	intersec := make([]interface{}, 0, len(s.Elems())/2)
	allKeys := make(map[interface{}]struct{})

	for _, key := range s.Elems() {
		allKeys[key] = struct{}{}
	}

	for _, key := range s1.Elems() {
		allKeys[key] = struct{}{}
	}

	for key := range allKeys {
		if s.Find(key) && s1.Find(key) {
			intersec = append(intersec, key)
		}
	}

	return intersec
}

// Pop random elem
func (s *Set) Pop() (interface{}, bool) {
	if len(s.data) == 0 {
		return nil, false
	}

	var key interface{}
	for k := range s.data {
		key = k
		break
	}
	delete(s.data, key)

	return key, true
}

// Remove elem
func (s *Set) Remove(x ...interface{}) {
	for _, val := range x {
		delete(s.data, val)
	}
}

// Union with s1
func (s *Set) Union(s1 SetInterface) []interface{} {
	union := make([]interface{}, 0, len(s.Elems())/2)
	allKeys := make(map[interface{}]struct{})

	for _, key := range s.Elems() {
		allKeys[key] = struct{}{}
	}

	for _, key := range s1.Elems() {
		allKeys[key] = struct{}{}
	}

	for key := range allKeys {
		if s.Find(key) || s1.Find(key) {
			union = append(union, key)
		}
	}

	return union
}

// Find elem
func (s *Set) Find(x interface{}) bool {
	if _, ok := s.data[x]; !ok {
		return false
	}

	return true
}
