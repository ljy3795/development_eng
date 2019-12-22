package main

import (
	"fmt"
	"strconv" // string conversion
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// method of the struct
// func (value receiver) func_name(parameters) return_type {code}
func (p Person) hello() string {
	return "Hello, I am" + p.firstName + " " + p.lastName + ", " +
		strconv.Itoa(p.age) + " years old"
}

// method (pointer receiver) - this modifies data
// to keep track of the struct, need to use pointer receiver
func (p *Person) hasBirthday() {
	p.age++
}

// value receiver
func (p Person) hasBirthday_2() {
	p.age++
	fmt.Println(p.age)
}

// Stringer interface
func (p Person) String() string {
	return fmt.Sprintf("%v, (%d)", p.firstName, p.age)
}

// declare IPAddr type
type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (a IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", a[0], a[1], a[2], a[3])
}

func main() {
	// initialize person using struct
	p1 := Person{firstName: "Steven", lastName: "King", city: "Chicago", gender: "m", age: 23}
	p2 := Person{"Neena", "Kochhar", "Boston", "f", 13}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.firstName, p2.firstName)
	p2.age++
	fmt.Println(p2.age)
	fmt.Printf("%T\n", p2)

	// struct pointer
	p3 := new(Person)
	p3.firstName = "Juyeong"
	p3.lastName = "Lee"
	p3.city = "Seoul"
	p3.gender = "Male"
	p3.age = 30
	fmt.Println(p3)

	fmt.Println(p3.hello())

	fmt.Println(p2.hello())
	fmt.Println(&p2)
	p2.hasBirthday()
	fmt.Println(p2.hello())
	fmt.Println(&p2)
	p2.hasBirthday()
	fmt.Println(p2.hello())
	fmt.Println(&p2)

	fmt.Println(p2.hello())
	p2.hasBirthday_2()
	fmt.Println(p2.hello())
	p2.hasBirthday_2()
	fmt.Println(p2.hello())

	my_str := p2.String()
	fmt.Println(p2)
	fmt.Println(my_str)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {1, 2, 3, 4},
	}

	fmt.Println(hosts)
	fmt.Printf("%T\n", hosts["loopback"])
	fmt.Printf("%T\n", hosts)

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
