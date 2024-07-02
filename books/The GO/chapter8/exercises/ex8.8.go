package exercises

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func Scan(r io.Reader, lines chan<- string) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		lines <- s.Text()
	}
	// scan will most likely try to read from the connection after it's closed
	// by handleConn. I don't know how to avoid this. Go seems to shun async io
	// in favour of goroutines, so it probably isn't worth avoiding.
	if s.Err() != nil {
		log.Print("scan: ", s.Err())
	}
}

// !+
func HandleConn(c net.Conn) {
	lines := make(chan string)
	ticker := time.NewTicker(4 * time.Second)
	go Scan(c, lines)
	for {
		select {
		case text := <-lines:
			ticker.Reset(4 * time.Second)
			go echo(c, text, 1*time.Second)
		case <-ticker.C:
			c.Close()
			return
		}
	}
}

//!-

func Ex88() {
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
		go HandleConn(conn)
	}
}
