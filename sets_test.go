package containers

import (
	"testing"
)

func TestSetClear(t *testing.T) {
	s := NewSet(1)
	s.Add(1, 2, 3)

	if len(s.data) != 3 {
		t.Errorf("Wrong size. Expected : %d, Have: %d", 3, len(s.data))
	}

	s.Clear()
	if len(s.data) != 0 {
		t.Errorf("Wrong size. Expected : %d, Have: %d", 0, len(s.data))
	}

}

func TestSetDifference(t *testing.T) {
	s := NewSet(1)
	s.Add(1, 2, 3)

	s1 := NewSet(1)
	s1.Add(3, 4, 5, 6)

	difference := s.Difference(s1)
	expectedSize := 2
	if len(difference) != expectedSize {
		t.Errorf("Wrong result. Expected: %v, Have: %v", expectedSize, len(difference))
	}
}

func TestSetFind(t *testing.T) {
	s := NewSet(1)
	s.Add(1, 2, 3)

	if !s.Find(2) {
		t.Errorf("Error. This elem must be exist's")
	}

	if s.Find(4) {
		t.Error("Error. This elem must not be exist's")
	}
}

func TestSetIntersection(t *testing.T) {
	s := NewSet(1)
	s.Add(1, 2, 3)

	s1 := NewSet(1)
	s1.Add(2, 3, 4)

	intersection := s.Intersection(s1)
	expectedSize := 2
	if len(intersection) != expectedSize {
		t.Errorf("Wrong result. Expected: %v, Have: %v", expectedSize, len(intersection))
	}
}

func TestSetPop(t *testing.T) {
	s := NewSet(1)
	s.Add(1, 2, 3)
	s.Pop()

	expectedSize := 2
	if s.Count() != expectedSize {
		t.Errorf("Wrong result. Expected: %v, Have: %v", expectedSize, s.Count())
	}
}

func TestSetRemove(t *testing.T) {
	s := NewSet(1)
	s.Add(1, 2, 3)
	s.Remove(2)

	if s.Find(2) {
		t.Error("Error. This elem must not be exist's")
	}

	if !s.Find(1) {
		t.Error("Error. This elem must  be exist's")
	}
}

func TestSetUnion(t *testing.T) {
	s := NewSet(1)
	s.Add(1, 2, 3)

	s1 := NewSet(1)
	s1.Add(2, 3, 4)

	union := s.Union(s1)
	expectedSize := 4
	if len(union) != expectedSize {
		t.Errorf("Wrong result. Expected: %v, Have: %v", expectedSize, len(union))
	}
}
