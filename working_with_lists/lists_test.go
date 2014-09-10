package lists

import (
	"testing"
)

var testSlice = StringSlice{"a", "b", "c", "d"}
var emptyTestSlice = StringSlice{}

// P01 (*) Find the last element of a list.
// Example:
// ?- my_last(X,[a,b,c,d]).
// X = d

func TestLastElement(t *testing.T) {
	result, _ := testSlice.LastElement()
	if result != "d" {
		t.Errorf("Expected %v but got %v", "d", result)
	}
}

func TestLastElementThrowsError(t *testing.T) {
	result, err := emptyTestSlice.LastElement()
	if result != "" {
		t.Errorf("Expected to throw error: ", err)
	}
}

// P02 (*) Find the last but one element of a list.
// (zweitletztes Element, l'avant-dernier élément)

func TestLastElementButOne(t *testing.T) {
	result, _ := testSlice.LastElementButOne()
	if result != "c" {
		t.Errorf("Expected %v but got %v", "c", result)
	}
}

func TestLastElementButOneThrowsError(t *testing.T) {
	result, err := emptyTestSlice.LastElementButOne()
	if result != "" {
		t.Errorf("Expected to throw error: ", err)
	}
}

// P03 (*) Find the K'th element of a list.
// The first element in the list is number 1.
// Example:
// ?- element_at(X,[a,b,c,d,e],3).
// X = c

func TestNthElement(t *testing.T) {
	result, _ := testSlice.NthElement(3)
	if result != "c" {
		t.Errorf("Expected %v but got %v", "d", result)
	}
}

func TestNthElementThrowsError(t *testing.T) {
	result, err := emptyTestSlice.NthElement(3)
	if result != "" {
		t.Errorf("Expected to throw error: ", err)
	}
}

// P04 (*) Find the number of elements of a list.

func TestLength(t *testing.T) {
	result := testSlice.Length()
	if result != 4 {
		t.Errorf("Expected %v but got %v", 4, result)
	}
}

func TestLengthEmptySlicer(t *testing.T) {
	result := emptyTestSlice.Length()
	if result != 0 {
		t.Errorf("Expected %v but got %v", 0, result)
	}
}

// P05 (*) Reverse a list.

// SIDEBAR: Rotate a list...

func TestReverse(t *testing.T) {

	strToRot := StringSlice{"a", "b", "c", "d", "e", "f", "g", "h"}

	rotdStr := make(StringSlice, strToRot.Length())
	copy(rotdStr, strToRot)
	rotdStr.Reverse()

	for i, j := 0, strToRot.Length()-1; i < strToRot.Length(); i, j = i+1, j-1 {
		if strToRot[i] != rotdStr[j] {
			t.Errorf("Expected %v to equal %v", strToRot[i], rotdStr[j])
		}
	}
}

// P06 (*) Find out whether a list is a palindrome.
// A palindrome can be read forward or backward; e.g. [x,a,m,a,x].

// P07 (**) Flatten a nested list structure.
// Transform a list, possibly holding lists as elements into a `flat' list by replacing each list with its elements (recursively).

// Example:
// ?- my_flatten([a, [b, [c, d], e]], X).
// X = [a, b, c, d, e]

// Hint: Use the predefined predicates is_list/1 and append/3

// P08 (**) Eliminate consecutive duplicates of list elements.
// If a list contains repeated elements they should be replaced with a single copy of the element. The order of the elements should not be changed.

// Example:
// ?- compress([a,a,a,a,b,c,c,a,a,d,e,e,e,e],X).
// X = [a,b,c,a,d,e]

// P09 (**) Pack consecutive duplicates of list elements into sublists.
// If a list contains repeated elements they should be placed in separate sublists.

// Example:
// ?- pack([a,a,a,a,b,c,c,a,a,d,e,e,e,e],X).
// X = [[a,a,a,a],[b],[c,c],[a,a],[d],[e,e,e,e]]

// P10 (*) Run-length encoding of a list.
// Use the result of problem P09 to implement the so-called run-length encoding data compression method. Consecutive duplicates of elements are encoded as terms [N,E] where N is the number of duplicates of the element E.

// Example:
// ?- encode([a,a,a,a,b,c,c,a,a,d,e,e,e,e],X).
// X = [[4,a],[1,b],[2,c],[2,a],[1,d][4,e]]
