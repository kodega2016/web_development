package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	res, err := divideBy(10, 20)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result:", res)
}

func divideBy(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("cannot divide by zero")
	} else {
		result = x / y
		return result, nil
	}
}
