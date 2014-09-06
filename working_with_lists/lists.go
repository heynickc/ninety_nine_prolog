package lists

import (
	"errors"
)

// P01
// Find the last element of a list.
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

// P02
// Find the last but one element of a list.

func (s StringSlice) LastElementButOne() (string, error) {
	if len(s) < 2 {
		return "", errors.New("String slice isn't long enough!")
	}
	return s[len(s)-2], nil
}

// P03
// Find the K'th element of a list.
// The first element in the list is number 1.
// Example:
// ?- element_at(X,[a,b,c,d,e],3).
// X = c

func (s StringSlice) ElementAtIndex(index int) (string, error) {
	k := index - 1
	if len(s) < 1 {
		return "", errors.New("String slice isn't long enough!")
	}
	return s[k], nil
}
