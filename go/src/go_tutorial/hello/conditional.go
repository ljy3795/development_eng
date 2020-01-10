package main

import "fmt"

func main() {
	a, b := 1, 0
	if a < b {
		fmt.Printf("%d is less then %d\n", a, b)
	} else if a == b {
		fmt.Printf("%d is equal to %d\n", a, b)
	} else {
		fmt.Printf("%d is greater than %d\n", a, b)
	}

	galaxy := "M87"
	switch galaxy {
	case "Milky way":
		fmt.Printf("Galaxy name is '%s' 1\n", galaxy)
	case "Andromeda":
		fmt.Printf("Galaxy name is '%s' 2\n", galaxy)
	case "M87":
		fmt.Printf("Galaxy name is '%s' 3\n", galaxy)
	}
}
