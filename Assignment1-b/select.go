package main

import (
	"fmt"
	"time"
)

func even(l1 []int, c chan int) {
	sum := 0
	for i, val := range l1 {
		if l1[i] % 2 == 0 {
			sum += val
		}
	}
	time.Sleep(3 * time.Millisecond)
	c <- sum
}

func odd(l2 []int, c chan int) {
	sum := 0
	for i, val := range l2 {
		if l2[i] % 2 != 0 {
			sum += val
		}
	}
	time.Sleep(2 * time.Millisecond)
	c <- sum
}

func main() {
	l := [6]int{2, 9, 3, 4, 10, 1}
	c1 := make(chan int)
	c2 := make(chan int)

	go even(l[:], c1)
	go odd(l[:], c2)

	select {
	case a := <-c1:
		fmt.Println("Even may sum greater!", a)
	case b := <-c2:
		fmt.Println("Odd may sum greater!", b)
	default:
		fmt.Println("Evens and odds sum up equal!")
	}
}
