package main

import (
	"fmt"
)

type Number interface {
	int64 | float64
}

func sumInt(m map[string]int64) int64 {
	var sum int64
	for _, v := range m {
		sum += v
	}
	return sum
}

func sumFloat(m map[string]float64) float64 {
	var sum float64
	for _, v := range m {
		sum += v
	}
	return sum
}

// instead of the above non generic functions, we can use the below generic function

func sum[K comparable, T Number](m map[K]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	ints := map[string]int64{"a": 1, "b": 2, "c": 3}
	floats := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
	fmt.Println(sumInt(ints))
	fmt.Println(sumFloat(floats))
	fmt.Println("Hello, World!")
	fmt.Println("With respect to the generics the output is below")
	fmt.Println(sum(ints))
	fmt.Println(sum(floats))
}
