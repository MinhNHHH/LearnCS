package exercises

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func Ex83() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	tcpConn := conn.(*net.TCPConn) // Type assertion to *net.TCPConn
	done := make(chan int)
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- 1 // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	tcpConn.CloseWrite()
	<-done // wait for background goroutine to finish
	tcpConn.Close()
}
