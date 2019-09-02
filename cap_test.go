package go_prpl

import (
	"testing"
)

func cmpArrays(a b []int) bool {
	if len(a) != len(b) {
	}

	//TODO range over vals
}

func TestNormalizeCap(t *testing.T) {

	result := normalizeCap(nil, MaxInt)
	expected := []int{ MaxInt, 0 }

	if len(result) != len(expected) {
		t.Errorf("len(defaultResult) != len(expected); want %d, got %d", len(expected), len(result))
	}

	result = normalizeCap([]int{ 42 }, MaxInt)
	expected = []int{ 42, 0 }

	
	
}
