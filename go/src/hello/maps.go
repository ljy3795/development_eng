package main

import "fmt"

//[(])
func isBalanced(p string) bool {
	// define map
	moons := make(map[string]string)
	moons["Earth"] = "Moon"
	moons["Jupiter"] = "Europa"
	fmt.Println(moons)

	pairs := map[rune]rune{'(': ')', '{': '}', '[': ']'}
	stack := []rune{}

	for _, c := range p {
		// fmt.Printf("1_%c\n", c)
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else {
			if len(stack) < 1 {
				return false
			}
			// fmt.Printf("2_%c\n", stack)
			// fmt.Printf("3_%c\n", stack[len(stack)-1])
			// fmt.Printf("4_%c\n", pairs[stack[len(stack)-1]])
			// fmt.Println("")
			if c == pairs[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
				// fmt.Printf("5_%c\n", stack)
			} else {
				return false
			}
		}
	}

	return true
}

func main() {
	// To initialize map, make() function

	// define map
	moons := make(map[string]string)

	// Assign
	moons["Earth"] = "Moon"
	moons["Jupiter"] = "Europa"
	moons["Saturn"] = "Titan"

	fmt.Println(moons)

	moons_2 := []string{}

	moons_2 = append(moons_2, "Moon")
	moons_2 = append(moons_2, "Europa")
	moons_2 = append(moons_2, "Titan")

	fmt.Println(moons_2)

	moons_3 := map[string]string{"Earth": "Moon", "Jupiter": "Europa", "Saturn": "Titan"}
	fmt.Println(moons_3)

	paraenthese := []string{"()", ")(", "[(])", "(()){}[]", "())[]{}", "()[]{}(([])){[()][]}"}
	for _, p := range paraenthese {
		fmt.Printf("%s: %t\n", p, isBalanced(p))
	}
}
