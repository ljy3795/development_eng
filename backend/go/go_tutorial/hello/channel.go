package main

import (
	"fmt"
)

// func hello() {
// 	fmt.Println("Say Hello")
// }

func hello(finished chan bool){
	fmt.Println("Hello GoRoutine")
	finished <- true
}

// iterative using closure
func fibo_1() func() int {
	x, y := 0, 1
	return func() int {
		r := x
		x, y = y, x+y
		return r
	}
}

// iterative using channel
func fibo_2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibo_select(c, quit chan int) {

}

func main() {
	// create int channel
	ch := make(chan int)

	go func(){
		ch <- 123
	}()

	var i int
	i = <- ch // 채널로부터 123을 받는다
	println(i)
	println("")

	// Go rountine이 끝날 때 까지 기다리는 기능
	// -> Go routine에서 어떤 작업이 실행되고 있을 때, 메인루틴은 <-done에서 계속 수신하면서 대기
	done := make(chan bool)
	go func(){
		for i := 0; i < 5; i++{
			fmt.Println(i)
		}
		done <- true
	}()
	<- done


	var my_int_chan chan int
	my_string_chan := make(chan string)
	fmt.Printf("The type of my_int_chan is '%T'\n", my_int_chan)
	fmt.Printf("The type of my_string_chan is '%T'\n", my_string_chan)

	finished := make(chan bool)
	fmt.Printf("'%v'\n", finished)
	go hello(finished)
	<- finished // blocking (receive data from finished channel)
	fmt.Println("main")

	fibo_1 := fibo_1()
	fmt.Println(fibo_1())
	fmt.Println(fibo_1())
	fmt.Println(fibo_1())
	fmt.Println(fibo_1())
	fmt.Println(fibo_1())

	c := make(chan int, 10)
	go fibo_2(10, c)
	for i := range c {
		fmt.Printf("%d ",i)
	}
	fmt.Println("")


	// c = make(chan int)
	// q := make(chan int)
	// go func() {
	// 	fmt.Println("Goroutine started")
	// 	for i := 0; i < 5; i++ {
	// 		value := <- c
	// 		fmt.Printf("received %d\n", value)
	// 	}
	// 	q <- 999
	// }()

	// for i := range c {
	// 	fmt.Printf("%d ",i)
	// }
	// fmt.Println("")
}
