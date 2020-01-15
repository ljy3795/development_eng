package main

import "fmt"

type point struct {
	x, y, z int
}

func main() {
	p := point{1,2,3}
	fmt.Println(p)

	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)

	var b bool
	b = true
	fmt.Printf("%t\n", b)

	fmt.Printf("%s\n", "quoted")
	fmt.Printf("%s\n", "\"quoted\"")
	fmt.Printf("%q\n", "\"quoted\"")
}