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
	Find(x interface{}) int
	Insert(int, interface{})
	Load(int) interface{}
	Pop(pos int) (interface{}, bool)
	Remove(x ...interface{})
	Sort()
	Store(int, interface{})
}

func NewList(size int) ListInterface {
	return &List{data: make([]interface{}, 0, size)}
}

type List struct {
	data []interface{}
}

// Append to the end
func (l *List) Append(x ...interface{}) {
	l.data = append(l.data, x...)
}

// Clear all elems
func (l *List) Clear() {
	l.data = make([]interface{}, 0)
}

// Count elements
func (l *List) Count() int {
	return len(l.data)
}

// Elems return data
func (l *List) Elems() []interface{} {
	return l.data
}

// Pop elem at index
func (l *List) Pop(index int) (interface{}, bool) {
	var (
		ok   bool
		elem interface{}
	)

	if index > l.Count()-1 {
		return nil, false
	}
	elem = l.data[index]
	l.data = append(l.data[:index], l.data[index+1:]...)

	return elem, ok
}

// Insert elem at index
func (l *List) Insert(index int, x interface{}) {
	if index < 0 || index >= l.Count() {
		fmt.Println("Wrong index")
		return
	}
	p1 := l.data[:index-1]
	p1 = append(l.data, x)
	l.data = append(l.data[index:], p1...)
}

// Load elem at index
func (l *List) Load(index int) interface{} {
	if index < 0 || index >= l.Count() {
		fmt.Println("Wrong index")
		return nil
	}
	return l.data[index]
}

// Remove elem by value
func (l *List) Remove(x ...interface{}) {
	for _, val := range x {
		for i, elem := range l.data {
			if elem == val {
				l.data = append(l.data[:i], l.data[i+1:]...)
				break
			}
		}
	}
}

// Sort elems
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

// Store elem by index
func (l *List) Store(index int, x interface{}) {
	if index < 0 || index >= l.Count() {
		fmt.Println("Wrong index")
		return
	}
	l.data[index] = x
}

// Find elem by value
func (l *List) Find(x interface{}) int {
	for i, elem := range l.data {
		if elem == x {
			return i
		}
	}

	return -1
}
