package util

import "testing"

func TestReadIntArray(t *testing.T) {
	g := ReadToIntArray("../../data/11-test2.txt")

	if len(g) != 5 {
		t.Fail()
	}

	if len(g[0]) != 5 {
		t.Fail()
	}
}
