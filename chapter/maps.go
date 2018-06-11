package chapter

import "fmt"

// Maps Chapter
func Maps() {
	fmt.Println("keys have to have the equality operator defined")
	fmt.Println("maps are like pointers, if you pass them into a function, and mutate them, they'll be mutated outside of the function")
	fmt.Println("maps are not thread safe")

	var simpleMap map[string]string // map with key of string and value of string
	simpleMap = make(map[string]string)
	simpleMap["steve"] = "programmer"
	simpleMap["bob"] = "mechanic"
	fmt.Println("steve is a " + simpleMap["steve"])
	simpleMap["steve"] = "awesome " + simpleMap["steve"]
	fmt.Println("steve is an " + simpleMap["steve"])

	easyMap := map[string]string{
		"jelly": "dog",
		"steve": "human",
		"thor":  "god",
	}
	for name, what := range easyMap {
		fmt.Println(name + " is a " + what)
	}

	purgeFakes(easyMap)
	checkIfExists(easyMap, "steve")
	checkIfExists(easyMap, "thor")
}

func purgeFakes(myMap map[string]string) {
	delete(myMap, "thor")
}

func checkIfExists(myMap map[string]string, name string) {
	if value, exists := myMap[name]; exists {
		fmt.Println("this " + value + " exists!")
	} else {
		fmt.Println(name + " does NOT exist")
	}
}
