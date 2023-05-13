package main

func forgottenSender(ch chan int) {
	data := 3

	// This will NOT block
	ch <- data
}

func handler() {
	// Declare a BUFFERED channel
	ch := make(chan int, 1)

	go forgottenSender(ch)
}

func main() {
	handler()
}
