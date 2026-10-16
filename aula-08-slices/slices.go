package main

import "fmt"

func main() {
	Array()
	slicesAndArrays()
	createMatrix(10, 10)
	iterateFruits()
}

//	[2, 3, 1] -> A conventional array, with "2" being the index 0, "3" being the index 1 and "1" being the index 2
//	In Go, arrays have a fixed size, so the array above would be [3] (its size) and int (because it only stores integers)
// Something that wasn't mentioned in the "functions" class is that it is possible to store functions within variables, just like in JavaScript

var array [10]int //	A new array called "array" of 10 integers

var initializedArray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //Initialized literal

func Array() [3]string {
	return [3]string{"A", "B", "C"}
}

//	Now, Slices.
//	In Go, Slices are declared in a very similar form to Arrays
//	To declare a Slice, it's just necessary to -> [](type)
//	The size is missing, which means that, while Arrays have fixed sizes, Slices have dynamic sizes
//	Slices are created on top of Arrays.

//	Example

var a = [5]int{6, 3, 2, 6, 5} //Array of 5 integers
var b = a[1:4]                //	"b" is a Slice of "a", going from index 1 and stop before index 4 (The first number is included and the second one is excluded)

func slicesAndArrays() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[0:2]
	fmt.Print(b)
}

//	Make -> Let a slice be created without creating an array under the hood
//	Slices created with "make" will be filled with the zero value of the type

//	func make([]T, len, cap) []T
var mySlice = make([]int, 5, 10)

//	The capacity argument is usually omitted and defaults to the length
var mySlice2 = make([]int, 5)

//	Slices literals
var mySliceLiteral = []string{"Eu", "Amo", "Go"} //	To be a Slice, it cannot have a "3" inside the brackets, otherwise it'd be an Array

//	Length -> the length of a Slice is simply the number of elements it contains. built-in len() function

//	Capacity -> the capacity of a Slice is the number of elements in the underlying Array. built-in cap() function

//	Variadic functions
//	When a function can take an arbitrary number of final arguments, so it's used "..." in the function signature

func sum(nums ...int) {
	//	nums is just a slice
	for i := 0; i < 10; i++ {
		num := nums[i]
		fmt.Print(num)
	}
}

//	Spread Operator -> Allows passing a Slice into a variadic function
func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func callFunctionAbove() {
	names := []string{"Bob", "Sue", "Alice"} // Slice of values
	printStrings(names...)                   // Passing the slice into the variadic function "printStrings"
}

//	Append -> the built-in variadic "append()" function is used to dynamically add elements to a slice
//	It's first parameter is the slice that will have an addition and the second one is the variable that will be appended

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

//	TRICKY SLICES

// 	The "append()" function changes the underlying Array of its paremeter AND returns a new Slice.
//	This means that using "append()" on anything other than itself is usually a BAD idea
//	Do not do this -> someSlice = append(otherSlice, element)
//	<----------------------------------------------->

//	RANGE
//	Go provides syntatic sugar to iterate easily over elements of a Slice
//	for INDEX, ELEMENT := range SLICE{}

func iterateFruits() {
	fruits := []string{"banana", "apple", "grape"}
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
}
