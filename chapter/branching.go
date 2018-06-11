package chapter

import "fmt"

// Branching Chapter
func Branching() {
	if true {
		fmt.Println("it's true!")
	}

	if localScope := "only inside block"; true {
		fmt.Println(localScope)
	}

	// switch example with matching
	switchExample("a")
	switchExample("aaaa")

	// switch example with booleans
	booleanSwitchExample("kk")
	booleanSwitchExample("sdjkfs")

	typeSwitchExample(2)
}

func switchExample(thing string) {
	var val string
	switch thing {
	case "a":
		val = "It's a"
		fallthrough
	case "b":
		val = "It's a, jk, it's b, bc fallthrough"
	default:
		val = "idk"
	}
	fmt.Println(val)
}

func booleanSwitchExample(thing string) {
	var val string
	switch {
	case len(thing) == 1:
		val = "one"
	case len(thing) == 2:
		val = "two"
	default:
		val = "I can't count that high"
	}
	fmt.Println("len is " + val)
}

// interface is kind of like any
func typeSwitchExample(thing interface{}) {
	var val string
	switch thing.(type) {
	case int:
		val = "an int"
	case string:
		val = "a string"
	default:
		val = "idk what that is"
	}
	fmt.Println(val)
}
