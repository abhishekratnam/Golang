package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	ch := make(chan int)
	wg.Add(2)

	go func() {
		i := <-ch
		fmt.Println("Calculated Random Value: ", i)
		wg.Done()
	}()

	go func() {
		gen := rand.Intn(500)
		ch <- gen

		wg.Done()
	}()
	wg.Wait()
}
