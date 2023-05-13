package main

import (
	"fmt"
	"time"
)

func access(ch chan int) {
	time.Sleep(time.Second)
	fmt.Println("start accessing channel")

	for i := range ch {
		fmt.Println(i)
		time.Sleep(time.Second)
	}

	fmt.Println("end accessing channel")
}

func main() {
	// only modify this line to defined the capacity
	ch := make(chan int, 3)
	defer close(ch)

	go access(ch)

	for i := 0; i < 9; i++ {
		ch <- i
		fmt.Println("Filled")
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Done")
}
