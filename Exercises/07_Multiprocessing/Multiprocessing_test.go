package main

import (
	"slices"
	"testing"
)

func TestMean(t *testing.T) {
	numbers := []int{90, 81, 78, 95, 79, 72, 85}
	expected := 82

	resultChannel := make(chan int, 1)
	dataChannel := make(chan []int, 1)

	dataChannel <- numbers

	mean(dataChannel, resultChannel)

	given := <-resultChannel

	if expected != given {
		t.Errorf(`mean(): wrong result. Expected %v, given %v`, expected, given)
	}
}

func TestNumberGenerator(t *testing.T) {
	N := 100
	limit := 1000
	numbers := numberGenerator(N, limit)

	if len(numbers) != N {
		t.Error("given slice has wrong length")
	}

	biggestNumber := slices.Max(numbers)
	smallestNumber := slices.Min(numbers)

	if biggestNumber == 0 {
		t.Errorf(`expected biggest number to not be 0. %v given.`, biggestNumber)
	}

	if biggestNumber >= limit {
		t.Errorf(`expected biggest number to be less than %v. %v given.`, limit, biggestNumber)
	}

	if smallestNumber < 0 {
		t.Errorf(`expected smallest number to not be negative. %v given.`, smallestNumber)
	}
}
