package main

import "net"
import "fmt"
import "bufio"
import "time"

//declaring the queue as a global variable
var queue  = make([]string, 0, 20)

func push(element string){
	queue = append(queue, element) // Simply append to enqueue.
	// log q content
	fmt.Println(queue)
}

func pull() string {
	// log q content
	if len(queue) == 0 {
		// log "returning null from pull"
		return ""
	}

	element := queue[0] // The first element is the one to be dequeued.
	// Slice off the element once it is dequeued.
	queue = queue[1:]
	// log q content
	// log returned element
	return element
}

func getMessage(conn net.Conn){
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(message)

	push(message)
}

func sendMessage(conn1 net.Conn){
	var message string
	message = pull()

	if message != "" {
		fmt.Fprintf(conn1, message+"\n")
	}
}

func sendingLoop(conn1 net.Conn){
	for{
		time.Sleep(10 * time.Second)
		sendMessage(conn1)
	}
}

func main() {
	fmt.Println("Starting Messaging Queue...")

	// connect to port 9000
	conn1, _ := net.Dial("tcp", "127.0.0.1:9000")


	// listen on port 8000(client)
	ln, _ := net.Listen("tcp", ":8000")

	// accept connection (port 8000)
	conn, _ := ln.Accept()


	// run loop forever (or until ctrl-c)
	go sendingLoop(conn1)

	for {
		// get message
		getMessage(conn)
	}

}