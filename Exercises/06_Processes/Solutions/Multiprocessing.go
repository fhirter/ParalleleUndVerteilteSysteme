package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numbers []int

	var mean int = mean(numbers)
	fmt.Println(mean)
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

func numberGenerator(N int) []int {
	var numbers []int
	for i := 0; i < N; i++ {
		numbers = append(numbers, rand.Int())
	}
	return numbers
}
