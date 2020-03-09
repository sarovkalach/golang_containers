package containers

import (
	"fmt"
	"reflect"
	"sort"
)

type ListInterface interface {
	Append(x ...interface{})
	Clear()
	Count() int
	Elems() []interface{}
	Pop() (interface{}, bool)
	Insert(int, interface{})
	Load(int) interface{}
	Remove(x interface{}) bool
	Sort()
	Store(int, interface{})
	Find(x interface{}) int
}

func NewList(size int) ListInterface {
	return &List{data: make([]interface{}, 0, size)}
}

type List struct {
	data []interface{}
}

func (l *List) Append(x ...interface{}) {
	l.data = append(l.data, x...)
}

func (l *List) Clear() {
	l.data = make([]interface{}, 0)
}

func (l *List) Count() int {
	return len(l.data)
}

func (l *List) Elems() []interface{} {
	return l.data
}

func (l *List) Pop() (interface{}, bool) {
	var (
		ok   bool
		elem interface{}
	)
	if l.Count()-1 < 0 {
		return nil, false
	}

	elem = l.data[l.Count()-1]
	l.data = l.data[:l.Count()-1]

	return elem, ok
}

func (l *List) Insert(index int, x interface{}) {
	p1 := l.data[:index-1]
	p1 = append(l.data, x)
	l.data = append(l.data[index:], p1...)
}

func (l *List) Load(index int) interface{} {
	return l.data[index]
}

func (l *List) Remove(x interface{}) bool {
	for i, elem := range l.data {
		if elem == x {
			l.data = append(l.data[:i], l.data[i+1:]...)
			return true
		}
	}
	return false
}

func (l *List) Sort() {
	if l.Count() == 0 {
		fmt.Println("Nothing to sort")
		return
	}
	elem := l.data[0]
	switch reflect.TypeOf(elem).Kind() {
	case reflect.Int:
		tmp := make([]int, l.Count())
		for i, elem := range l.data {
			tmp[i] = elem.(int)
		}
		sort.Ints(tmp)
		for i := range l.data {
			l.data[i] = tmp[i]
		}
		return

	case reflect.Float64:
		tmp := make([]float64, l.Count())
		for i, elem := range l.data {
			tmp[i] = elem.(float64)
		}
		sort.Float64s(tmp)
		for i := range l.data {
			l.data[i] = tmp[i]
		}

	case reflect.String:
		tmp := make([]string, l.Count())
		for i, elem := range l.data {
			tmp[i] = elem.(string)
		}
		sort.Strings(tmp)
		for i := range l.data {
			l.data[i] = tmp[i]
		}
	}
}

func (l *List) Store(index int, x interface{}) {
	l.data[index] = x
}

func (l *List) Find(x interface{}) int {
	for i, elem := range l.data {
		if elem == x {
			return i
		}
	}

	return -1
}
