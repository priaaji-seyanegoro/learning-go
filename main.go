package main

import "fmt"

// map is type data on golang like obj on javascript but must have key and value same tipe data

func main() {
	menus := map[string]float64{
		"soup":    6.33,
		"chikens": 10.00,
		"egg":     2.20,
	}
	fmt.Println(menus)
	fmt.Println(menus["soup"])

	for k, v := range menus {
		fmt.Println(k, "-", v)
	}
}
