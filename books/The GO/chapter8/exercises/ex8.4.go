// Exercise 8.4: Modify the reverb2 server to use a sync.WaitGroup per connection to count
// the number of active echo goroutines. When it falls to zero, close the write half of the TCP
// connection as described in Exercise 8.3. Verify that your modified netcat3 client from that
// exercise waits for the final echoes of multiple concurrent shouts, even after the standard input
// has been closed.

package exercises

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func Echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure the counter is decremented when the goroutine completes
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func HandleConn1(c *net.TCPConn) {
	var wg sync.WaitGroup // Create a new WaitGroup for each connection
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go Echo(c, input.Text(), 1*time.Second, &wg)
	}
	wg.Wait()      // Wait for all echo goroutines to finish for this connection
	c.CloseWrite() // Close the write half of the connection
	c.Close()      // Fully close the connection
}

func EX84() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		tcpConn := conn.(*net.TCPConn)
		go HandleConn1(tcpConn)
	}
}
