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
	Remove(x interface{}) bool
	Union(s1 SetInterface) []interface{}
}

func newSet(size int) SetInterface {
	return &Set{
		data: make(map[interface{}]struct{}, size),
	}
}

type Set struct {
	data map[interface{}]struct{}
}

func (s *Set) Add(x ...interface{}) {
	for _, v := range x {
		s.data[v] = struct{}{}
	}
}

func (s *Set) Clear() {
	s.data = make(map[interface{}]struct{}, 0)
}

func (s *Set) Count() int {
	return len(s.data)
}

func (s *Set) Difference(s1 SetInterface) []interface{} {
	diff := make([]interface{}, 0, len(s.data)/2)
	data := s.Elems()

	for _, k := range data {
		if !s1.Find(k) {
			diff = append(diff, k)
		}
	}

	return diff
}

func (s *Set) Elems() []interface{} {
	data := make([]interface{}, 0, len(s.data))
	for key := range s.data {
		data = append(data, key)
	}

	return data
}

func (s *Set) Intersection(s1 SetInterface) []interface{} {
	intersec := make([]interface{}, 0, len(s.Elems())/2)
	allKeys := make([]interface{}, 0, len(s.Elems())+len(s1.Elems()))

	for _, key := range s.Elems() {
		allKeys = append(allKeys, key)
	}

	for _, key := range s1.Elems() {
		allKeys = append(allKeys, key)
	}

	for _, key := range allKeys {
		if s.Find(key) && s1.Find(key) {
			intersec = append(intersec, key)
		}
	}

	return intersec
}

func (s *Set) Pop() (interface{}, bool) {
	if len(s.data) == 0 {
		return nil, false
	}

	var key interface{}
	for k := range s.data {
		key = k
		break
	}

	return key, true
}

func (s *Set) Remove(x interface{}) bool {
	if _, ok := s.data[x]; !ok {
		return false
	}
	delete(s.data, x)

	return true
}

func (s *Set) Union(s1 SetInterface) []interface{} {
	union := make([]interface{}, 0, len(s.Elems())/2)
	allKeys := make([]interface{}, 0, len(s.Elems())+len(s1.Elems()))

	for _, key := range s.Elems() {
		allKeys = append(allKeys, key)
	}

	for _, key := range s1.Elems() {
		allKeys = append(allKeys, key)
	}

	for _, key := range allKeys {
		if s.Find(key) || s1.Find(key) {
			union = append(union, key)
		}
	}

	return union
}

func (s *Set) Find(x interface{}) bool {
	if _, ok := s.data[x]; !ok {
		return false
	}

	return true
}
