package main

import (
	"testing"
)

func TestMean(t *testing.T) {
	numbers := []int{90, 81, 78, 95, 79, 72, 85}
	expected := 82
	given := mean(numbers)

	if expected != given {
		t.Errorf(`mean(): wrong result. Expected %v, given %v`, expected, given)
	}
}

func TestMin(t *testing.T) {
	numbers := []int{90, 81, 78, 95, 79, 72, 85}
	expected := 72

	given := minimal(numbers)

	if expected != given {
		t.Errorf(`mean(): wrong result. Expected %v, given %v`, expected, given)
	}
}

func TestMax(t *testing.T) {
	numbers := []int{90, 81, 78, 95, 79, 72, 85}
	expected := 95

	given := maximal(numbers)

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

	biggestNumber := maximal(numbers)
	smallestNumber := minimal(numbers)

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
