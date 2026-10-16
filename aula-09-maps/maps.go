package main

import "fmt"

//	Just like in other programming languages, a Map is just a way to store key:value elements
//	The syntax is: make(map[T]T) or map[T]T{}
func main() {
	var maps = make(map[string]int)
	maps["John"] = 37
	maps["Mary"] = 24
	maps["Mary"] = 21 //	Overwrites 24

	var ages = map[string]int{
		"John": 1,
		"Mary": 2,
	}
	fmt.Println(len(ages))
	fmt.Println(len(maps))

	//	The "len()" function also works on Maps, returning the total number of key-value pairs

	//	MUTATIONS

	//	Insert an element
	//	m[key] = elem

	//	Get an element
	//	elem = m[key]

	//	Delete an element
	//	delete(m, key)

	//	Check if an element exists
	//	elem, ok := m[key] -> if the "key" is in "m", then "ok" is true. If not, "ok" is false

	// It is not possible to use Slices, Maps or Functions as Map keys

	//	hits := make(map[string]map[string]int) //	A nested string-integer Map inside a String map -> similar to a JSON

	// 	n := hits["/doc"]["au"] //	Acessing the "au" index in the second Map through the "/doc" index in the first Map
	//	This will probably become an unwieldy approach when adding data, because it's not possible to be sure that the inner map exists

	type Key struct {
		Path, Country string
	}

	hits := make(map[Key]int)
	hits[Key{"/", "vn"}]++

	//	When attempting to get a value from a map where the key doesn't exist, it returns the zero value
	//	If the Map itself doesn't exist, the code will panick

	//	NESTED MAPS

	var nestedMap1 = make(map[string]map[string]int)
	var nestedMap2 = make(map[rune]map[string]int)
	var nestedMap3 = make(map[int]map[string]map[string]int)

	fmt.Println(nestedMap1)
	fmt.Println(nestedMap2)
	fmt.Println(nestedMap3)
}
