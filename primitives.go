package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}

// entry point/parent function
func main() { // start of main function, note: main function is also a go routine function!!!
	// all someFunc become asynchronous
	go someFunc("1")            // go routine, main is going to run and fork someFunc off of its own process, not going to wait
	go someFunc("2")            // go routine, main is going to run and fork someFunc off of its own process, not going to wait
	go someFunc("3")            // go routine, main is going to run and fork someFunc off of its own process, not going to wait
	time.Sleep(time.Second * 2) // does not sync the go routine, not creating join point
	// block until someFunc can finish
	fmt.Println("hi")
} // end of main function
