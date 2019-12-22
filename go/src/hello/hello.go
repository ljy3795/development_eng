package main

import (
	"bytes"
	"fmt"
	"math"
	"github.com/jylee/go_tutorial/string_util"
)

var test string = "aa"

func main() {
	// basic print function
	fmt.Println("Hello New World")

	// set variable with type(could be skipped)
	var x string
	x = "Hello Newww World"

	var year = 2019.1
	fmt.Println(x)
	fmt.Println(year)

	fmt.Printf("type of x : %T\n", x)
	fmt.Printf("type of x : %T\n", year)

	// new variable short assignment (:=)
	// outside function(main()) is not working
	i, j := 1, 2
	fmt.Println(i, j)

	c, python, java := true, false, "no!"
	fmt.Println(c, python, java)
	fmt.Printf("Type of c, python, java is %T %T %T\n", c, python, java)

	a := 1
	a = 2
	fmt.Println(a)

	// Character var
	var myLetter = 'T'
	var anotherLetter byte = 'B'
	fmt.Println(myLetter) // essentially integers
	fmt.Println(anotherLetter)

	my_byte := byte('b')
	my_rune := '~'
	fmt.Printf("%c = %d	%c = %U\n", my_byte, my_byte, my_rune, my_rune)
	// %c : character, %d : base10, %U : Unicode,

	d := true
	fmt.Printf("%t\n", d)

	fmt.Println(c, math.Floor(1.7), math.Ceil(9.2))

	i1, j1 := 0, 0

	fmt.Println(i1, j1)

	var aa string = "STRING"
	fmt.Println(aa)

	bb := []rune(aa)
	fmt.Printf("bb = %c\n", bb)
	fmt.Printf("bb type = %T\n", bb)

	// Bytes
	b1 := byte('a')
	b2 := []byte("AB")          // ()안에있는 string은 각각의 character로 분리
	b3 := []byte{'a', 'b', 'c'} // {}는 각각 char로
	fmt.Printf("b1 = %c\n", b1)
	fmt.Printf("b2 = %c\n", b2)
	fmt.Printf("b3 = %c\n\n", b3)

	s1 := []byte("Hello")
	s2 := []byte("World")
	s3 := [][]byte{s1, s1}
	s4 := bytes.Join(s3, []byte(",")) // Join(s [][]byte, sep []byte) []byte
	fmt.Printf("s1 = %c\n", s1)
	fmt.Printf("s2 = %c\n", s2)
	fmt.Printf("s3 = %c\n", s3)
	fmt.Printf("s4 = %c\n", s4)

	fmt.Printf("%T\n", s1)
	fmt.Printf("%T\n", b1)

	// import custom package
	// when we import use full path, but using use last part of the name (string_util)
	test_str := "reverse me"
	fmt.Println(string_util.Reverse(test_str))

}