package main

// func main() {
// 	// charChannel := make(chan string, 3) // use buffered channel is needed for asynchronous communication send and forget up to the buffered capacity
// 	// chars := []string{"a", "b", "c"}

// 	// for _, s := range chars {
// 	// 	select {
// 	// 	case charChannel <- s:
// 	// 	}
// 	// }

// 	// close(charChannel)

// 	// for result := range charChannel { // data stay in channel
// 	// 	fmt.Println(result)
// 	// }
// 	done := make(chan bool)

// 	go doWork(done)

// 	time.Sleep(time.Second * 3)

// 	close(done) // parent go routine cancel the child go routine after 3 seconds
// }

// func doWork(done <-chan bool) { // read only channel
// 	for {
// 		select { // infinitely running go routine
// 		// removing the select, would not have the case to have parent cancel the go routine
// 		case <-done:
// 			return
// 		default:
// 			fmt.Println("DOING WORK")
// 		}
// 	}
// }
