package main

import (
	"fmt"
	"time"
)

func abandonedWorker(ch chan string) {
	for data := range ch {
		fmt.Println(data)
	}

	fmt.Println("Worker is done, shutting down")
}

func handler(inputData []string) {
	ch := make(chan string, len(inputData))
	defer close(ch) // this is the solution

	for _, data := range inputData {
		ch <- data
	}

	go abandonedWorker(ch)

}

func main() {
	inputData := []string{"one", "two", "three"}
	handler(inputData)
	time.Sleep(3 * time.Second)
}
