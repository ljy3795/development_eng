package main

import "fmt"

func sum(s []int, c chan int){
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c channel
}

func main(){
	s := []int{1,2,3,4,5,6,7,8,9,10}
	c := make(chan int)
	go sum(s[:len(s)-1], c)
	go sum(s[len(s)-1:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}