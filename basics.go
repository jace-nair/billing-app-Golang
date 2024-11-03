package main

import (
	"fmt"
	"sort"
)

func basicfunction() {

	var myString string = "This is a variable string"

	const myConst string = "This is a const string"

	var myNumber int = 3

	var myFloatNumber float64 = 2.33

	var myArray [3]int = [3]int{20, 25, 30}

	var anotherArray [2]int = [2]int{4, 5}

	//var oneMoreArray [3]uint = [3]uint{6, 7, 8}

	//var mySlice []int = []int{1, 2, 3}
	//var mySlice2 = []int{1, 2, 3, 4, 5}
	//mySlice3 := []int{1, 2, 3, 4, 5, 6}

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

	var ages []int = []int{2, 3, 5, 4, 6}

	fmt.Println(sort.SearchInts(ages, 4))

	var names []string = []string{"one", "two", "three", "four", "five"}

	fmt.Println(names)

	//var x int = 0

	//for x < 5 {
	//fmt.Println(x)
	//x++
	//}

	for x := 0; x < 5; x++ {
		fmt.Printf("The value of x is %v\n", x)
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	var age int = 35

	fmt.Println(age > 40)

	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age > 30 && age < 40 {
		fmt.Println("Age is between 30 and 40")
	} else {
		fmt.Println("Age is greater than 40")
	}

}
