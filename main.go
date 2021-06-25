package main

import "fmt"

func main() {
	makajiBill := newBill("Makajikenduri")
	makajiBill.updateTip(1000)
	makajiBill.addItem("Pisang Ijo", 10000)
	makajiBill.addItem("Gorengan", 5000)

	fmt.Println(makajiBill.format())
}
