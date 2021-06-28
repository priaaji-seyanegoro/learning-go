package main

import "fmt"

func sumAll(number ...int) int {
	//parameter berubah menjadi slice
	total := 0
	for _, value := range number {
		total += value
	}

	return total
}

func main() {
	// parameter pd sumAll terakhir bs menerima banyak parameter dgn 1 argumen
	total := sumAll(10, 10, 10, 10, 10, 10)
	fmt.Println(total)

	//slice as parameters
	numbers := []int{20, 20, 20, 20}
	total2 := sumAll(numbers...)
	fmt.Println(total2)
}
