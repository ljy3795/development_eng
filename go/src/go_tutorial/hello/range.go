package main

import "fmt"

func main() {
	for _, i := range []int{1, 2, 3, 4, 5, 6, 7} {
		fmt.Println(i)
	}

	sum := 0
	ids := []int{10, 20, 30, 40, 50}
	for i, id := range ids {
		fmt.Println(i, id)

		sum += id
		fmt.Println(sum)
	}

	// range with string indexes and runes
	for i2, c := range "한국어" {
		fmt.Printf("%#U starts at byte position %d\n", c, i2)
	}

	// range with a map
	moons := map[string]string{"Earth": "Moon", "Jupiter": "Europa", "Saturn": "Titan"}
	for k, v := range moons {
		fmt.Printf("%s: %s\n", k, v)
	}

	// range with a map
	numbers := map[string]int{
		"Uno":    1,
		"Dos":    2,
		"Tres":   3,
		"Cuatro": 4,
		"Cinco":  5,
	}

	for k, v := range numbers {
		fmt.Printf("%s: %d\n", k, v)
	}

	// range with a channel
	// iterate over 5 values in the 'queue' channel
	queue := make(chan string, 5)
	queue <- "Enceladus"
	queue <- "Titan"
	queue <- "Europa"
	queue <- "Ganemede"
	queue <- "Io"
	close(queue)

	fmt.Println(queue)

	for q := range queue {
		fmt.Println(q)
	}

	// range with a struct
	type Cities struct {
		name     string
		location [2]int
	}

	// create empty slice of struct pointers
	cities := []*Cities{}

	ct := new(Cities)
	ct.name = "London"
	ct.location[0] = 5
	ct.location[1] = 1
	fmt.Println(ct)
	cities = append(cities, ct)
	fmt.Println(cities)

	ct = new(Cities)
	ct.name = "Sydney"
	ct.location[0] = 34
	ct.location[1] = 51
	fmt.Println(ct)
	cities = append(cities, ct)
	fmt.Println(cities)

	for i := range cities {
		fmt.Println(i)
		c3 := cities[i]
		fmt.Println("City : ", *c3)
	}

	fmt.Println(*cities[0])
}
