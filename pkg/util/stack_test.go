package util

import "testing"

func TestStack(t *testing.T) {
	in := []int{1, 2, 3}

	s := Stack{}

	for _, n := range in {
		s.Push(n)
	}

	for i := len(in); i > 0; i-- {
		n := s.Pop()
		if n != i {
			t.Fatalf("Expected %v, got %v\n", i, n)
		}
	}

	if len(s) != 0 {
		t.Fail()
	}
}
