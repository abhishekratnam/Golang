package main

import (
	"fmt"
	"time"
)

// The pinger prints a ping and waits for a pong
func p1(p1 <-chan int, p2 chan<- int) {
	for {
		<-p1
		fmt.Println("ping")
		time.Sleep(time.Second)
		p2 <- 1
	}
}

// The ponger prints a pong and waits for a ping
func p2(p1 chan<- int, p2 <-chan int) {
	for {
		<-p2
		fmt.Println("pong")
		time.Sleep(time.Second)
		p1 <- 1
	}
}

func main() {
	ping := make(chan int)
	pong := make(chan int)

	go p1(ping, pong)
	go p2(ping, pong)

	// The main goroutine starts the ping/pong by sending into the ping channel
	ping <- 1

	select {}

}
