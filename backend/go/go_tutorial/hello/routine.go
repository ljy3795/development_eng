package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello Goroutine")
}

func main() {
	go hello()
	time.Sleep(100*time.Millisecond)
	fmt.Println("Hello")
}