package main

import "fmt"

func main() {

	var myString string = "This is a variable string"
	var myNumber int = 3
	var myFloatNumber float64 = 2.33
	const myConst string = "This is a const string"

	fmt.Println(myString)
	fmt.Println(myNumber)
	fmt.Println(myFloatNumber)
	fmt.Println(myConst)

	myString = "This is a variable string 2"
	myNumber = 6

	fmt.Println(myString)
	fmt.Println(myNumber)

}
