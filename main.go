package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOption(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill , t - add tip) : ", reader)
	switch opt {
	case "a":
		fmt.Println("You Choose item")
		name, _ := getInput("Item Name :", reader)
		price, _ := getInput("Item Price :", reader)
		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Price Must Be Number")
			promptOption(b)
		}
		b.addItem(name, priceFloat)
		fmt.Println("Item Added - ", name, price)
		promptOption(b)
	case "t":
		tip, _ := getInput("Enter tip amount (Rp): ", reader)
		tipFloat, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Tip Must Be Number")
			promptOption(b)
		}
		b.updateTip(tipFloat)
		fmt.Println("Your Tip $", tip)
		promptOption(b)
	case "s":
		b.save()
		fmt.Print("You name file saved - ", b.name)
	default:
		fmt.Println("Wrong Input ...")
		promptOption(b)
	}
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create new bill name : ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func main() {
	myBill := createBill()
	promptOption(myBill)
	// fmt.Println(myBill)
}
