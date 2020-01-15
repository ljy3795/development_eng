package main

import "fmt"

func main() {
	// Arrays
	var fruits [2]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"

	fmt.Println(fruits)

	fruits_2 := [2]string{"Apple", "Banana"}
	fmt.Println(fruits_2)

	ids := []int{1, 2, 3, 4, 5}
	fmt.Println(ids)

	// Slices (build on array)
	//  -> []T (unlike array, no specified length)
	//  -> describe a piece of an array
	//  -> length(len) / capacity(cap)

	alphabet := []string{"a", "b", "c", "d", "e"}
	fmt.Println(alphabet)
	fmt.Println(len(alphabet), cap(alphabet))

	// make (type, len, cap)
	s := make([]int, 5, 10)
	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	fruits_slice := []string{"banana", "tomato", "mango", "apple"}
	fmt.Println(fruits_slice)
	fmt.Println(len(fruits_slice), cap(fruits_slice))
	fmt.Println(fruits_slice[1:3])
	fmt.Println(fruits_slice[1:4])
	fmt.Println(fruits_slice[1:])
	fmt.Println(fruits_slice[:])

	// byte slices
	s2 := []byte("abc")
	fmt.Printf("%c\n", s2)
	fmt.Println(s2)

	// append a btye
	s2 = append(s2, byte('d'))
	fmt.Printf("%c\n", s2)
	fmt.Println(s2)

	// first byte
	fmt.Printf("%c\n", s2[0])
	fmt.Println(s2[0])
	fmt.Printf("%c\n", s2[2:])
}
