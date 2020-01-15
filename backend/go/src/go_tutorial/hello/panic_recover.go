package main

import (
	"fmt"
)

func foo() {
	fmt.Println("foo()")
	defer bar()
	panic("Argh!")
}

func bar() {
	fmt.Println("bar()")
	err := recover()
	fmt.Printf("err = %+v\n", err)
}

func main(){
	foo()
	fmt.Println()

	f()
}


func f() {
	defer func() {
		if r := recover(); r != nil{
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int){
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v",i)) // panic
	}

	defer fmt.Println("Defer in g",i)
	fmt.Println("Printing in g", i)
	g(i+1)
}