package main

import (
	"fmt"
)

func sayGoodMorning(name string) {
	fmt.Printf("Good Morning %v\n", name)
}

func sayGoodNight(name string) {
	fmt.Printf("Good Night %v\n", name)
}

func sayGoodMorningAll(names []string) {
	for _, v := range names {
		fmt.Printf("Good Morning %v\n", v)
	}
}

func namesAndGreetings(names []string, greetings func(string)) {
	for _, v := range names {
		greetings(v)
	}
}

func main() {

	//Slice
	var names []string = []string{"mario", "luigi", "yoshi", "peach"}

	//Map
	var menu map[string]float64 = map[string]float64{"soup": 4.99, "pie": 7.99, "salad": 6.99, "toffee pudding": 3.55}

	//Looping through a Slice
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for i, v := range names {
		fmt.Println(i, v)
	}

	for _, v := range names {
		fmt.Println(v)
	}

	//Looping through a Map
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// Invoking sayGoodMorning()
	sayGoodMorning("Ava")

	// Invoking sayGoodNight()
	sayGoodNight("Alex")

	// Invoking sayGoodMorningAll()
	sayGoodMorningAll(names)

	// Invoking namesAndGreetings()
	namesAndGreetings(names, sayGoodNight)

	// Printing Slice
	fmt.Println(names)

	// Printing Map
	fmt.Println(menu)
}
