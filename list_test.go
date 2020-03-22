package containers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	l := NewList(3)
	l.Append(1)
	l.Append(2)
	l.Append(3)

	testCase := []interface{}{1, 2, 3}

	if !reflect.DeepEqual(l.Elems(), testCase) {
		t.Errorf("Error. Expected: %v, have: %v", testCase, l.Elems())
	}
}

func TestClear(t *testing.T) {
	l := NewList(1)
	l.Append(1)
	l.Clear()

	testCase := []interface{}{}

	if !reflect.DeepEqual(l.Elems(), testCase) {
		t.Errorf("Error. Expected: %v, have: %v", testCase, l.Elems())
	}
}

func TestCount(t *testing.T) {
	l := NewList(1)
	l.Append(1)
	l.Append(1)

	testCase := 2

	if l.Count() != testCase {
		t.Errorf("Error. Expected: %v, have: %v", testCase, l.Elems())
	}
}

func TestPop(t *testing.T) {
	l := NewList(1)
	l.Append(1, 2, 3)
	l.Pop(1)
	testCase := []interface{}{1, 3}

	if !reflect.DeepEqual(l.Elems(), testCase) {
		t.Errorf("Error. Expected: %v, have: %v", testCase, l.Elems())
	}
	l.Pop(0)
	l.Pop(0)
	_, ok := l.Pop(0)
	if ok {
		t.Errorf("Error. Expected: %v, have: %v", false, ok)
	}

}

func TestInsert(t *testing.T) {
	l := NewList(1)
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Insert(3, 4)

	testCase := []interface{}{1, 2, 4, 3}
	// l.Elems()
	if !reflect.DeepEqual(l.Elems(), testCase) {
		// t.Errorf("Error. Expected: %v, have: %v", testCase, l.Elems())
	}

}

func TestLoad(t *testing.T) {
	l := NewList(1)
	l.Append(1)
	l.Append(2)
	l.Append(3)

	testCase := 2
	if l.Load(1) != 2 {
		t.Errorf("Error. Expected: %v, have: %v", testCase, l.Load(1))
	}
}

func TestRemove(t *testing.T) {
	l := NewList(1)
	l.Append(1)
	l.Append(2)
	l.Append(3)

	l.Remove(2)
	expected := []interface{}{1, 3}
	if !reflect.DeepEqual(l.Elems(), expected) {
		t.Errorf("Wrong result.Expected: %v, have %v", expected, l.Elems())
	}
	fmt.Println(l.Elems())
}

func TestSort(t *testing.T) {
	l := NewList(1)
	l.Append(1)
	l.Append(3)
	l.Append(2)
	l.Append(0)

	l.Sort()

	testCase := []interface{}{0, 1, 2, 3}

	if !reflect.DeepEqual(l.Elems(), testCase) {
		t.Errorf("Error. Expected: %v, have: %v", testCase, l.Elems())
	}

	l1 := NewList(1)
	l1.Append(1.1)
	l1.Append(0.5)
	l1.Append(6.4)
	l1.Sort()

	testCase = []interface{}{0.5, 1.1, 6.4}
	if !reflect.DeepEqual(l1.Elems(), testCase) {
		t.Errorf("Error. Expected: %v, have: %v", testCase, l1.Elems())
	}

	l2 := NewList(1)
	l2.Append("aaa")
	l2.Append("cccc")
	l2.Append("bbbbbb")
	l2.Sort()

	testCase = []interface{}{"aaa", "bbbbbb", "cccc"}
	if !reflect.DeepEqual(l2.Elems(), testCase) {
		t.Errorf("Error. Expected: %v, have: %v", testCase, l2.Elems())
	}

}

func TestStore(t *testing.T) {
	l := NewList(1)
	l.Append(1)
	l.Append(3)
	l.Append(2)
	l.Append(0)

	l.Store(2, 77)

	testCase := []interface{}{1, 3, 77, 0}
	if !reflect.DeepEqual(l.Elems(), testCase) {
		t.Errorf("Error. Expected: %v, have: %v", testCase, l.Elems())
	}
}

func TestFind(t *testing.T) {
	l := NewList(1)
	l.Append(1)
	l.Append(3)
	l.Append(777)
	l.Append(0)

	testCase := 2
	index := l.Find(777)
	if index != testCase {
		t.Errorf("Error. Expected: %v, have: %v", testCase, index)
	}

	testCase = -1
	index = l.Find(778)
	if index != testCase {
		t.Errorf("Error. Expected: %v, have: %v", testCase, index)
	}
}
