package main

import "fmt"

func main() {
	var name string = "Khadga Bahadur Shrestha"
	var age int = 25
	fmt.Println(name, age)

	username, role, salary := "kodega2016", "software developer", 4500.0
	fmt.Println(username, role, salary)

	var (
		email    string = "khadgalovecoding2016@gmail.com"
		address  string = "Madhumalla Morang Nepal"
		isActive bool   = true
	)

	fmt.Println(email, address, isActive)
}
