package main

import (
	"os"
	"fmt"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	// main 마지막에 파일 close 실행 (python에서 finally로 보면 됨)
	defer f.Close()

	// 파일 읽기
	bytes := make([]byte, 1024)
	f.Read(bytes)
	fmt.Println(len(bytes))
	// fmt.Printf(string(bytes))

	for i:= 0; i < 5; i++ {
		// LIFO
		defer fmt.Println(i)
	}
	fmt.Println("main")
}