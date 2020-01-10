package main

import "fmt"

func say_hello(msg string) {
	fmt.Println(msg)
}

// regular function
// returning "func(string)" - anonymous function taking a string
func regular_f_returning_anonymous_f() func(string) {
	// returns an anonymous function which is an inner function
	return func(msg string) {
		fmt.Println(msg)
	}
	// func(){}() -> func(){} -> we only declare the anonymous function without calling it
}

// return a function which is returning int
// 익명함수는 내부가 아닌 외부(int_seq)의 변수 i를 참조하고 있음. 이에 따라 외부 변수 i가 상태를 계속 유지하는 방식
func int_seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// recursive
func fibo(n int) int {
	if n <= 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}

// iterative
func fibo_ano() func() int {
	x, y := 0, 1
	return func() int {
		r := x
		x, y = y, x+y
		return r
	}
}

func main() {
	// regular function
	say_hello("Hello")

	// anonymous function (w/o func name)
	func(msg string) {
		fmt.Println(msg)
	}("Hello from an anonymous function")

	// () means call and execute
	print_fnc := regular_f_returning_anonymous_f()
	print_fnc("Hello from returned anonymous function")

	next_int := int_seq() // keeping track of the values of i even though it's gone out of function body
	fmt.Println(next_int())
	fmt.Println(next_int())
	fmt.Println(next_int())
	// A CLOSURE is a function value (next_int()) that references variable (i) from outside of its body(int_seq()). But still it remembers the value
	// function is "bound" to the variable

	next_int2 := int_seq() // keeping track of the values of i even though it's gone out of function body
	fmt.Println(next_int2())
	fmt.Println(next_int2())
	fmt.Println(next_int2())

	// Fibonacci
	n := 10
	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", fibo(i))
	}
	fmt.Println("")
	next_fibo := fibo_ano()
	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", next_fibo())
	}
	fmt.Println("")

	next_fibo = fibo_ano()
	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", next_fibo())
	}
}
