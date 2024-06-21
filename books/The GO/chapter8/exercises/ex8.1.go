package exercises

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func client() {
	locationtz := map[string]string{
		"NewYork": "localhost:8010",
		"London":  "localhost:8020",
		"Tokyo":   "localhost:8030",
	}
	for tz, url := range locationtz {
		go connectToServer(url, tz)
	}
	select {}
}

func server() {
	// Define the port flag
	port := flag.Int("port", 8020, "Server accept port")
	tz := flag.String("tz", "Asia/Tokyo", "Location timezone")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server listen port %d", *port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, *tz) // handle one connection at a time
	}
}
func handleConn(c net.Conn, tz string) {
	defer c.Close()
	location, err := time.LoadLocation(tz)
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	for {
		_, err := io.WriteString(c, time.Now().In(location).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func connectToServer(url, tz string) {
	conn, err := net.Dial("tcp", url)
	log.Printf("Connect to %s\n", url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	input := bufio.NewScanner(conn)

	for input.Scan() {
		fmt.Printf("%s: %s\n", tz, input.Text())
	}
}
