package main

import "fmt"

func main() {
	// arrays()
	// slices()
	// maps()
	exercise4()
}

func arrays() {
	// Arrays
	x := [4]float64{
		98,
		77,
		85,
		89,
	}

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
func slices() {
	// Slices
	x1 := make([]float64, 5, 6)
	fmt.Println(x1, append(x1, 6))
	x2 := append(x1, 6, 7, 8, 9, 10, 11, 12, 13, 14)
	fmt.Println(x2)

	arr := []float64{1, 2, 3, 4, 5}
	x3 := arr[0:5]
	fmt.Println(x3)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice3, slice4)
	fmt.Println(slice3, slice4)
}

func maps() {
	// Maps
	x1 := make(map[string]int)
	x1["key"] = 10
	fmt.Println(x1["key"])

	x2 := make(map[int]int)
	x2[1] = 10
	delete(x2, 1)
	fmt.Println(x2)

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
	fmt.Println(elements["Li"])
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
}

func exercise4() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	min := 100
	for _, v := range x {
		if v <= min {
			min = v
		}
	}
	fmt.Println(min)
}

/* Exercises
1: x[3]
2: 3, with capacity of 9. Above 9 it will create a new array and return a new slice.
3: c, d, e, f
4: see exercise4()
*/
