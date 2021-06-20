package main

import "fmt"

func main() {
	fmt.Println("ARRAYS AND SLICE LESSONS")

	//ARRAYS
	ages := [3]int{21, 22, 23}
	fmt.Println("arrays => ", ages)
	// change value of array
	ages[2] = 10
	fmt.Println("modify arrays => ", ages)
	//length arrays
	fmt.Println("length of arrays => ", len(ages))

	//SLICE use arrays under the hood
	hobbys := []string{"football", "singing", "dancing"}
	fmt.Println("Slice => ", hobbys)
	// add value of slice acctully not add into prev slice but return new slice(array)
	hobbys = append(hobbys, "drawing")
	fmt.Println("Add New Slice =>", hobbys)
	//length arrays
	fmt.Println("length of slice => ", len(hobbys))

	//Slice - range
	hobbyOne := hobbys[1:3] //its mean get a slice from index 1 to index 3 (not includes latest values)
	fmt.Println("hobbyOne =>", hobbyOne)
	hobbyTwo := hobbys[:3] //its mean get a slice from first index (0) to index 3
	fmt.Println("hobby two =>", hobbyTwo)
	hobbyThree := hobbys[2:] //its mean get a slice from index (2) to lastest index of slice
	fmt.Println("hobby three =>", hobbyThree)
}
