package main

import "fmt"

func main() {
	firstInts, secondInts := splitAnySlice([]int{0, 1, 2, 3})
	fmt.Println(firstInts, secondInts)

	firstStrings, secondStrings := splitAnySlice([]string{"japão", "jamaica", "jalapão", "jaraguá", "guiné"})
	fmt.Println(firstStrings, secondStrings)

}

//	Put simply, generics allows us to use variables to refer to specific types.
//	This is an amazing feature because it allows us to write abstract functions that drastically reduce code duplication

func splitIntSlice(s []int) ([]int, []int) { // If it was to write another function that splits a string slice, it'd be needed to write everything again
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func splitAnySlice[T any](s []T) ([]T, []T) { // "T" is the type paramether of the function, which is "any". Then, "s" receives an empty "T" slice. And then, the function returns two empty "T" slices
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

//	A reason to use generics is to "dry up" the code, with "dry" standing for "do not repeat yourself".
//	It's a way of writing code that can be reused. So instead of writing two nearly equal functions, we write a type-safe generic function once and it can be used throughout all the application.

//	CONSTRAINTS
//	Constraints allow us to create generics that are only useful for a specific set of types

type Stringer interface {
	String() string
}

func concat[T Stringer](vals []T) string {
	//	Differently from the other functions, here, the type "T" isn't able to be aything. It has to be a Stringer.
	//	In this situation, a Stringer is anything that has the "String()" function and returns a string
	result := ""
	for _, val := range vals {
		//	This is where the "String()" method is used.
		//	That's why we need a more specific constraint instead of the any constraint
		result += val.String()
	}
	return result
}

//	INTERFACE TYPE LISTS
//	When generics were released, a new way of writing interfaces was also released at the same time.
//	We can now simply list a bunch of types to get a new interface/constraint

//	Ordered is a type constraint that matches any ordered type.
//	An ordered type is one that supports the <, <=, > and >= operators.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~string
}

//	PARAMETRIC CONSTRAINTS
//	Your interface definitions, which can later be used as constraints, can accept type parameters as well.

//	The "store" interface represents a store that sells products.
//	It takes a type paremeter P that represents the type of the product
type store[P product] interface {
	Sell(P)
}

type product interface {
	Price() float64
	Name() string
}

type book struct {
	title  string
	author string
	price  float64
}

func (b book) Price() float64 {
	return b.price
}

func (b book) Name() string {
	return b.title
}
