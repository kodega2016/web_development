package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// consume json data to structs
	var rawData string = string(`[
		{
			"username":"kodega2016",
			"email":"khadgalovecoding2016@gmail.com",
			"role":"software developer"
		},
		{
			"username":"rakasidebus",
			"email":"rakasidebus@gmail.com",
			"role":"software developer"
		}
	]`)

	isValid := json.Valid([]byte(rawData))

	if isValid {
		var users []User
		json.Unmarshal([]byte(rawData), &users)
		for _, u := range users {
			u.print()
		}
	}

	// structs to json data
	var userData []User = []User{
		{"kodega", "khadgalovecoding2016@gmail.com", "software developer"},
		{"rakas", "rakasidebus@gmail.com", "frontend developer"},
	}

	bs, err := json.MarshalIndent(userData, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("json data:", string(bs))

}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func (u User) print() {
	fmt.Println("username:", u.Username, "email:", u.Email, "role:", u.Role)
}
