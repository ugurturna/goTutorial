package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// TCP Server example, required telnet for mac
// brew install telnet
// telnet localhost 8080

func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}

	defer conn.Close()

	fmt.Println("Code got here.")
}
