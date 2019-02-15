package main

import (
	"fmt"
	"log"
	"math/big"
	"sort"
	"sync"
)

var from int64 = 2
var to int64 = 100

var primes []int
var goRoutines int = 10

func main() {
	wg := new(sync.WaitGroup)
	queue := make(chan int64)
	startGoRoutines(goRoutines, wg, queue)

	go func() {
		wg.Wait()
		log.Println("done waiting")
		close(queue)
	}()

	for elem := range queue {
		primes = append(primes, int(elem))
	}
	sort.Ints(primes)
	fmt.Println(primes)
}

func startGoRoutines(amt int, wg *sync.WaitGroup, queue chan<- int64) {
	totalRange := to - from
	var ranges []int64

	for i := 0; i < amt; i++ {
		ranges = append(ranges, totalRange/int64(amt)) // integer division
	}
	//divide up the remainder
	for i := 0; i < int(totalRange)%amt; i++ {
		ranges[i] += 1
	}
	fmt.Println("Range per goroutine", ranges)
	start := from

	for i := 0; i < amt; i++ {
		wg.Add(1)
		fmt.Println(i, start, "+", ranges[i], "=", start+ranges[i])
		start1 := start
		_range := ranges[i]

		go func() {
			findPrimes(start1, start1+_range, wg, queue)
			defer wg.Done()
		}()
		start += ranges[i]
	}
}

func findPrimes(from int64, to int64, wg *sync.WaitGroup, queue chan<- int64) {
	for i := from; i <= to; i++ {
		if isPrime(i) {
			queue <- i
		}
	}
}

// Checks if the number is prime
func isPrime(number int64) bool {
	if big.NewInt(number).ProbablyPrime(0) {
		return true
	} else {
		return false
	}
}
