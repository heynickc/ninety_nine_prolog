package lists

import (
	"errors"
	// "fmt"
)

// P01 (*) Find the last element of a list.
// Example:
// ?- my_last(X,[a,b,c,d]).
// X = d

type StringSlice []string

func (s StringSlice) LastElement() (string, error) {
	if len(s) < 1 {
		return "", errors.New("Nothing in the string slice!")
	}
	return s[len(s)-1], nil
}

// P02 (*) Find the last but one element of a list.
// (zweitletztes Element, l'avant-dernier élément)

func (s StringSlice) LastElementButOne() (string, error) {
	if len(s) < 2 {
		return "", errors.New("String slice isn't long enough!")
	}
	return s[len(s)-2], nil
}

// P03 (*) Find the K'th element of a list.
// The first element in the list is number 1.
// Example:
// ?- element_at(X,[a,b,c,d,e],3).
// X = c

func (s StringSlice) NthElement(index int) (string, error) {
	k := index - 1
	if len(s) < 1 {
		return "", errors.New("String slice isn't long enough!")
	}
	return s[k], nil
}

// P04 (*) Find the number of elements of a list.

func (s StringSlice) Length() int {
	return len(s)
}

// SIDEBAR: Rotate a list...
// http://www.cs.bell-labs.com/cm/cs/pearls/rotate.c
// Programming Pearls 2.3

func (s *StringSlice) ReverseAt(i, j int) {
	sCopy := *s
	var t string
	for i < j {
		t = sCopy[i]
		sCopy[i] = sCopy[j]
		sCopy[j] = t
		i++
		j--
	}
	*s = sCopy
}

func (s *StringSlice) Rotate(rotDist, n int) {
	s.ReverseAt(0, rotDist-1)
	s.ReverseAt(rotDist, n-1)
	s.ReverseAt(0, n-1)
}

// P05 (*) Reverse a list.
// Uses ReverseAt()

func (s *StringSlice) Reverse() {
	s.ReverseAt(0, s.Length()-1)
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
