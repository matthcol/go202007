package main

import (
	"calcul"
	"fmt"
	"time"
)

func main() {
	c := make(chan uint)
	// send goroutines in //
	go gofibo(4, c)
	go gofibo(10, c)
	go gofibo(7, c)
	// wait for all results via a channel
	n1 := <-c
	n2 := <-c
	n3 := <-c
	// then deal with all results
	fmt.Println(n1, n2, n3)
}

func gofibo(rank uint, c chan uint) {
	n := calcul.Fibo(rank)
	time.Sleep(2 * time.Second)
	c <- n
}
