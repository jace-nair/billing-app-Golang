package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Helper function that returns the user input along with any error
func getInput(prompt string, r *bufio.Reader) (string, error) { // accepts a prompt and a reader r and returns string and error
	fmt.Print(prompt)
	input, err := r.ReadString('\n')     // reader reads everything before new line or enter
	return strings.TrimSpace(input), err //get rid of any white space and return the string and the error
}

// Create a new bill with user input. Returns a bill object.
func createBill() bill {
	reader := bufio.NewReader(os.Stdin) //create a reader using bufio package's NewReader method and read from Stdin (terminal)

	//fmt.Print("Create a new bill name: ")
	//name, _ := reader.ReadString('\n') // reader reads everything before new line or enter
	//name = strings.TrimSpace(name)     //get rid of any white space

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name) // create a new bill with name passed in
	fmt.Println("Created the bill - ", b.name)
	return b // return the bill
}

func promptOptions(b bill) { //giving user prompt options to add an item, tip or save the file
	reader := bufio.NewReader((os.Stdin)) //local varible
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt { //switch options for user
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		fmt.Println(name, price)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		fmt.Println(tip)
	case "s":
		fmt.Println("You chose s")
	default:
		fmt.Println("that was not a valid option")
		promptOptions(b)
	}
}

func main() {

	//Version 1

	// Create a new instance of the bill object
	//harrysBill := newBill("Harry's bill")

	// Add Items
	//harrysBill.addItem("Onion soup", 4.50)
	//harrysBill.addItem("Veg Pie", 8.95)
	//harrysBill.addItem("toffee pudding", 4.95)
	//harrysBill.addItem("Coffee", 3.25)

	// Update tip
	//harrysBill.updateTip(10)

	//Print a formatted Harry's bill
	//fmt.Println(harrysBill.format())

	mybill := createBill()
	promptOptions(mybill)

}
