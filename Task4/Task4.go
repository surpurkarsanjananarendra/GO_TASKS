package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func squareWorkers(numbers chan int, squares chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range numbers {
		squares <- num * num
	}
	close(squares)
}

func printAggregate(squares chan int) {
	total := 0
	fmt.Println("\n\nSquares of numbers: ")
	for i := range squares {
		fmt.Print(i, " ")
		total += i
	}
	fmt.Printf("\n\nAggregate is %d", total)
}

func main() {
	var size int
	fmt.Println("Enter size: ")
	fmt.Scan(&size)

	nums := make([]int, size)

	fmt.Println("\nEnter numbers:")
	for i := 0; i < size; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println("Numbers: ")
	for _, v := range nums {
		fmt.Print(v, " ")
	}

	numbers := make(chan int, size)
	squares := make(chan int, size)

	for i := range nums {
		numbers <- nums[i]
	}
	close(numbers)

	wg.Add(1)
	go squareWorkers(numbers, squares, &wg)

	printAggregate(squares)

	wg.Wait()
}
