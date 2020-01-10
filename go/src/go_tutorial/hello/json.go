package main

import (
	"fmt"
	"encoding/json"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := Message{"Interfaces", "Empty Interfaces", 155468231}
	fmt.Println(m)

	b, err := json.Marshal(m)
	fmt.Println(string(b))
	fmt.Println(err)

	var m2 Message // create a place where the decoded data will be stored
	err = json.Unmarshal(b, &m2)
	fmt.Println(m)
}