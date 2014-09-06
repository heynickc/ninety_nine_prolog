package lists

import (
	"testing"
)

// P01
// Find the last element of a list.
// Example:
// ?- my_last(X,[a,b,c,d]).
// X = d

func TestLastElement(t *testing.T) {
	testSlice := StringSlice{"a", "b", "c", "d"}
	result, _ := testSlice.LastElement()
	if result != "d" {
		t.Errorf("Expected %v but got %v", "d", result)
	}
}

func TestLastElementThrowsError(t *testing.T) {
	testSlice := StringSlice{}
	result, err := testSlice.LastElement()
	if result != "" {
		t.Errorf("Expected to throw error: ", err)
	}
}

// P02
// Find the last but one element of a list.

func TestLastElementButOne(t *testing.T) {
	testSlice := StringSlice{"a", "b", "c", "d"}
	result, _ := testSlice.LastElementButOne()
	if result != "c" {
		t.Errorf("Expected %v but got %v", "c", result)
	}
}

func TestLastElementButOneThrowsError(t *testing.T) {
	testSlice := StringSlice{}
	result, err := testSlice.LastElementButOne()
	if result != "" {
		t.Errorf("Expected to throw error: ", err)
	}
}

// P03
// Find the K'th element of a list.
// The first element in the list is number 1.
// Example:
// ?- element_at(X,[a,b,c,d,e],3).
// X = c

func TestElementAtIndex(t *testing.T) {
	testSlice := StringSlice{"a", "b", "c", "d", "e"}
	result, _ := testSlice.ElementAtIndex(3)
	if result != "c" {
		t.Errorf("Expected %v but got %v", "d", result)
	}
}

func TestElementAtIndexThrowsError(t *testing.T) {
	testSlice := StringSlice{}
	result, err := testSlice.ElementAtIndex(3)
	if result != "" {
		t.Errorf("Expected to throw error: ", err)
	}
}
