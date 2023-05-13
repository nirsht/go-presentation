package main

import (
	"errors"
	"fmt"
	"time"
)

func networkCall() int {
	time.Sleep(3 * time.Second)
	return 1
}

func continueToValidateOtherData() error {
	return errors.New("data is invalid! Returning")
}

func forgottenSender(ch chan int) {
	data := networkCall()

	ch <- data
}

func handler() error {
	ch := make(chan int)
	go forgottenSender(ch)

	err := continueToValidateOtherData()
	if err != nil {
		return errors.New("data is invalid! Returning")
	}

	data := <-ch
	fmt.Println(data)
	return nil
}

func main() {
	_ = handler()
}
