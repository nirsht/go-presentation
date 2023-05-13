package main

func forgottenSender(ch chan int) {
	data := 3

	// This is blocked as no one is receiving the data
	ch <- data
}

func main() {
	ch := make(chan int)

	go forgottenSender(ch)
}
