package main

import "fmt"

func main() {

	var myString string = "This is a variable string"
	var myNumber int = 3
	var myFloatNumber float64 = 2.33
	const myConst string = "This is a const string"

	var myArray [3]int = [3]int{20, 25, 30}
	var anotherArray [2]int = [2]int{4, 5}

	var mySlice []int = []int{1, 2, 3}
	var mySlice2 = []int{1, 2, 3, 4, 5}
	mySlice3 := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(myString)
	fmt.Println(myNumber)
	fmt.Println(myFloatNumber)
	fmt.Println(myConst)

	myString = "This is a variable string 2"
	myNumber = 6

	fmt.Println(myString)
	fmt.Println(myNumber)

	fmt.Printf("%v and %v. And %v is an integer while %v is a float\n", myString, myConst, myNumber, myFloatNumber)

	fmt.Println(myArray, len(myArray), anotherArray, len(anotherArray))

}
