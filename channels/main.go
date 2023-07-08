package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int)
	defer close(c)

	go calculateValue(c)
	fmt.Println("the value from the channel:", <-c)
	fmt.Println("total go routines:", runtime.NumGoroutine())

}

func calculateValue(c chan int) {
	c <- 10
}
