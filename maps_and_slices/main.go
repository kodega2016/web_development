package main

import (
	"fmt"
	"strings"
)

func main() {
	// maps
	var userInfo map[string]string
	userInfo = make(map[string]string)

	// assigning value to the map
	userInfo["name"] = "khadga bahadur shrestha"
	userInfo["role"] = "software developer"
	userInfo["address"] = "madhumalla morang"
	userInfo["password"] = "password"

	fmt.Println(userInfo)

	// deleting the value from the map
	delete(userInfo, "password")

	//accessing the value from the map
	name, isOk := userInfo["name"]
	if isOk {
		fmt.Println("name:", name)
	}

	// updating the value of the map
	userInfo["name"] = "Arjun shrestha"
	fmt.Println(userInfo)

	// accessing the key-value pair of the map
	for key, value := range userInfo {
		fmt.Printf("%v:%v\n", key, value)
	}

	fmt.Println("total number of key-value pair:", len(userInfo))

	// slices
	var users []string = []string{
		"khadga", "sakar", "sabin", "rakesh",
	}

	fmt.Println("total users:", len(users))
	fmt.Println("first user and last user:", users[0], users[len(users)-1])
	fmt.Println("users:", strings.Join(users, ","))

	// add new user to the users
	users = append(users, "Asok", "Aashis")
	fmt.Println("users:", users)
}
