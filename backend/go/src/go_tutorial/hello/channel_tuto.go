package main

import (
	"fmt"
	"time"
)

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


	// ----------------------------------------- //
	// Un-Buffer
	// c := make(chan int)
	// c <- 1 // 수신루틴이 없으므로 데드락
	// fmt.Println(<-c)

	// buffered channel
	ch = make(chan int, 10)
	ch <- 1 
	fmt.Println(<-ch)


	// channel params.
	ch_2 := make(chan string, 1)
	sendChan(ch_2)
	receiveChan(ch_2)

	// channel close
	ch_3 := make(chan int, 2)

	ch_3 <- 1
	ch_3 <- 2

	// 채널닫기
	close(ch_3)

	// 채널수신
	println(<-ch_3)
	println(<-ch_3)

	if _, success := <-ch_3; !success {
		println("No remaining data")
	}

	ch_3 = make(chan int, 2)
	ch_3 <- 7
	ch_3 <- 7

	close(ch_3)

	// println(<-ch_3)
	// println(<-ch_3)

	for i := range ch_3 {
		println(i)
	}


	done_1 := make(chan bool)
	done_2 := make(chan bool)

	// go routine을 이용한 비동기적 실행
	go run_1(done_1)
	go run_2(done_1)

// EXIT:
// 	for { //for loop를 돌면서 done_1이 완료되면 print, 그리고 다시 for loop를 실행해서
// 		select {
// 		case <-done_1:
// 			println("Run_1 완료")
// 			// close(done_1)
// 		case <- done_2:
// 			println("Run_2 완료")
// 			close(done_1)
// 			close(done_2)
// 			break EXIT // EXIT로 이동. 하지만 해당 for loop는 pass
// 		}
// 	}
// 	// fmt.Println("Start")

EXIT:
    for {
        select {
        case <-done_1:
            println("run1 완료")
 
        case <-done_2:
            println("run2 완료")
            break EXIT
        }
    } 
}

// channel param.
func sendChan(ch chan<- string) {
	ch <- "Data"
	// x := <- ch // 에러발생
}

func receiveChan(ch <-chan string) {
	data := <- ch
	fmt.Println(data)
}



func run_1(done chan bool) {
	time.Sleep(1*time.Second)
	done <- true
}


func run_2(done chan bool) {
	time.Sleep(2*time.Second)
	done <- true
}