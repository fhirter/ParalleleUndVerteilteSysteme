package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	const N int = 10000
	const limit int = 1000
	numbers := numberGenerator(N, limit)

	start := time.Now()

	mean := mean(numbers)
	minimal := minimal(numbers)
	maximal := maximal(numbers)

	log.Printf("main, execution time %s\n", time.Since(start))

	fmt.Println(mean)
	fmt.Println(minimal)
	fmt.Println(maximal)
}

func mean(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	var sum int
	for _, d := range numbers {
		sum += d
	}
	return sum / len(numbers)
}

func minimal(numbers []int) int {
	m := 0
	for i, number := range numbers {
		if i == 0 || number < m {
			m = number
		}
	}
	return m
}

func maximal(numbers []int) int {
	m := 0
	for i, number := range numbers {
		if i == 0 || number > m {
			m = number
		}
	}
	return m
}

func numberGenerator(N int, limit int) []int {
	var numbers []int
	for i := 0; i < N; i++ {
		numbers = append(numbers, rand.Intn(limit))
	}
	return numbers
}
