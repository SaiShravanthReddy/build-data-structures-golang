package main

import "fmt"

// type Value struct {
// 	integer int
// 	float   float32
// 	str     string
// 	boolean bool
// }

// type Array struct {
// 	index int
// 	Value  
// }

type IntArray struct {
	index int
	intValue int 
}

type FloatArray struct {
	index int
	floatValue float32 
}

type StringArray struct {
	index int
	stringValue string 
}

type BoolArray struct {
	index int
	boolValue bool 
}

const (
	integer int
	float   float32
	str     string
	boolean bool
)

// type ArrayType struct {
// 	integer int
// 	float   float32
// 	str     string
// 	boolean bool
// }

func main() {
	// var a Array	
	// fmt.Print(a)

	// arr is array name, type is array type
	create(arr, integer)

	// fmt.Print(get(a, 1))
}

func create(arr string, type ArrayType) {

}

func get(arr Array, ind int) Value {
	// fmt.Print(arr.Value)
	return arr.Value
}

