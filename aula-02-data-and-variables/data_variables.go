package main

import "fmt"

//import "fmt"

func main() {
	// uint - Unsigned Integer -> Sets automatically uint32 or uint64
	// uint8  - Unsigned Integer -> Only positive numbers
	// uint16 - Unsigned Integer -> Only positive numbers
	// uint32 - Unsigned Integer -> Only positive numbers
	// uint64 - Unsigned Integer  -> Only positive numbers

	// int - Integer -> Sets automatically int32 or int64
	// int8 - Integer -> Both positive and negative numbers
	// int16 - Integer -> Both positive and negative numbers
	// int32 - Integer -> Both positive and negative numbers
	// int64 - Integer -> Both positive and negative numbers

	// float - Float -> Decimal points numbers
	// float32 - Float -> Decimal points numbers
	// float64 - Float -> Decimal points numbers

	//byte - Byte -> Equals to int8

	//rune - Rune -> Equals to int32

	//bool - Boolean -> True or false

	//string - String -> Stores text with double quotation marks, single quotation marks only stores one char

	//nil -> Undefined/Null -> None, Non-existing, Null, etc

	var name string = "Daniel" //Var -> can change later
	const pi float32 = 3.14    //Const -> Immutable
	fmt.Printf("Olá, seu nome é %v \n", name)
	fmt.Println(pi)

	//Implicit type assignment operator
	y := 3
	//Type cast in implicit assignment operators
	x := uint32(3)
	fmt.Println(y)
	fmt.Printf("%T", y)
	fmt.Println(x)
	fmt.Printf("%T \n", x)

	if y < 1 {
		fmt.Printf("O número %v é menor que 1", y)
	} else if y > 5 {
		fmt.Printf("O número %v é maior que 5", y)
	} else {
		fmt.Println("Outra condição")
	}
}
