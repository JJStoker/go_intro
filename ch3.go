package main

import "fmt"

func main() {
	// Exercise 5
	fmt.Print("Enter fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	celsius := (input - 32) * 5 / 9

	fmt.Println(celsius)

	// Exercise 6
	fmt.Print("Enter feet: ")
	fmt.Scanf("%f", &input)

	metres := (input * 0.3048)

	fmt.Println(metres)
}

/*
Exercises

1: var string x = "test" / x := "test"
2: 6
3: determines where the variable can be used
4: const can't be assigned a new value after first assignment
5: see main()
6: see main()

*/
