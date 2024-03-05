package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	log.SetFlags(0)
	concurrency := 1
	sampleSize := 10
	log.Printf("Number of slices: %v", concurrency)
	for i := 0; i < 8; i++ {
		calculationTime := concurrentMean(sampleSize, concurrency)
		log.Printf("%v, %s", sampleSize, calculationTime.Truncate(time.Microsecond).String())
		sampleSize = sampleSize * 10
	}
}

func concurrentMean(sampleSize int, concurrency int) time.Duration {
	numberOfSlices := concurrency

	if sampleSize%numberOfSlices > 0 {
		panic("sampleSize divided by number of slices should not leave a remainder!")
	}

	const maxNumber int = 1000
	numbers := numberGenerator(sampleSize, maxNumber)

	resultChannel := make(chan int, numberOfSlices)
	dataChannel := make(chan []int, numberOfSlices)

	start := time.Now()
	// fill data channel
	for i := 0; i < numberOfSlices; i++ {
		sliceSize := sampleSize / numberOfSlices
		dataChannel <- numbers[sliceSize*i : sliceSize*(i+1)]
	}
	// run calculations
	for i := 0; i < numberOfSlices; i++ {
		go mean(dataChannel, resultChannel)
	}

	//get results and calculate mean of means
	var meanNumber = 0
	for i := 0; i < numberOfSlices; i++ {
		intermediateMean := <-resultChannel
		meanNumber += intermediateMean / numberOfSlices
	}
	//log.Println(meanNumber)
	duration := time.Since(start)

	return duration
}

func mean(in chan []int, out chan<- int) {
	numbers := <-in
	if len(numbers) == 0 {
		out <- 0
	}
	var sum int
	for _, d := range numbers {
		sum += d
	}
	out <- sum / len(numbers)
}

func numberGenerator(N int, limit int) []int {
	var numbers []int
	for i := 0; i < N; i++ {
		numbers = append(numbers, rand.Intn(limit))
	}
	return numbers
}
