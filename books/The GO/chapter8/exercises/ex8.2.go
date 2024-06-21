package exercises

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/google/uuid"
)

type Server struct {
	clients map[string]*Client
	mutex   sync.Mutex
}

type Client struct {
	ID   string
	conn net.Conn
}

func New() *Server {
	return &Server{
		clients: make(map[string]*Client),
	}
}

func (c *Client) handleCommand(cmd string) {
	command := strings.Split(strings.Trim(cmd, " "), " ")
	switch strings.ToLower(command[0]) {
	case "ls":
		// Get the current directory
		dir, err := ioutil.ReadDir(".")
		if err != nil {
			log.Fatal(err)
		}
		// Print the contents of the directory
		for _, file := range dir {
			io.WriteString(c.conn, fmt.Sprintf("%s\n", file.Name()))
		}
	case "cd":
		path := command[1]
		// Change the current working directory
		err := os.Chdir(path)
		if err != nil {
			log.Fatalf("Error changing directory: %v", err)
		}

	case "close":
		c.conn.Close()
	case "get":
		filename := command[1]
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(c.conn, "Error reading file: %v\n", err)
			return
		}
		_, err = c.conn.Write(data)
		if err != nil {
			log.Printf("Error sending file: %v\n", err)
		}
	}
}
func (s *Server) HandleConn(conn net.Conn) {
	client := &Client{
		ID:   uuid.NewString(),
		conn: conn,
	}
	s.AddClient(client)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		client.handleCommand(scanner.Text())
	}
}

func (s *Server) AddClient(conn *Client) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	fmt.Printf("client %s connected \n", conn.ID)
	s.clients[conn.ID] = conn
}

func (s *Server) Remove(conn Client) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.clients, conn.ID)
	fmt.Printf("client %s disconnected \n", conn.ID)
	defer conn.conn.Close()
}

func Start(url string) {
	listener, err := net.Listen("tcp", url)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	log.Println("Listenning to port: 8000")
	defer listener.Close()
	server := New()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error connect:", err)
			continue
		}
		go server.HandleConn(conn)
	}
}
