package channel

import (
	"fmt"
)

func receiveAndSend(c chan int) {
	fmt.Printf("Received: %d\n", <-c)
	fmt.Printf("Sending 2...\n")
	c <- 2
}
func recieveOnly(c <-chan int) {
	fmt.Printf("Received: %d\n", <-c)
	// c <- 2 // error
}

func sendOnly(c chan<- int) {
	c <- 2 // OK
	// fmt.Printf("Received: %d\n", <-c) // error
}

func main() {
	myChan := make(chan int)

	go receiveAndSend(myChan)
	myChan <- 1

	fmt.Printf("Value from receiveAndSend: %d\n", <-myChan)
}
