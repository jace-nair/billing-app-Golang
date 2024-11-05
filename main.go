package main

import "fmt"

func main() {

	// Create a new instance of the bill object
	harrysBill := newBill("Harry's bill")

	//Print a formatted Harry's bill
	fmt.Println(harrysBill.format())
}
