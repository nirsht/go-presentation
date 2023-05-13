package main

import "fmt"

func someFunction(ch chan int) {
	data := <-ch
	fmt.Println(data)
}

func main() {
	ch := make(chan int)

	go someFunction(ch)
}
