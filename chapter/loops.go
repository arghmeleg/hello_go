package chapter

import (
	"fmt"
	"strconv"
)

// Loops Chapter
func Loops() {
	for i := 0; i < 4; i++ {
		fmt.Println("Go only has one loop, for")
	}

	i := 0
	for i < 4 {
		fmt.Println("FOUR")
		i++
	}

	i = 0
	for {
		fmt.Println("i is " + strconv.Itoa(i))
		i++

		if i%2 == 1 {
			continue
		}

		if i > 4 {
			break
		}
	}

	fmt.Println("ranges are like arrays or slices, strings, maps, channels")

	ages := []string{
		"28",
		"29",
		"25",
	}

	// _ is index, but we don't use it, so _n
	for _, n := range ages {
		fmt.Println(n)
	}
}
