package main

import "fmt"

type bill struct {
	name string
	item map[string]float64
	tip  float64
}

//make new bills
func newBill(name string) bill {
	b := bill{
		name: name,
		item: map[string]float64{},
		tip:  0,
	}
	return b
}

//receiver func is only trigger with associate fun with it
func (b *bill) format() string {
	formatString := "Bill breakdown \n"
	var total float64 = 0
	//list of items
	for k, v := range b.item {
		formatString += fmt.Sprintf("%-15v  : Rp.%0.2f \n", k, v)
		total += v
	}
	total += b.tip
	//tip
	formatString += fmt.Sprintf("%-16v : Rp.%0.2f \n", "Tip", b.tip)
	//total
	formatString += fmt.Sprintf("%-16v : Rp.%0.2f \n", "Total", total)
	return formatString
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip += tip
}

// add an item
func (b *bill) addItem(name string, price float64) {
	b.item[name] = price
}
