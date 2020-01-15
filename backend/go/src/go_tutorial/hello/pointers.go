package main

import "fmt"

func main() {
	a := 10 // Assign value to a
	b := &a // Assign memory address(location) where a is stored
	// & operator generates a pointer to its operand

	fmt.Println(a, b)
	fmt.Printf("%T, %T\n", a, b)

	// The * operator denotes the pointer's underlying value
	fmt.Println(*b)

	i := 42
	p := &i

	fmt.Println(p)
	fmt.Println(*p) // read i through the pointer p, this should print 42.
	*p = 21         // set i through the pointer p
	fmt.Println(p)
	fmt.Println(*p) // this should print 21
	fmt.Println(i)  // this should print 21 (i is replaced py *operator)

	i2, j := 42, 2701

	p2 := &i2
	fmt.Println(*p2)
	fmt.Println(p2)
	*p2 = 21
	fmt.Println(*p2)
	fmt.Println(i2)
	fmt.Println(p2)

	p2 = &j
	*p2 = *p2 / 37
	fmt.Println(j)
	fmt.Println(*p2)
	fmt.Println(i)

}
