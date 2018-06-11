package chapter

import (
	"fmt"
	"strconv"
)

// Functions Chapter
func Functions() {
	fmt.Println("Hello from functions")

	var pure = Motorcycle{"BMW", 2017}
	Rev(pure)

	// variable length example
	fmt.Println(VariadicFunc(1, 2, 3, 4, 5) + " things")

	// passing a function example
	takesAfunc("potato", makeExcite)

	// closure examples
	megaExcite := createAppender("!!!!")
	megaExcite("whoa")
	confuse := createAppender("????")
	confuse("uhhh")
}

// Motorcycle stores some basic stuff about Motorcycles
type Motorcycle struct {
	make string
	year int
}

func motoMessage(year int, make string) (string, string) { // second string is retrun type
	return "(" + strconv.Itoa(year) + ") " + make, "It's a " + make
}

func motoMessage2(year int, make string) (first string, second string) { // second string is retrun type
	first = "[" + strconv.Itoa(year) + "] " + make
	second = "Das " + make
	return
}

// Rev prints some info about a Motorcycle
func Rev(moto Motorcycle) {
	greet1, greet2 := motoMessage(moto.year, moto.make)
	fmt.Println(greet1 + " | " + greet2)

	greet3, greet4 := motoMessage2(moto.year, moto.make)
	fmt.Println(greet3 + " | " + greet4)
}

// VariadicFunc (very addict) functions - variable number of parameters in a functions
func VariadicFunc(things ...int) string {
	return strconv.Itoa(len(things))
}

// you can pass functions around
func makeExcite(something string) {
	fmt.Println(something + "!")
}

func takesAfunc(something string, ok func(string)) {
	fmt.Print("Make excite? ")
	ok(something)
}

// Printer is something that modifies and prints stuff
type Printer func(string)

func createAppender(toAppend string) Printer {
	return func(mainString string) {
		fmt.Println(mainString + toAppend)
	}
}
