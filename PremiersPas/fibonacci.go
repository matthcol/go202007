package main

import (
	"fmt"
	"time"
)

func main() {
	rank := uint(45)
	t0 := time.Now()
	nb := fibo(rank)
	t1 := time.Now()
	deltaLoop := t1.Sub(t0)
	fmt.Printf("Le terme %d de la suite de Fibonacci (loop) est %d (time %v)\n", rank, nb, deltaLoop)
	//go timefibor(47)
	//go timefibor(47)
	//go timefibor(47)
	n1, d1 := timefibor(5)
	fmt.Printf("%v %v\n", n1, d1)
	n2, _ := timefibor(6)
	fmt.Printf("%v\n", n2)
	_, d3 := timefibor(7)
	fmt.Printf("%v\n", d3)
	timefibor(8)
}

func fibo(rank uint) uint {
	var a uint = 0
	var b uint = 1
	switch rank {
	case 0:
		return a
	case 1:
		return b
	default:
		for i := uint(2); i < rank+1; i++ {
			c := a + b
			a, b = b, c
		}
		return b
	}
}

func fibor(rank uint) uint {
	switch rank {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibor(rank-2) + fibor(rank-1)
	}
}

// func timefibor(rank uint) {
// 	t0 := time.Now()
// 	nb := fibor(rank)
// 	t1 := time.Now()
// 	deltaRecursive := t1.Sub(t0)
// 	fmt.Printf("Le terme %d de la suite de Fibonacci (recursive) est %d (time %v)\n", rank, nb, deltaRecursive)
// }

func timefibor(rank uint) (uint, time.Duration) {
	t0 := time.Now()
	nb := fibor(rank)
	t1 := time.Now()
	deltaRecursive := t1.Sub(t0)
	return nb, deltaRecursive
}
