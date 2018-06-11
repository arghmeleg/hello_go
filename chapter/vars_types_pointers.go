package chapter

import "fmt"

type salutation struct { // capital S makes this visible, lowercase would make it private
	name     string
	greeting string
}

// First is a constant
const First = "first const"

const (
	aa = 1
	bb = 2
)

const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

const (
	d = iota // 0
	e        // 1
	f        // 2
)

// VarsTypesPointers Chapter
func VarsTypesPointers() {
	fmt.Println("Hello!")

	var stevesVar string
	stevesVar = "Hello Go World!"
	fmt.Println(stevesVar)

	var a, b, c int
	fmt.Println(a, b, c)

	var d, e, f int = 1, 2, 3
	fmt.Println(d, e, f)

	var g, h, i = 3, false, "ok"
	fmt.Println(g, h, i)

	j := ":= short syntax does initiallization and assignment (only works inside of a function)"
	fmt.Println(j)

	var message = &j      // & gets memory location, message is a pointer
	fmt.Println(&j)       // memory location
	fmt.Println(message)  // also a memory location
	fmt.Println(*message) // value at that memory location

	*message = "new value!"
	fmt.Println(j)

	var s = salutation{name: "steve", greeting: "yeehaw"}
	s = salutation{"steve", "yeehaw"} // can also initialize like this
	fmt.Println(s)
	s.name = "bob"
	fmt.Println(s.name, s.greeting)
}

// # Types
// ## Basic Types
// bool
// string
// int, int8, int16, int32, int64
// uint, uint8, uint16, uint32, uint64, uintptr
// byte (uint8)
// rune (int32), like a char
// float32, float64
// complex64, complex128
//
// ## Other Types
// Array
// Slice - Like array, vector or list. Allowed to grow.
// Struct
// Pointer
// Function
// Interface - Template for defining methods.
// Map
// Channel
