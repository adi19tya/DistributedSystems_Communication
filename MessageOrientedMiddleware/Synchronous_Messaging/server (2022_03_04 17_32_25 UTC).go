package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {
	fmt.Println("Start server...")

	// listen on port 8000
	ln, _ := net.Listen("tcp", ":8000")
	// accept connection
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// get message, output
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))
		// message to be sent
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Message sent: ")
		text, _ := reader.ReadString('\n')
		// send to server
		fmt.Fprintf(conn, text + "\n")
	}
}
