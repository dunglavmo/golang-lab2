package main

import (
	"fmt"
	"net"
	"time"
)

func scanPort(protocol, host string, port int) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout(protocol, address, 2*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func main() {
	host := "localhost"

	results := make(chan bool)

	for i := 1; i <= 65535; i++ {
		go func(port int) {
			isOpen := scanPort("tcp", host, port)
			results <- isOpen
		}(i)
	}

	for i := 1; i <= 65535; i++ {
		open := <-results
		if open {
			fmt.Printf("Port %d is open\n", i)
		} else {
			fmt.Printf("Port %d is closed\n", i)
		}
	}
}
