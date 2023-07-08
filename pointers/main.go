package main

import "fmt"

func main() {
	var p *int
	x := 120
	p = &x

	fmt.Printf("The value of the address %v is %v\n", p, *p)
	changeValue(p, 10)
	fmt.Printf("The new value for the address %v is %v\n", p, *p)
}

func changeValue(s *int, newValue int) {
	*s = newValue
}
