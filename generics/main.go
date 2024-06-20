package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumInts(m map[string]int64) int64 {
	var s int64

	for _, v := range m {
		s += v
	}
	return s
}

func Sumfloats(m map[string]float64) float64 {
	var s float64

	for _, v := range m {
		s += v
	}

	return s
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	floats := map[string]float64{
		"first":  34.01,
		"second": 12.23,
	}

	fmt.Printf("Non-generic sums: %v and %v\n", SumInts(ints), Sumfloats(floats))
	fmt.Printf(
		"Generic sums: %v and %v\n",
		SumIntOrFloat[string, int64](ints),
		SumIntOrFloat[string, float64](floats),
	)
}

func SumIntOrFloat[K comparable, V Number](m map[K]V) V {
	var s V

	for _, v := range m {
		s += v
	}

	return s
}
