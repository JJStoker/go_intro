package main

import "fmt"

func main() {
	//// Your second function
	// xs := []float64{98, 93, 77, 82, 83}
	// fmt.Println(average(xs))

	//// Variadic functions
	// fmt.Println(add(1, 2, 3))
	// xs := []int{1, 2, 3}
	// fmt.Println(add(xs...))

	//// Closure
	// add := func(x, y int) int {
	// 	return x + y
	// }
	// fmt.Println(add(1, 1))
	// x := 0
	// increment := func() int {
	// 	x++
	// 	return x
	// }
	// fmt.Println(increment())
	// fmt.Println(increment())
	// nextEven := makeEvenGenerator()
	// fmt.Println(nextEven()) // 0
	// fmt.Println(nextEven()) // 2
	// fmt.Println(nextEven()) // 4

	//// Recursion
	// fmt.Println(factorial(10))

	//// defer, panic and recover
	// defer second()
	// first()
	//// panic and recover
	// defer func() {
	// 	str := recover()
	// 	fmt.Println(str)
	// }()
	// panic("PANIC")

	//// Pointers
	// x := 5
	// zero(&x)
	// fmt.Println(x) // x is 0

	//// new
	// xPtr := new(int)
	// one(xPtr)
	// fmt.Println(*xPtr) // x is 1

	fmt.Println(exercise2(16))
	fmt.Println(exercise3(16, 5))
	gen := makeOddGenerator()
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(fib(20))

	x := 1
	y := 2
	exercise11(&x, &y)
	fmt.Println(x, y)
}

func average(xs []float64) float64 {
	total := 0.0
	for _, value := range xs {
		total += value
	}

	return total / float64(len(xs))
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func zero(x *int) {
	*x = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func exercise2(x uint) (uint, bool) {
	return uint(x / 2), bool(x%2 == 0)
}

func exercise3(x ...uint) uint {
	var max uint = 0
	for _, v := range x {
		if v >= max {
			max = v
		}
	}
	return max
}

func makeOddGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		if ret == 0 {
			i++
		} else {
			i += 2
		}
		return
	}
}

func fib(x uint) uint {
	if x == 0 {
		return 0
	}

	if x == 1 {
		return 1
	}

	return fib(x-1) + fib(x-2)
}

func exercise11(x, y *int) {
	x2 := *x
	y2 := *y
	*y = x2
	*x = y2
}

/* Exercises
1: func sum(a []int) uint
2: see exercise2()
3: see exercise3()
4: see makeOddGenerator() / main()
5: see fib()
6:
- defer runs calls after func completion
- panic stops execution
- recover stops the panic and is passed the input of panic
7: & operator
8: * operator; *x = 3
9: new operator; xPtr = new(int) or pass variable with &
10: 2.25
11: see exercise11()
*/
