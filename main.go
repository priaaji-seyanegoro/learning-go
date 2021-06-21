package main

import "fmt"

func main() {
	fmt.Println("Lopping lesson")

	names := []string{"Aji", "Agus", "Hafiz"}

	// for i := 1; i < len(names); i++ {
	// 	fmt.Printf("The value of %v at index %v \n", names[i], i)
	// }

	for index, value := range names {
		fmt.Printf("The value of %v at index %v \n", value, index)
	}
}
