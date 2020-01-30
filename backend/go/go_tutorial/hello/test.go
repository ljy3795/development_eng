package main

import (
	"bytes"
	"fmt"
	"github.com/jylee/go_tutorial/string_util"
	"math"
)



func main() {
	fmt.Println('1')

	x := "Hello"
	fmt.Printf("%T\n", x)

	var i, j int = 1, 2
	fmt.Println(i, j)

	var myLetter = 'T'
	var anotherLetter byte = 'B'
	fmt.Printf("%T\n", myLetter)
	fmt.Printf("%T\n", anotherLetter)

	fmt.Println(myLetter)
	fmt.Println(anotherLetter)

	fmt.Printf("%c\n", myLetter)
	fmt.Printf("%c\n", anotherLetter)

	b1 := byte('a')
	fmt.Println(b1)
	b2 := []byte("AB")
	fmt.Println(b2)
	fmt.Printf("%c\n", b2)
	b3 := []byte{'a', 'b', 'c'}
	fmt.Printf("%c\n", b3)

	s1 := []byte("Hello")
	fmt.Println(s1)
	fmt.Printf("%c\n", s1)
	s2 := []byte("World")
	s3 := [][]byte{s1, s2, s1}
	fmt.Printf("%c\n", s3)
	s4 := bytes.Join(s3, []byte("--"))
	fmt.Printf("%c\n", s4)

	fmt.Printf("%c\n", s4[5])

	s5 := [][]byte{[]byte("Hello"), []byte("New"), []byte("World")}
	fmt.Printf("%c\n", s5)
	s6 := bytes.Join(s5, []byte("--"))
	fmt.Printf("%c\n", s6)

	fmt.Println(math.Floor(3.6))

	str_test := "Heloooo"
	fmt.Println(string_util.Reverse(str_test))

	var fruits [2]string
	fruits[0] = "Apple"
	fruits[1] = "rambutan"

	fmt.Println(fruits)
	fmt.Printf("%d\n",len(fruits))

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	fruits_2 := [2]string{"Apple", "Banana"}
	fmt.Println(fruits_2)

	ids := []int{1,2,3,4,5}
	fmt.Println(ids)

	alp := []string{"AA", "BB", "CC"}
	fmt.Println(alp)
	fmt.Println(len(alp))
	fmt.Println(cap(alp))

	alp2 := make([]string, 3, 5) // length & capacity
	fmt.Println(len(alp2))
	fmt.Println(cap(alp2))

	a1 := []byte("abc")
	a2 := append(a1, byte('d'))
	fmt.Printf("%c\n",a2)
	fmt.Println(string(a2))
	

	range_ids := []int{11,2,3,4,5}
	for _, i := range range_ids {
		fmt.Println(i)
	}

	moons_3 := map[string]string{"Earth": "Moon", "Jupiter": "Europa", "Saturn": "Titan"}
	fmt.Println(moons_3)

	for i, ids := range moons_3 {
		fmt.Println(i, ids)
	}



	t1, t2 := 42, 2701

	p := &t1
	fmt.Println(p)
	fmt.Println(*p)
	*p = 21
	fmt.Println(p)
	fmt.Println(*p)

	fmt.Println(t2)


	p1 := new(Person)
	p1.firstName = "A"
	p1.lastName = "B"
	p1.age = 30
	fmt.Println(p1)

	p2 := Person{lastName : "1", firstName:"2"}
	fmt.Println(p2)
	p2.age++
	fmt.Println(p2)
	p2.hello()

	p2.add_age()
	fmt.Println(p2.age)
}

func (p Person) hello(){
	fmt.Println(p.firstName + " Hello " + p.lastName)
}

func (p *Person) add_age(){
	p.age++
}

type Person struct{
	firstName string
	lastName string
	age int
}