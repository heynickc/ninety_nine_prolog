package lists

import (
	"errors"
	"fmt"
)

type StringSlice []string
type NestedSlice []interface{}

// P01 (*) Find the last element of a list.
// Example:
// ?- my_last(X,[a,b,c,d]).
// X = d

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
	if len(s) < index {
		return "", errors.New("String slice isn't long enough!")
	}
	return s[k], nil
}

// P04 (*) Find the number of elements of a list.

func (s StringSlice) Length() int {
	return len(s)
}

// SIDEBAR: Rotate a list...
// Instead of full reversal, this will rotate the
// list by rotDist # of elements
// http://www.cs.bell-labs.com/cm/cs/pearls/rotate.c
// Programming Pearls 2.3

func (s StringSlice) ReversePortion(i, j int) StringSlice {
	result := make(StringSlice, len(s))
	copy(result, s)
	var t string
	for i < j {
		t = result[i]
		result[i] = result[j]
		result[j] = t
		i++
		j--
	}
	return result
}

func (s StringSlice) Rotate(rotDist, n int) StringSlice {
	result := make(StringSlice, len(s))
	copy(result, s)
	result = result.ReversePortion(0, rotDist-1)
	result = result.ReversePortion(rotDist, n-1)
	result = result.ReversePortion(0, n-1)
	return result
}

// P05 (*) Reverse a list.
// Uses ReversePortion() algorithm

func (s StringSlice) Reverse() StringSlice {
	result := make(StringSlice, len(s))
	copy(result, s)
	result = result.ReversePortion(0, s.Length()-1)
	return result
}

// P06 (*) Find out whether a list is a palindrome.
// A palindrome can be read forward or backward; e.g. [x,a,m,a,x].

func (s StringSlice) IsPalindrome() bool {
	var i, mid, last int
	i = 0
	mid = s.Length() / 2
	last = s.Length() - 1
	for i < mid {
		if s[i] != s[last-i] {
			return false
		}
		i++
	}
	return true
}

// P07 (**) Flatten a nested list structure.
// Transform a list, possibly holding lists as elements into a `flat' list by replacing each list with its elements (recursively).

// Example:
// ?- my_flatten([a, [b, [c, d], e]], X).
// X = [a, b, c, d, e]

// Hint: Use the predefined predicates is_list/1 and append/3

// http://rosettacode.org/wiki/Flatten_a_list#Go
// Uses type switching http://golang.org/doc/effective_go.html#type_switch

func (n NestedSlice) Flatten() StringSlice {
	var result []string
	for _, s := range n {
		switch i := s.(type) {
		case string:
			result = append(result, i)
		case NestedSlice:
			result = append(result, i.Flatten()...)
		}
	}
	return result
}

// P08 (**) Eliminate consecutive duplicates of list elements.
// If a list contains repeated elements they should be replaced with a single copy of the element. The order of the elements should not be changed.

// Example:
// ?- compress([a,a,a,a,b,c,c,a,a,d,e,e,e,e],X).
// X = [a,b,c,a,d,e]

func (s StringSlice) Compress() StringSlice {
	var result StringSlice
	for _, item := range s {
		if last, _ := result.LastElement(); item != last {
			result = append(result, item)
		}
	}
	return result
}

// P09 (**) Pack consecutive duplicates of list elements into sublists.
// If a list contains repeated elements they should be placed in separate sublists.

// Example:
// ?- pack([a,a,a,a,b,c,c,a,a,d,e,e,e,e],X).
// X = [[a,a,a,a],[b],[c,c],[a,a],[d],[e,e,e,e]]

func (s StringSlice) Pack() NestedSlice {
	var result NestedSlice
	var group StringSlice
	for _, item := range s {
		if len(group) == 0 {
			group = append(group, item)
		} else {
			if last := group[len(group)-1]; item != last {
				group = append(group, item)
			}
		}
		// } else {
		// 	result = append(result, group)
		// 	group = make(StringSlice, 0)
		// }
		fmt.Println(group)
	}
	return result
}

// P10 (*) Run-length encoding of a list.
// Use the result of problem P09 to implement the so-called run-length encoding data compression method. Consecutive duplicates of elements are encoded as terms [N,E] where N is the number of duplicates of the element E.

// Example:
// ?- encode([a,a,a,a,b,c,c,a,a,d,e,e,e,e],X).
// X = [[4,a],[1,b],[2,c],[2,a],[1,d][4,e]]
