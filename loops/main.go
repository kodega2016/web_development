package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("total sum:", sum)

	total := 2

	for total < 100 {
		total += total
	}

	fmt.Println("total:", total)
}
