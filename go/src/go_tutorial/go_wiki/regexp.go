package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`a(x*)b(y|z)c`)
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbzc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbzc-")[1])

	re = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$") 
	fmt.Printf("%q\n", re.FindStringSubmatch("/view/testpage"))

}
