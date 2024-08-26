package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primiTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number"},
		{"is not prime", 8, false, "8 is not prime number because it is divisible by 2"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primiTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got fale", e.name)
		}
		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("expected %s but got %s", e.name, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write
	r, w, _ := os.Pipe()

	// set os.Stdout to our write Pipe

	os.Stdout = w

	prompt()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	// perform our test
	if string(out) != "--->" {
		t.Errorf("incorrect prompt: expected ---> but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write
	r, w, _ := os.Pipe()

	// set os.Stdout to our write Pipe

	os.Stdout = w

	intro()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	// perform our test
	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("intro text not correct; got %s", string(out))
	}
}

func Test_checkNumebrs(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "prime", input: "7", expected: "7 is a prime number"},
		{name: "quit", input: "q", expected: ""},
		{name: "negative", input: "-11", expected: "Negative numbers are not prime, by definition!"},
		{name: "quit", input: "q", expected: ""},
		{name: "is not prime", input: "8", expected: "8 is not prime number because it is divisible by 2"},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)
		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s, but got %s", e.name, e.expected, res)
		}
	}

}

func Test_readUserInput(t *testing.T) {
	// to test thif function, we need a channel, and an instance of an io.reader
	doneChan := make(chan bool)

	// create a reference to a bytes.Buffer'

	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)

	<-doneChan

	close(doneChan)
}
