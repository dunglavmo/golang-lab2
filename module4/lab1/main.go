package main

import (
	"fmt"
	"math/rand"
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

func main() {
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
