package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// print a wellcome message
	intro()
	// create a channel to indicate when the user wants to quit
	doneChan := make(chan bool)
	// start a goroutine to read user input and run program
	go readUserInput(os.Stdin, doneChan)
	// block until the doneChan gets a value
	<-doneChan
	close(doneChan)
	fmt.Println("Good bye")
}

func intro() {
	fmt.Println("It is Prime")
	fmt.Println("-----------")
	fmt.Println("Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit")
	prompt()
}

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)
	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}
		fmt.Println(res)
		prompt()
	}
}

func prompt() {
	fmt.Print("--->")
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// scan user input
	scanner.Scan()

	// check to see if the user wants to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// try to convert what the user typed into a int
	numberToCheck, err := strconv.Atoi(scanner.Text())

	if err != nil {
		return "Please enter a whole number!", false
	}

	_, msg := isPrime(numberToCheck)
	return msg, false
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime number by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}
	// negative number is not prime number
	if n < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime numbers
			return false, fmt.Sprintf("%d is not prime number because it is divisible by %d", n, i)
		}
	}
	return true, fmt.Sprintf("%d is a prime number", n)
}
