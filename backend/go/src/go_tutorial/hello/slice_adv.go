package main

import "fmt"

func main() {
	var array = [6]int{1,2,3,4,5,6}
	var slices = array[1:4]

	fmt.Printf("slice : cap %v, len %v, %v\n", cap(slices), len(slices), slices)
	fmt.Printf("array : cap %v, len %v, %v\n", cap(array), len(array), array)

	// make -> func make([]T, len, cap) []T
	slice := make([]int, 4, 10)
	fmt.Printf("slice : cap %v, len %v, %v\n", cap(slice), len(slice), slice)
	slice = make([]int, 4)
	fmt.Printf("slice : cap %v, len %v, %v\n", cap(slice), len(slice), slice)
}