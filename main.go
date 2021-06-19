package main

import "fmt"

func main() {
	// variables
	myName := "aji"
	myAge := 22
	//print
	fmt.Println("Hello First Golang Wish Me Luck")
	fmt.Print("Spacing Line \n")
	fmt.Println("Hello My Name is:", myName, "And My Age is: ", myAge)

	//Formatting String - printf(formatted string)
	fmt.Printf("My name is %v and My Age %v \n", myName, myAge) //v is variable
	fmt.Printf("My name is %q and My Age %q \n", myName, myAge) // q is quote (just read string)
	fmt.Printf("Type Variable Of Age => %T \n", myAge)          // T is Typeof Variable
	fmt.Printf("Your Score is => %0.1f", 225.55)                // f is formatter float ||  %<decimal formatter>f
}
