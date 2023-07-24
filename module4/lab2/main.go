package main

import (
	"fmt"
	"time"
)

const (
	redDuration    = 15 * time.Second
	greenDuration  = 30 * time.Second
	yellowDuration = 3 * time.Second
)

type States int

const (
	Red States = iota
	Green
	Yellow
)

func (state States) String() string {
	switch state {
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Unknown"
	}
}

func trafficLight(States chan<- States) {
	// Start with the Red light
	States <- Red

	for {
		// Transition from Red to Green
		time.Sleep(redDuration)
		States <- Green

		// Transition from Green to Yellow
		time.Sleep(greenDuration)
		States <- Yellow

		// Transition from Yellow to Red
		time.Sleep(yellowDuration)
		States <- Red
	}
}

func main() {
	States := make(chan States)

	go trafficLight(States)

	for {
		state := <-States
		fmt.Printf("Traffic Light: %s\n", state)
	}
}
