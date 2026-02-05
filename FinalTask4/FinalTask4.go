package main

import (
	"fmt"
	"sync"
)

func squareWorker(num int, squares chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	squares <- num * num
}

func printAggregate(squares <-chan int, result chan<- int) {
	total := 0

	fmt.Println("\nSquares of numbers:")
	for sq := range squares {
		fmt.Print(sq, " ")
		total += sq
	}
	result <- total
}

func main() {
	var size int
	fmt.Print("\nEnter size: ")
	fmt.Scan(&size)

	numbers := make([]int, size)
	fmt.Println("\nEnter numbers:")
	for i := 0; i < size; i++ {
		fmt.Scan(&numbers[i])
	}

	squares := make(chan int)
	result := make(chan int)

	var wg sync.WaitGroup

	go printAggregate(squares, result)

	for _, val := range numbers {
		wg.Add(1)
		go squareWorker(val, squares, &wg)
	}

	/*go func() {
		wg.Wait()
		close(squares)
	}()*/

	wg.Wait()
	close(squares)

	aggregate := <-result
	fmt.Println("\n\nAggregate of Numbers:", aggregate)
	fmt.Println("\nTask Done!")
}
