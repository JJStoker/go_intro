package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type MultiShape struct {
	shapes []Shape
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r Rectangle) perimeter() float64 {
	// Gespiekt
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) perimeter() float64 {
	// Gespiekt
	return 2 * math.Pi * c.r
}

func (m MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func main() {
	a := new(Android)
	a.Name = "Test"
	a.Talk()

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	c := Circle{0, 0, 5}
	fmt.Println(c.area())
	fmt.Println(totalArea(&c, &r))

	multiShape := MultiShape{
		shapes: []Shape{
			Circle{0, 22, 4},
			Rectangle{033, 0, 10, 110},
			MultiShape{
				shapes: []Shape{
					Circle{1, 22, 1},
					Rectangle{033, 0, 10, 10},
				},
			},
		},
	}

	fmt.Println(multiShape.area())

	// Exercise 3
	fmt.Println(r.perimeter())
	fmt.Println(c.perimeter())
}

/* Exercises
1: methods use receiver instances; it lets you use the . operator in stead of defining arguments
2: it lets you inherit properties, flattening the struct
2a: Lets you use methods directly on the containing type
3: see main()

*/
