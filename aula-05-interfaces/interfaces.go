package main

/*
Interfaces are collections of method signatures
A type "implements" an interface if it has all of the methods of the given interface defined on it
*/

import (
	"fmt"
	"math"
)

func main() {
	circle := Circle{radius: 8}
	fmt.Print(circle.area())
	checkNumberAbove()
}

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct { //In Go, interfaces are implicit implemented so it is not necessary to use something as "implements {interface name}"
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

/* 	Because the rectangle implements both area() and perimeter() methods, it can be interpreted as a "shape"
because a shape is anything that can have its
area and perimeter calculated
*/

type Circle struct {
	radius float64
}

func (circle Circle) area() float64 {
	result := math.Pi * math.Pow(circle.radius, 2)
	return math.Floor(result)
}

type GlobalInterface interface{} //An empty interface is implemented by all types because it has no requirements

// Multiple interfaces implementation -> a type can implement multiple interfaces as long as it fulfills all the methods of all different interfaces
type Expensive interface {
	cost() float64
}

type Printer interface {
	print()
}

type Email struct {
	isSubscribed bool
	body         string
}

func (email Email) cost() float64 {
	if !email.isSubscribed {
		return 0.05 * float64(len(email.body))
	}
	return 0.01 * float64(len(email.body))
}
func (email Email) print() {
	fmt.Println(email.body)
}

//type switches

func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

func checkNumberAbove() {
	printNumericValue(1)
	printNumericValue("1")
	printNumericValue(struct{}{})
}

/*
	Interfaces are not classes. They're slimmer.
	Interfaces don't have constructors or deconstructors that require that data is created or destroyed.
	Interfaces aren't hierarchical by nature, though there is syntatic sugar to create interfaces that happen to be supersets of other interfaces.
	Interfaces define function signatures, but not underlying behavior.


*/
