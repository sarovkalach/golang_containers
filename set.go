package containers

type SetInterface interface {
	Add(x ...interface{})
	Clear()
	Count() int
	Difference()
	Elems() []interface{}
	Intersection()
	Pop() (interface{}, bool)
	Remove(x interface{}) bool
	Union()
	Find(x interface{}) int
}

func newSet(size int) *Set {
	return &Set{
		data: make(map[interface{}]struct{}, size),
	}
}

type Set struct {
	data map[interface{}]struct{}
}
