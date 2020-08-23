package main

import "fmt"

func main() {
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i, "even")
	// 	} else {
	// 		fmt.Println(i, "odd")
	// 	}
	// }

	fmt.Println("Exercise 2")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("\nExercise 3")
	for i := 1; i <= 100; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

/*
Exercises
1: Small !(10 > 10)
2: see main()
3: see main()

*/
