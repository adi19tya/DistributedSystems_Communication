
package main

import "net"
import "fmt"
import "bufio"

func main() {
	fmt.Println("Start server...")



	// listen on port 9000(server)
	ln1, _ := net.Listen("tcp", ":9000")

	// accept connection (port 8000)
	conn, _ := ln1.Accept()


	// run loop forever (or until ctrl-c)
	for {
		// get message, output
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", message)
	}
}
