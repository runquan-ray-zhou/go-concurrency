package main

import "fmt"

func sliceToChannel(nums []int) <-chan int { // return read only channel
	out := make(chan int) // return an unbuffered channel
	go func() {
		for _, n := range nums { // asynchronous
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// input
	nums := []int{2, 3, 4, 7, 1}
	// stage 1
	dataChannel := sliceToChannel(nums)
	// stage 2
	finalChannel := sq(dataChannel)
	// stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
}
