package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	var N = 10000000
	const limit int = 1000
	numbers := numberGenerator(N, limit)

	meanChannel := make(chan int, 1)
	minChannel := make(chan int, 1)
	maxChannel := make(chan int, 1)

	log.Printf("N,sequential,concurrent")

	var meanNumber int
	var minimalNumber int
	var maximalNumber int

	start := time.Now()

	mean(numbers, meanChannel)
	minimal(numbers, minChannel)
	maximal(numbers, maxChannel)

	meanNumber = <-meanChannel
	minimalNumber = <-minChannel
	maximalNumber = <-maxChannel

	sequentialTime := time.Since(start)

	start = time.Now()
	go mean(numbers, meanChannel)
	go minimal(numbers, minChannel)
	go maximal(numbers, maxChannel)

	meanNumber = <-meanChannel
	minimalNumber = <-minChannel
	maximalNumber = <-maxChannel

	concurrentTime := time.Since(start)

	log.Printf("%v, %s, %s", N, sequentialTime, concurrentTime)

	time.Sleep(100 * time.Millisecond)

	log.Println(meanNumber, minimalNumber, maximalNumber)

}

func mean(numbers []int, out chan<- int) {
	if len(numbers) == 0 {
		out <- 0
	}
	var sum int
	for _, d := range numbers {
		sum += d
	}
	out <- sum / len(numbers)
}

func minimal(numbers []int, out chan<- int) {
	m := 0
	for i, number := range numbers {
		if i == 0 || number < m {
			m = number
		}
	}
	out <- m
}

func maximal(numbers []int, out chan<- int) {
	m := 0
	for i, number := range numbers {
		if i == 0 || number > m {
			m = number
		}
	}
	out <- m
}

func numberGenerator(N int, limit int) []int {
	var numbers []int
	for i := 0; i < N; i++ {
		numbers = append(numbers, rand.Intn(limit))
	}
	return numbers
}
