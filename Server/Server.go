package main

import (
	"errors"
	"fmt"
	"net"
	"os"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	} else {
		fmt.Println("yes listning to port 8081")
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on tcp", ":8081")
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		defer conn.Close()
		handleRequest(conn)
		/*if hndlRqstErr != nil {
			fmt.Println("Operation failed on connection")
		}*/
	}

}
func checkmessage(message string) string {
	if message != "" {
		if message == "Hello" {
			message = "Hi"
		} else if message == "Name" {
			message = "Chitty"
		} else if message == "Goodbye" {
			message = "bye"
		} else {

			message = "Sorry!!"
		}

	} else {

		//panic("sorry message can't be interpreted")
		message = "sorryMCI"
	}
	return message

}
func generateMsg(conn net.Conn, buf []byte) error {
	reqLen, Rerr := conn.Read(buf)
	if Rerr != nil {
		return errors.New("errReadEmpty")
	}
	message := string(buf[:reqLen-2])
	rplymessage := checkmessage(message)
	fmt.Println(" msg is : ", rplymessage)
	msgln, Werr := conn.Write([]byte(rplymessage))
	fmt.Println("lenghth of message is ", msgln)
	if Werr != nil {
		return errors.New("errWriteFullHR")
	}
	return nil
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		err := generateMsg(conn, buf)
		if err != nil {

			fmt.Println("Error : ", err)
		}
	}
}
