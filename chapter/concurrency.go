package chapter

import (
	"fmt"
	"strconv"
	"time"
)

// Concurrency Chapter
func Concurrency() {
	go fmt.Println("\n----> I'm in another thread! <----\n ")

	fmt.Println("Communication between threads controls actions, not the data itself")
	fmt.Println("Don't communicate by sharing memory, share memory by communicating")
	fmt.Println("One thread owns the data and communicates about it using a channel")
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Goroutine is a lightweight thread")
	fmt.Println("Keyword \"go\" means run this concurrently, don't block")
	fmt.Println()
	fmt.Println("Channels are how threads communicate")

	done := make(chan bool, 1) // buffer is one
	go func() {
		fmt.Println("hmm, ok")
		done <- true
	}()

	if <-done {
		fmt.Println("I'm done!")
	}

	fmt.Println("\nThread Race!")
	multi := make(chan string)
	aCount := 0
	bCount := 0

	go func() {
		for i := 0; i < 10; i++ {
			multi <- "A"
		}
		close(multi)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			multi <- "B"
		}
		close(multi)
	}()

	for chanLoop := range multi {
		fmt.Print(chanLoop)
		if chanLoop == "A" {
			aCount++
		} else {
			bCount++
		}
	}

	fmt.Println()
	fmt.Println(strconv.Itoa(aCount) + " to " + strconv.Itoa(bCount))

	switch {
	case aCount == bCount:
		fmt.Println("Tie!")
	case aCount > bCount:
		fmt.Println("A Wins")
	default:
		fmt.Println("B Wins")
	}

	fmt.Println("\nSelect statement is like a switch, but for thread/channel communication")

	fmt.Println()
	c := make(chan string)
	c2 := make(chan string)
	go func() {
		c <- "A"
	}()
	go func() {
		c2 <- "B"
	}()

	running := true

	for running {
		select {
		case s, ok := <-c:
			if ok {
				fmt.Println(s, ":1")
				running = false
			} else {
				return
			}
		case s, ok := <-c2:
			if ok {
				fmt.Println(s, ":2")
				running = false
			} else {
				return
			}
		default:
			fmt.Println("waiting..")
		}
	}
	// something is buggy with the above example
	// panic: send on closed channel
}
