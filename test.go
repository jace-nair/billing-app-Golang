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

	var names []string = []string{"mario", "luigi", "yoshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for i, v := range names {
		fmt.Println(i, v)
	}

	for _, v := range names {
		fmt.Println(v)
	}

	// Invoking sayGoodMorning()
	sayGoodMorning("Ava")

	// Invoking sayGoodNight()
	sayGoodNight("Alex")

	// Invoking sayGoodMorningAll()
	sayGoodMorningAll(names)

	// Invoking namesAndGreetings()
	namesAndGreetings(names, sayGoodNight)
}
