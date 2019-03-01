package main

import (
	"fmt"
	"github.com/flatten"
)
type StringSlice []string
type EncodedPair struct {
	c int
	s string
}
func main() {

	//Welcome message
	fmt.Printf("This is a golang solution for 99 Problems of Prolog \n")

	// Variable declarations
	my_list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10 ,11, 12}
	palli := []int{1,2,3,2,1}
	nested := map[string]interface{}{
		"a": "b",
		"c": map[string]interface{}{
			"d": "e",
			"f": "g",
		},
		"z": 1.4567,
	}
	my_chaos := []int{1,1,1,1,1,1,1,1,1,1,13,2,4,3,3,3,3,3,3,3,5,5,5,55,32,12,12,12,12,12,12,12,12,12,34,34,54,45,34,34,45,45,34,34,34,45,45,45,45,45,45}
	my_chaos_string := []string{"a","a","a","a","a","a","a","a","a","a","a","a","a","a","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","c","c","c","c","c","c","c","c","c","c","c","c","c","c","c","c","c","c","c","c","c","c","d","d","d","d","d","d","d","d","d","d","d","d","d","d","d","d","d","d","d","d","d","d","d","d","d","d","d","d","d","d","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","e","r","f","z"}
	my_string_simple := []string{"a","b","c","d","e","f"}

	// Testing the functions

	//1. Last element of the list
	fmt.Println("P01: Last element of the list")
	//Test
	last := my_last(my_list)
	fmt.Printf("%v \n",last)

	//2. Zweiletztes element
	fmt.Println("P02: Zweiletztes element")
	//Test
	zweiletztes := my_zweiletztes(my_list)
	fmt.Printf("%v\n", zweiletztes)

	//3. K'th element
	fmt.Println("P03: K'th element")
	//Test
	kth := element_at(4,my_list)
	fmt.Printf("4th is:%v \n",kth)

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
	compre :=encode(my_chaos_string)
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
	freewill := dupli_freewill(my_string_simple,4)
	fmt.Println(freewill)

	//16. Drop every N'th element from a list
	fmt.Println("P16: Drop every N'th element from a list")
	//Test
	dropped := drop(my_string_simple,3)
	fmt.Println(dropped)

	//17. Split a list into two parts; the length of the first part is given.
	fmt.Println("P17: Split a list into two parts with the length of the first part")
	//Test
	First, Second := split(my_string_simple,3)
	fmt.Println(First,Second)

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

