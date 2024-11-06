package main

import (
	"fmt"
	"os"
)

// Custom Struct Type acts as a blue print for any bill object
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Function that takes a bill name and returns a new bill object
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// Receiver function - receives a unformatted bill and returns a formatted bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	//Loop through items to create a formatted string fs and total price. -25 represents spacing to the right
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	//Add tip
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip)

	//Add total to formatted string fs. -25 represents spacing to the right
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "total:", total+b.tip)

	//Return the formatted fs string
	return fs
}

// Receiver function - function to update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// Receiver function - function to add item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// Receiver funcion - function to save bill
func (b *bill) save() {
	data := []byte(b.format()) //turn fomatted string of the bill into byte slice and store in data variable

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644) //write files to bills folder with billname.txt and 0644 permissions. Returns error if any.
	if err != nil {
		panic(err) // stop the flow of the program and print out an error
	}
	fmt.Println("bill was saved to file")
}
