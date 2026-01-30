package main

import (
	"fmt"
)

func even(l1 []int) int {
	sum := 0
	for i, val := range l1 {
		if l1[i] % 2 == 0 {
			sum += val
		}
	}
	return int(sum)
}

func odd(l2 []int) int {
	sum := 0
	for i, val := range l2 {
		if l2[i] % 2 != 0 {
			sum += val
		}
	}
	return int(sum)
}

func main() {
	l := [6]int{2, 9, 3, 4, 10, 15}

	a := even(l[:])
	b := odd(l[:])

	switch {
	case a > b:
		fmt.Println("Sum of evens is greater!")
	case b > a:
		fmt.Println("Sum of odds is greater!")
	default:
		fmt.Println("Evens and odds sum up equal!")
	}
}
