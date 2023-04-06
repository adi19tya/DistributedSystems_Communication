package main

import "net"
import "fmt"
import "bufio"
import "os"

var conn *net.Conn


func main() {

	// connect to server
	conn, _ := net.Dial("tcp", "127.0.0.1:8000")

	for {

		// message to be sent
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Message Sent: ")
		text, _ := reader.ReadString('\n')
		// send to server
		fmt.Fprintf(conn, text + "\n")
		// wait for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received: "+message)
	}
}