package chapter

import (
	"fmt"
)

// MethodsAndInterfaces Chapter
func MethodsAndInterfaces() {
	pure := motorcycle{year: "2017", model: "R9T", make: "BMW"}
	fmt.Println(pure.year)

	// Look, it's like a method!
	pure.print()

	pure.renameModel("R9T Pure")
	pure.print()

	// printable type interface thing
	printIt(pure)
}

type motorcycle struct {
	year  string
	make  string
	model string
}

func (moto motorcycle) print() {
	fmt.Println(moto.year + " " + moto.make + " " + moto.model)
}

func (moto *motorcycle) renameModel(modelName string) {
	moto.model = modelName
}

type printable interface {
	print()
}

func printIt(something printable) {
	something.print()
}
