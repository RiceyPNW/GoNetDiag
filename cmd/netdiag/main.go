package main

import "fmt"

func main() {
	printMe()
}

func printMe() {
	fmt.Println("Hello World!")
	fmt.Println(addition(8, 29))

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"John": 23, "Casper": 24}
	fmt.Println(myMap2["John"])

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}
}

func addition(a int, b int) int {
	var result int = a + b
	return result
}
