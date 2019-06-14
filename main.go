package main

import (
	"fmt"
	"github.com/flatten"
	"math/rand"
	"time"
)

// Type declaration
type StringSlice []string
type EncodedPair struct {
	c int
	s string
}

func main() {

	//Welcome message
	fmt.Printf("This is a golang solution for 99 Problems of Prolog \n")

	// Variable declarations
	my_list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	palli := []int{1, 2, 3, 2, 1}
	nested := map[string]interface{}{
		"a": "b",
		"c": map[string]interface{}{
			"d": "e",
			"f": "g",
		},
		"z": 1.4567,
	}
	my_chaos := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 13, 2, 4, 3, 3, 3, 3, 3, 3, 3, 5, 5, 5, 55, 32, 12, 12, 12, 12, 12, 12, 12, 12, 12, 34, 34, 54, 45, 34, 34, 45, 45, 34, 34, 34, 45, 45, 45, 45, 45, 45}
	my_chaos_string := []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "e", "r", "f", "z"}
	my_string_simple := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"}
	my_workers := []string{"aldo", "beat", "carla", "david", "evi", "flip", "gary", "hugo", "ida"}
	my_sublisted := []StringSlice{{"a", "b", "c"}, {"d", "e"}, {"f", "g", "h"}, {"d", "e"}, {"i", "j", "k", "l"}, {"m", "n"}, {"o"}}

	// Testing the functions

	//Lists
	fmt.Println("List Functions:")
	//1. Last element of the list
	fmt.Println("P01: Last element of the list")
	//Test
	last := my_last(my_list)
	fmt.Printf("%v \n", last)

	//2. Zweiletztes element
	fmt.Println("P02: Zweiletztes element")
	//Test
	zweiletztes := my_zweiletztes(my_list)
	fmt.Printf("%v\n", zweiletztes)

	//3. K'th element
	fmt.Println("P03: K'th element")
	//Test
	kth := element_at(4, my_list)
	fmt.Printf("4th is:%v \n", kth)

	//4. Number of elements of a list
	fmt.Println("P04: Number of elements")
	//Test
	num := num_elements(my_list)
	fmt.Printf("%v elements\n", num)

	//5. Reversing a list
	fmt.Println("P05: Reversing a list")
	//Test
	my_rlist := rev_list(my_list)
	fmt.Println(my_rlist)

	//6. Palindrome
	fmt.Println("P06: Palindrome")
	//Test
	id_palindrome(palli)

	//7. Flatten a nested map structure
	fmt.Println("P07: Flatten a nested map structure")
	//Test
	flat, _ := flatten.Flatten(nested, "", flatten.DotStyle)
	fmt.Println(flat)

	//flat, err := flatten.Flatten(nested, "", flatten.RailsStyle)
	//fmt.Println(flat)
	//fmt.Println(err)

	//8. Eliminate consecutive duplicates of list elements
	fmt.Println("P08: Eliminate consecutive duplicates of list elements.")
	//Test
	clean_list := el_consec(my_chaos)
	fmt.Println(my_chaos)
	fmt.Println(clean_list)

	//9. Pack Consicutive
	fmt.Println("P09: Pack Consicutive")
	//Test
	pack := Pack(my_chaos_string)
	fmt.Println(pack)

	//10. Run-length encoding of a list
	fmt.Println("P10: Run-length encoding of a list")
	//Test
	compre := encode(my_chaos_string)
	fmt.Println(compre)

	//11. Modified run-length encoding
	fmt.Println("P11: Modified Run-length encoding of a list")
	//Test
	recomp := mod_encode(my_chaos_string)
	fmt.Println(recomp)

	//TODO #2 : Test the first written function after debbug
	//12. Decode a run-length encoded list
	fmt.Println("P12: Decode a run-length encoded list")
	//Test
	decoded := Decode(recomp)
	fmt.Println(decoded)

	//14. Duplicate the elements of a list
	fmt.Println("P14: Element duplication of a list")
	//Test
	duble := dupli(my_string_simple)
	fmt.Println(duble)

	//15. Duplicate the elements of a list for a given number
	fmt.Println("P15: Duplicate the elements of a list")
	//Test
	freewill := dupli_freewill(my_string_simple, 4)
	fmt.Println(freewill)

	//16. Drop every N'th element from a list
	fmt.Println("P16: Drop every N'th element from a list")
	//Test
	dropped := drop(my_string_simple, 3)
	fmt.Println(dropped)

	//17. Split a list into two parts; the length of the first part is given.
	fmt.Println("P17: Split a list into two parts with the length of the first part")
	//Test
	First, Second := split(my_string_simple, 3)
	fmt.Println(First, Second)

	//18. Extract a slice from a list.
	fmt.Println("P18: Extract a slice from a list.")
	//Test
	sliced := slice(my_string_simple, 1, 5)
	fmt.Println(sliced)

	//19. Rotate the list as n items
	fmt.Println("P19: Rotate the list n times to left")
	//Test
	rotated := rotate_L(my_string_simple, 4)
	fmt.Println(rotated)

	//20. Remove the K'th element from a list
	fmt.Println("P20: Remove the K'th element from a list")
	//Test
	x, newl := remove_at(my_string_simple, 5)
	fmt.Println(x)
	fmt.Println(newl)

	//21. Insert an element at a given position into a list
	fmt.Println("P21: Insert an element at a given position into a list")
	//Test
	yeni := insert_at(my_string_simple, 3, "n")
	fmt.Println(yeni)

	//22. Create a list containing all integers within a given range
	fmt.Println("P22: Create a list containing all integers within a given range")
	//Test
	my_int_list := range_int(4, 19)
	fmt.Println(my_int_list)

	//23. Extract a given number of randomly selected elements from a list
	fmt.Println("P23: Extract a given number of randomly selected elements from a list")
	//Test
	rnd_ext := rnd_select(my_string_simple, 3)
	fmt.Println(rnd_ext)

	//24. Lotto: Draw N different random numbers from the set 1..m
	fmt.Println("P24: Lotto")
	//Test
	results := lotto()
	fmt.Println(results)

	//25. Generate a random permutation of the elements of a list
	fmt.Println("P25: Generate a random permutation of the elements of a list")
	//Test
	L := rnd_permu(my_string_simple)
	fmt.Println(L)

	//26. Generate the combinations of K distinct objects chosen from the N elements of a list
	fmt.Println("P26: Generate the combination of K distinct objects chosen from the N elements of a list")
	//Test
	combo := combination(my_string_simple, 3)
	fmt.Println(combo)

	//27. Group the elements of a set into disjoint subsets
	fmt.Println("P27: Group the element of a set into disjoint subsets")
	//Test
	G1, G2, G3 := group3(my_workers, 2, 2, 5)
	fmt.Println(G1)
	fmt.Println(G2)
	fmt.Println(G3)

	//28. Sorting a list of lists according to length of sublists
	fmt.Println("P28: Sorting a list of lists according to length of sublists")
	//Test
	len_sorted := lsort(my_sublisted)
	fmt.Println(len_sorted)

	//Arithmetic Functions:
	fmt.Println("Arithmetic functions:")

	//P31: Determine whether a given integer number is prime.
	fmt.Println("P31: Determine whether a given integer number is prime")
	//Test
	fmt.Println(is_prime(7))
	fmt.Println(is_prime(2))
	fmt.Println(is_prime(15))
	fmt.Println(is_prime(0))
	fmt.Println(is_prime(151))
	fmt.Println(is_prime(113))
	fmt.Println(is_prime(1))

	//P32: Determine the greatest common divisor of two positive integer numbers
	fmt.Println("P32: Determine the greatest common divisor of two positive integer numbers")
	//Test
	fmt.Println(gcd(63, 36))

	//P33: Determine whether two positive integer numbers are comprime
	fmt.Println("P33: Determine two positive integer numbers are coprime")
	//Test
	fmt.Println(coprime(35, 64))

	//P34: Calculate Euler's totient fucntion phi(m)
	fmt.Println("P34: Calculate Euler's totient function phi(m)")
	//Test
	fmt.Println(phi(10))

	//P35: Determine the prime factors of a given positive integer
	fmt.Println("P35: Determine the prime factors of a given positive integer")
	//Test
	fmt.Println(prime_factors(315))

	//P36: Determine the prime factors of a given positive integer
	fmt.Println("P36: Determine the prime factors of a given positive intteger(2)")
	//Test
	fmt.Println(prime_factors_mult(prime_factors(315)))

	//P37: Calculate Euler's totient function phi(m)(improved)
	fmt.Println("Calculate Euler's toient function phi(m) (improved)")
	//Test

}

// Function declarations
func my_last(x []int)(y int){ // Problem 1
	fmt.Printf("%v \n",x[len(x)-1])
	y = x[len(x)-1]
	return y
}

func my_zweiletztes(x []int)(y int){// Problem 2
	fmt.Printf("%v \n",x[len(x)-2])
	y = x[len(x)-2]
	return y
}

func element_at(k int, x []int)(y int){// Problem 3
	if (k<len(x))||(k!=len(x)) {
		y = x[k-1]
		fmt.Printf("%v\n", y)
	} else{
		fmt.Println("You are out of bounds for the given list!")
	}
	return y
}

func element_at_s(k int, x StringSlice)(y StringSlice){// Problem 3.20
	for count,item := range x{
		if (count+1) == k {
			y=append(y,item)
		}
	}
	return y
}

func num_elements(x []int)(y int){// Problem 4
	y = len(x)
	fmt.Println(y)
	return y
}

func rev_list(x []int)(y []int){// Problem 5
	for range x {
		y=append(y,0)
	}
	l:=len(x)
	for i:=l;i>0;i--{
		y[l-i] = x[i-1]
	}
		fmt.Println(y)
	return y
}

func id_palindrome(x []int){// Problem 6
	y:=rev_list(x)
	for i:=0;i<len(x);i++{
		if x[i] != y[i]{
			fmt.Println("This list is not a palindrome.")
			return
		}else if i == len(x)-1{
			fmt.Println("This list is a palindrome")
		}
	}
}

/*func Flatten(nested []interface{}) (result []string){// Problem 7
	for _, s := range nested {
		switch i := s.(type) {
		case string:
			result = append(result, i)
		case []interface{}:
			result = append(result, Flatten(i)...)
		}
	}
	return result
}*/




func el_consec(x []int)(y []int){// Problem 8
	for i :=0; i<len(x); i++ {
		if i == 0{
			y=append(y,x[i])
		}
		for j:=0; j<len(y);j++ {
			if x[i] == y[j]{
				break
			}else if j == len(y)-1{
				y=append(y,x[i])
			}
		}

	}
	return y
}

func Pack (s StringSlice) (result []StringSlice) { // Problem 9 from cheat sheet
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

func encode(x StringSlice) (y []EncodedPair) {// Problem 10
	var compressed []StringSlice
	compressed = Pack(x)
	for _, item := range compressed {
		y = append(y, item.EncodePair())
	}
	return y
}

func (s StringSlice) EncodePair() EncodedPair {
	var char string
	var count int
	char = s[0]
	for i := 0; i < len(s); i++ {
		count++
	}
	return EncodedPair{count, char}
}

func mod_encode(x StringSlice)(y []interface{}){ // Problem 11
	var compressed []StringSlice
	compressed = Pack(x)
	for _, item := range compressed{
		if len(item) != 1 {
			y = append(y, item.EncodePair())
		}else{
			y=append(y,item[0])
		}
	}
	return y
}

//TODO #1: Find the Difference between the two functions and the reason why we can't range the interface with the first one
/*func dencode(x interface{})(y StringSlice){ // Problem 12
	for _,item := range x{
		switch i := item.(type) {
		case string:
			y = append(y, i)
		case EncodedPair:
			for j := 0; j < i.c; j++ {
				y = append(y, i.s)
			}
		}
	}
	return y
}*/

func Decode(g []interface{}) StringSlice {// Problem 12 cheat sheet
	var result StringSlice
	for _, item := range g {
		switch i := item.(type) {
		case string:
			result = append(result, i)
		case EncodedPair:
			for j := 0; j < i.c; j++ {
				result = append(result, i.s)
			}
		}
	}
	return result
}

func dupli(x StringSlice)(y StringSlice){// Problem 14
	var backup StringSlice
	backup = x
	for _,item := range backup{
		for j:=0;j<2;j++{
			y=append(y,item)
		}
	}
	return y
}

func dupli_freewill(x StringSlice, n int)(y StringSlice){// Problem 15
	var backup StringSlice
	backup = x
	for _,item := range backup{
		for j:=0;j<n;j++{
			y=append(y,item)
		}
	}
	return y
}

func drop(x StringSlice,n int)(y StringSlice){// Problem 16
	for count,item := range x{
		if (count+1)%n != 0{
			y=append(y,item)
		}
	}
	return y
}

func split(x StringSlice, n int )(y1 StringSlice, y2 StringSlice){// Problem 17
	for count,item := range x{
		if count<n {
			y1=append(y1,item)
		}else{
			y2=append(y2,item)
		}
	}
	return y1, y2
}

func slice(x StringSlice, min int, max int)(y StringSlice){// Problem 18
	for count, item := range x{
		if (count+1)>min && (count+1)<max{
			y=append(y,item)
		}
	}
	return y
}

func rotate_L(x StringSlice, n int)(y StringSlice){// Problem 19
	x1,x2 := split(x,n)
	for _,item := range x2{
		y=append(y,item)
	}
	for _,item := range x1{
		y=append(y,item)
	}
	return y
}

func remove_at(x StringSlice,n int)(k StringSlice, y StringSlice){// Problem 20
	y=drop(x,n)
	k=element_at_s(n,x)
	return y,k
}

func insert_at(x StringSlice, n int,s string)(y StringSlice){// Problem 21
	y = append(x[:n-1], append([]string{s}, x[n-1:]...)...)
	return y
}

func range_int(x int, y int)(z []int){// Problem 22
	for i:=x;i<y+1;i++{
		z=append(z,i)
	}
	return z
}

func rnd_select(x StringSlice, n int)(y StringSlice){ // Problem 23
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		y = append(y, x[r.Intn(len(x))])
	}
	return y
}

func lotto()(y []int){ // Problem 24
	var x []int
	x=range_int(1,49)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0; i<6; i++{
		y=append(y,x[r.Intn(len(x))])
	}
	return y
}

func rnd_permu(x StringSlice)(y StringSlice){ // Problem 25
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var perm []int
	perm = r.Perm(len(x))
	for _, index := range perm {
		y = append(y, x[index])
	}
	return y
}

func combination(x StringSlice, n int)(y StringSlice){ // Problem 26
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<n;i++{
		y=append(y,x[r.Intn(len(x))])
	}
	return y
}

func group3(x StringSlice,k int, l int, n int)(G1 StringSlice, G2 StringSlice, G3 StringSlice){// Problem 27
	G1,G2 = split(x,k)
	G2,G3 = split(G2,l)
	return G1, G2, G3
}

func lsort(x []StringSlice)(y []StringSlice){ // Problem 28

	var tok StringSlice
	for _, item := range x {
		y = append(y, item)
		if len(y) != 1 {
			for i := 0; i < len(y); i++ {
				for j:= 0; j < len(y); j++{
					if (len(y[i])<len(y[j]))&&(i>j) {
						tok = y[i]
						y[i]=y[j]
						y[j]=tok
					}
					if (len(y[i])>len(y[j]))&&(j>i){
						tok = y[i]
						y[i]=y[j]
						y[j]=tok
					}
				}
			}
		}
	}
	return y
}

func is_prime(n int)(token bool) {// Problem 31
	if n == 1 || n == 0 || n == 2{
		return false
	} else {
		for i := 2; i < n+1; i++ {

				if (n%i == 0) && (i != n){
					return false
				}

		}
		return true
	}
}

func gcd(n int, m int)(y int) {// Problem 32
	com := 1
	k := n
	l := m

	for k != 1 || l != 1 {
		for i := 0; i < m; i++ {
			if is_prime(i) || i == 2 {
				if k%i == 0 && l%i == 0 {
					com = com * i
					k = k / i
					l = l / i
					break
				} else if k%i == 0 && l%i != 0 {
					k= k / i
					break
				} else if k%i != 0 && l%i == 0 {
					l = l / i
					break
				} else if k == 1 || l == 1 {
					return com
				}
			}
		}
	}
	return com
}

func coprime(n int, m int)(result bool){// Problem 33
	if gcd(n,m) == 1{
		return true
	}else{
		return false
	}
}

func phi(m int)(y int){// Problem 34
	for i:=1;i<m;i++{
		if coprime(i,m){
			y++
		}
	}
	return y
}

func prime_factors(n int)(y []int){// Problem 35
	for n != 1{
		for i:=0;i<n+1;i++{
			if is_prime(i){
				if n%i == 0 {
					y=append(y,i)
					n=n/i
					break
				}
			}
		}
	}
return
}

func prime_factors_mult(x []int)(y [][]int){// Problem 36

	var k [][]int
	for _,item := range x{
		tok := 0
		for i:=0;i<len(x);i++{
			if item == x[i]{
				tok++
			}
		}
		k=append(k,[]int{item, tok})
	}
	y=append(y,k[0])
	for j:=1;j<len(k);j++{
		for i:=0;i<len(k);i++{
			if k[j][0] != y[i][0]{
				y = append(y,k[j])
				break
			}else{break}
		}
	}
	return
}