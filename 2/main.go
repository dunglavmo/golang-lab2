package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"time"
)

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func lab1() {
	messages := make(chan string)

	go func() {
		for {
			message := randStr(rand.Intn(20))
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			messages <- message
		}
	}()

	for {
		message := <-messages
		fmt.Println("Random string:", message)
	}
}

func scanPort(protocol, host string, port int) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout(protocol, address, 2*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func lab2() {
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

func lab3() {
	http.HandleFunc("/", helloServer)
	http.ListenAndServe(":8080", nil)
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	// lab1()
	// lab2()
	lab3()
}
