package chapter

import (
	"fmt"
)

// Slices Chapter
func Slices() {
	fmt.Println("You can't change the size of an array after you declare it")
	fmt.Println("Arrays type is its size and it's underlying type")
	fmt.Println("Arrays are initialized with 0 values")
	fmt.Println("Arrays are not pointers to memory locations, like some other languages")

	fmt.Println("Slices add features to arrays")
	fmt.Println("Sizes do affect the type of Slices")
	fmt.Println("A Slice is basically a pointer to an array")
	fmt.Println("Slices can grow, they are variable length")
	fmt.Println("Slices are passed by reference")

	var sliceOfIntegers []int
	sliceOfIntegers = make([]int, 3) // 3 is initial length
	sliceOfIntegers[0] = 1
	sliceOfIntegers[1] = 2
	sliceOfIntegers[2] = 3
	// sliceOfIntegers[3] = 4 this throws an error

	// easier intialization of slices
	anotherSlice := []string{"green", "puppy", "hat", "beans", "box"}
	fmt.Println(anotherSlice[2])

	// slicing slices
	newSlice := anotherSlice[1:3]
	printRange(newSlice)
	newSlice = anotherSlice[1:]
	printRange(newSlice)

	// append to slices
	newSlice = append(newSlice, "cool")
	printRange(newSlice)
	newSlice = append(newSlice, newSlice...)
	printRange(newSlice)

	// to delete you must actually use append..
	sliceDeleteAt(newSlice, 1)
	printRange(newSlice)
}

func printRange(slice []string) {
	for _, n := range slice {
		fmt.Print(n + " ")
	}
	fmt.Println("")
}

func sliceDeleteAt(slice []string, index int) {
	afterIndex := index + 1
	slice = append(slice[:index], slice[afterIndex:]...)
}
