package main

import "fmt"

func main() {

	// Create a new instance of the bill object
	harrysBill := newBill("Harry's bill")

	// Add Items
	harrysBill.addItem("Onion soup", 4.50)
	harrysBill.addItem("Veg Pie", 8.95)
	harrysBill.addItem("toffee pudding", 4.95)
	harrysBill.addItem("Coffee", 3.25)

	// Update tip
	harrysBill.updateTip(10)

	//Print a formatted Harry's bill
	fmt.Println(harrysBill.format())
}
