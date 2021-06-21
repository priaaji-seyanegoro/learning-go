package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreetings(n string) {
	fmt.Printf("Hello Good Morning %v \n", n)
}

func cycleFunction(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circleArea(r float64) float64 {
	return math.Phi * r * r
}

func getInitialName(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	fmt.Print(names)

	var intials []string

	for _, v := range names {
		intials = append(intials, v[:1])
	}

	if len(intials) > 1 {
		return intials[0], intials[1]
	}
	return intials[0], "-"

}

func main() {
	// names := []string{"Aji", "Leo", "Budi"}
	// cycleFunction(names, sayGreetings)

	// areaOne := 2.123
	// areaTwo := 3.323

	// fmt.Printf("Area One is %0.2f and Area Two is %0.2f", circleArea(areaOne), circleArea(areaTwo))

	a, h := getInitialName("aji hirosi")
	fmt.Println(" =>", a, h)

	m1, m2 := getInitialName("Mamad")
	fmt.Println(" =>", m1, m2)
}
