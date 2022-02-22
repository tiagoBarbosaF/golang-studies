package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	ints := map[string]int64{
		"first":  56,
		"second": 31,
	}

	floats := map[string]float64{
		"first":  43.87,
		"second": 12.36,
	}

	fmt.Printf("\nNon-Generic Sums: %v and %v\n", SumInts(ints), SumFloats(floats))
	fmt.Printf("Generic Sums: %v and %v\n", SumIntsOrFloats[string, int64](ints), SumIntsOrFloats[string, float64](floats))
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
	fmt.Printf("Generic Sums with Constraint: %v and %v\n", SumNumbers(ints), SumNumbers(floats))
}

// SumInts Non-Generic functions
func SumInts(m map[string]int64) int64 {
	var sumInt int64

	for _, value := range m {
		sumInt += value
	}

	return sumInt
}

// SumFloats Non-Generic functions
func SumFloats(m map[string]float64) float64 {
	var sumFloat float64

	for _, value := range m {
		sumFloat += value
	}

	return sumFloat
}

// SumIntsOrFloats Generic function
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var sum V

	for _, value := range m {
		sum += value
	}
	return sum
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, value := range m {
		sum += value
	}
	return sum
}
