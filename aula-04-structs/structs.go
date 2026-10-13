package main

import (
	"fmt"
)

type Car struct { //Structs contains key-value pairs
	Make            string
	Model           string
	FabricationYear int16
	Height          float32
	Width           float32
}

type Truck struct {
	Make            string
	Model           string
	FabricationYear int16
	Height          float32
	Width           float32
	FrontWheel      Wheel //Struct can be concatenated, like FrontWheel is a Wheel struct type
	BackWheel       Wheel
}

type Wheel struct {
	Radius   int
	Material string
}

func invokeCar(car Car) {
	fmt.Printf("Modelo do carro: %v", car.Model)
}

// Anonymous structs
type Bike struct {
	Make            string
	Model           string
	FabricationYear int
	Height          float32
	Width           float32
	Wheel           struct { //Nested anonymous struct within the Bike struct
		Radius   int
		Material string
	}
}

// Embbeded Structs
type car struct {
	make  string
	model string
}
type truck struct {
	// "car" is embbeded, so the definition of a
	// "truck" now also additionally contains all
	// of the fields of the car struct
	car
	bedSize int
}

// struct methods
type rect struct {
	width  int
	height int
}

// area has a receiver of (r rect)
func (r rect) area() int {
	return r.width * r.height
}

var r = rect{
	width:  5,
	height: 2,
}

// Main functon
func main() {
	invokeCar(Car{
		Model: "Nissan \n",
	})
	fmt.Println(r.area())

	truck := Truck{}
	truck.Height = 5
	truck.FrontWheel.Radius = 2
	fmt.Printf("Altura: %v \nDiâmetro da roda dianteira: %v", truck.Height, truck.FrontWheel.Radius)
}
