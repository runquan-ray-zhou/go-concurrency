package main

import "fmt"

// func someFunc(num string) {
// 	fmt.Println(num)
// }

// // entry point/parent function
// func main() { // start of main function, note. main function is also a go routine function!!!
// 	// all someFunc become asynchronous
// 	go someFunc("1")            // go routine, main is going to run and fork someFunc off of its own process, not going to wait
// 	go someFunc("2")            // go routine, main is going to run and fork someFunc off of its own process, not going to wait
// 	go someFunc("3")            // go routine, main is going to run and fork someFunc off of its own process, not going to wait
// 	time.Sleep(time.Second * 2) // does not sync the go routine, not creating join point
// 	// block until someFunc can finish
// 	fmt.Println("hi")
// } // end of main function

func main() { // main function and func is go routine join together
	myChannel := make(chan string) // make a channel that passes string type data

	// go routine of an anonymous function, forked away and executing on its own.
	go func() {
		myChannel <- "data" // sending data to myChannel
	}()

	// below is blocking, main function is waiting for the channel to close or a message to be received from the channel
	msg := <-myChannel // main function is reading data from the channel, blocking line of code is the join point

	fmt.Println(msg)
}
