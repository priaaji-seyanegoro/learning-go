package main

import (
	"fmt"
	"sort"
	"strings"
)

// https://golang.org/pkg/ - this list of standart package on golang

func main() {
	greeting := "Hello Friends!!"
	fmt.Println(strings.Contains(greeting, "hello"))          // methods to check variable contains something word - true/false
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hai")) // methods for replace word
	fmt.Println(strings.ToUpper(greeting))                    // uppercase word
	fmt.Println(strings.Split(greeting, " "))                 // convert word to splice / arrays

	sortHello := []string{"H", "E", "L", "L", "O"}
	sort.Strings(sortHello)
	fmt.Println("HELLO =>", sortHello)
}
