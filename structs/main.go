package main

import "fmt"

func main() {
	user := User{
		FirstName:     "Khadga",
		LastName:      "Shrestha",
		Email:         "khadgalovecoding2016@gmail.com",
		ContactNumber: "9811013594",
		Address: Address{
			city:    "Kathmandu",
			street:  "Lainchaur",
			country: "Nepal",
		},
	}

	user.print()

}

type User struct {
	FirstName     string
	LastName      string
	Email         string
	ContactNumber string
	Address       Address
}

type Address struct {
	city    string
	street  string
	country string
}

func (u User) print() {
	fmt.Println(u.FirstName, u.LastName, u.Email, u.ContactNumber, u.Address.city)
}
