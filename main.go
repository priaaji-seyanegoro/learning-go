package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Your Blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

func main() {
	// you can declare to variable
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", blacklist)

	//you can write function  on function
	registerUser("Eko", func(name string) bool {
		return name == "admin"
	})
}
