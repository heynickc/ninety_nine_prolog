package lists

import (
	"errors"
)

type StringSlice []string
type GenericSlice []interface{}
type EncodedPair struct {
	int
	string
}
type EncodedSlice []EncodedPair

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
	mid = len(s) / 2
	last = len(s) - 1
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

func (n GenericSlice) Flatten() StringSlice {
	var result []string
	for _, s := range n {
		switch i := s.(type) {
		case string:
			result = append(result, i)
		case GenericSlice:
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
		if len(result) == 0 {
			result = append(result, item)
		} else if last := result[len(result)-1]; item != last {
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

func (s StringSlice) Pack() []StringSlice {
	var result []StringSlice
	var group StringSlice
	for i, item := range s {
		if len(group) == 0 {
			group = append(group, item)
		} else if last := group[len(group)-1]; item == last {
			group = append(group, item)
			if i == len(s)-1 {
				result = append(result, group)
			}
		} else if last := group[len(group)-1]; item != last {
			result = append(result, group)
			group = make(StringSlice, 0)
			group = append(group, item)
		}
	}
	return result
}

// P10 (*) Run-length encoding of a list.
// Use the result of problem P09 to implement the so-called run-length encoding data compression method. Consecutive duplicates of elements are encoded as terms [N,E] where N is the number of duplicates of the element E.

// Example:
// ?- encode([a,a,a,a,b,c,c,a,a,d,e,e,e,e],X).
// X = [[4,a],[1,b],[2,c],[2,a],[1,d][4,e]]

func (s StringSlice) Encode() EncodedSlice {
	var result EncodedSlice
	var compressed []StringSlice
	compressed = s.Pack()
	for _, item := range compressed {
		result = append(result, item.CompressCount())
	}
	return result
}

func (s StringSlice) CompressCount() EncodedPair {
	var char string
	var count int
	char = s[0]
	for i := 0; i < len(s); i++ {
		count++
	}
	return EncodedPair{count, char}
}

// P11 (*) Modified run-length encoding.
// Modify the result of problem P10 in such a way that if an element has no duplicates it is simply copied into the result list. Only elements with duplicates are transferred as [N,E] terms.

// Example:
// ?- encode_modified([a,a,a,a,b,c,c,a,a,d,e,e,e,e],X).
// X = [[4,a],b,[2,c],[2,a],d,[4,e]]

// P12 (**) Decode a run-length encoded list.
// Given a run-length code list generated as specified in problem P11. Construct its uncompressed version.

// P13 (**) Run-length encoding of a list (direct solution).
// Implement the so-called run-length encoding data compression method directly. I.e. don't explicitly create the sublists containing the duplicates, as in problem P09, but only count them. As in problem P11, simplify the result list by replacing the singleton terms [1,X] by X.

// Example:
// ?- encode_direct([a,a,a,a,b,c,c,a,a,d,e,e,e,e],X).
// X = [[4,a],b,[2,c],[2,a],d,[4,e]]

// P14 (*) Duplicate the elements of a list.
// Example:
// ?- dupli([a,b,c,c,d],X).
// X = [a,a,b,b,c,c,c,c,d,d]

// P15 (**) Duplicate the elements of a list a given number of times.
// Example:
// ?- dupli([a,b,c],3,X).
// X = [a,a,a,b,b,b,c,c,c]

// What are the results of the goal:
// ?- dupli(X,3,Y).

// P16 (**) Drop every N'th element from a list.
// Example:
// ?- drop([a,b,c,d,e,f,g,h,i,k],3,X).
// X = [a,b,d,e,g,h,k]

// P17 (*) Split a list into two parts; the length of the first part is given.
// Do not use any predefined predicates.

// Example:
// ?- split([a,b,c,d,e,f,g,h,i,k],3,L1,L2).
// L1 = [a,b,c]
// L2 = [d,e,f,g,h,i,k]

// P18 (**) Extract a slice from a list.
// Given two indices, I and K, the slice is the list containing the elements between the I'th and K'th element of the original list (both limits included). Start counting the elements with 1.

// Example:
// ?- slice([a,b,c,d,e,f,g,h,i,k],3,7,L).
// X = [c,d,e,f,g]

// P19 (**) Rotate a list N places to the left.
// Examples:
// ?- rotate([a,b,c,d,e,f,g,h],3,X).
// X = [d,e,f,g,h,a,b,c]

// ?- rotate([a,b,c,d,e,f,g,h],-2,X).
// X = [g,h,a,b,c,d,e,f]

// Hint: Use the predefined predicates length/2 and append/3, as well as the result of problem P17.

// P20 (*) Remove the K'th element from a list.
// Example:
// ?- remove_at(X,[a,b,c,d],2,R).
// X = b
// R = [a,c,d]

// P21 (*) Insert an element at a given position into a list.
// Example:
// ?- insert_at(alfa,[a,b,c,d],2,L).
// L = [a,alfa,b,c,d]

// P22 (*) Create a list containing all integers within a given range.
// Example:
// ?- range(4,9,L).
// L = [4,5,6,7,8,9]

// P23 (**) Extract a given number of randomly selected elements from a list.
// The selected items shall be put into a result list.
// Example:
// ?- rnd_select([a,b,c,d,e,f,g,h],3,L).
// L = [e,d,a]

// Hint: Use the built-in random number generator random/2 and the result of problem P20.

// P24 (*) Lotto: Draw N different random numbers from the set 1..M.
// The selected numbers shall be put into a result list.
// Example:
// ?- rnd_select(6,49,L).
// L = [23,1,17,33,21,37]

// Hint: Combine the solutions of problems P22 and P23.
// P25 (*) Generate a random permutation of the elements of a list.
// Example:
// ?- rnd_permu([a,b,c,d,e,f],L).
// L = [b,a,d,c,e,f]

// Hint: Use the solution of problem P23.

// P26 (**) Generate the combinations of K distinct objects chosen from the N elements of a list
// In how many ways can a committee of 3 be chosen from a group of 12 people? We all know that there are C(12,3) = 220 possibilities (C(N,K) denotes the well-known binomial coefficients). For pure mathematicians, this result may be great. But we want to really generate all the possibilities (via backtracking).

// Example:
// ?- combination(3,[a,b,c,d,e,f],L).
// L = [a,b,c] ;
// L = [a,b,d] ;
// L = [a,b,e] ;
// ...

// P27 (**) Group the elements of a set into disjoint subsets.
// a) In how many ways can a group of 9 people work in 3 disjoint subgroups of 2, 3 and 4 persons? Write a predicate that generates all the possibilities via backtracking.

// Example:
// ?- group3([aldo,beat,carla,david,evi,flip,gary,hugo,ida],G1,G2,G3).
// G1 = [aldo,beat], G2 = [carla,david,evi], G3 = [flip,gary,hugo,ida]
// ...

// b) Generalize the above predicate in a way that we can specify a list of group sizes and the predicate will return a list of groups.

// Example:
// ?- group([aldo,beat,carla,david,evi,flip,gary,hugo,ida],[2,2,5],Gs).
// Gs = [[aldo,beat],[carla,david],[evi,flip,gary,hugo,ida]]
// ...

// Note that we do not want permutations of the group members; i.e. [[aldo,beat],...] is the same solution as [[beat,aldo],...]. However, we make a difference between [[aldo,beat],[carla,david],...] and [[carla,david],[aldo,beat],...].

// You may find more about this combinatorial problem in a good book on discrete mathematics under the term "multinomial coefficients".

// P28 (**) Sorting a list of lists according to length of sublists
// a) We suppose that a list (InList) contains elements that are lists themselves. The objective is to sort the elements of InList according to their length. E.g. short lists first, longer lists later, or vice versa.

// Example:
// ?- lsort([[a,b,c],[d,e],[f,g,h],[d,e],[i,j,k,l],[m,n],[o]],L).
// L = [[o], [d, e], [d, e], [m, n], [a, b, c], [f, g, h], [i, j, k, l]]

// b) Again, we suppose that a list (InList) contains elements that are lists themselves. But this time the objective is to sort the elements of InList according to their length frequency; i.e. in the default, where sorting is done ascendingly, lists with rare lengths are placed first, others with a more frequent length come later.

// Example:
// ?- lfsort([[a,b,c],[d,e],[f,g,h],[d,e],[i,j,k,l],[m,n],[o]],L).
// L = [[i, j, k, l], [o], [a, b, c], [f, g, h], [d, e], [d, e], [m, n]]

// Note that in the above example, the first two lists in the result L have length 4 and 1, both lengths appear just once. The third and forth list have length 3 which appears, there are two list of this length. And finally, the last three lists have length 2. This is the most frequent length.
