package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {

	//optimizeConcurrent()
	concurrentVsSequential()
}

func optimizeConcurrent() {
	numberOfSlices := 2

	log.SetFlags(0)
	var meanNumber int
	const limit int = 1000
	var N = 10
	//numbers := numberGenerator(N, limit)
	var channels []chan int
	for i := 0; i < numberOfSlices; i++ {
		channels = append(channels, make(chan int, 1))
	}

	start := time.Now()
	for i := 0; i < numberOfSlices; i++ {
		//	go mean(numbers[], channels[i])
	}
	for i := 0; i < numberOfSlices; i++ {
		_ = <-channels[i]
	}

	concurrentTime := time.Since(start)

	log.Printf("%v, %s", N, concurrentTime.Truncate(time.Millisecond).String())
	log.Println(meanNumber)

}

func concurrentVsSequential() {
	log.SetFlags(0)

	var meanNumber int
	var minimalNumber int
	var maximalNumber int

	log.Printf("N,sequential,concurrent")

	var N = 10
	const limit int = 1000
	for i := 0; i < 10; i++ {
		numbers := numberGenerator(N, limit)

		meanChannel := make(chan int, 1)
		minChannel := make(chan int, 1)
		maxChannel := make(chan int, 1)

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

		log.Printf("%v, %s, %s", N, sequentialTime.Truncate(time.Microsecond).String(), concurrentTime.Truncate(time.Microsecond).String())

		time.Sleep(100 * time.Millisecond)

		N = N * 10
	}
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
