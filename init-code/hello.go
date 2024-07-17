package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("hello")
	one()
	three()
	four()
}

func one() {
	fmt.Println("World!")
	two()
	fmt.Println(quote.Go())
	four()
}

func two() {
	fmt.Println("two")
}

func three() {
	fmt.Println("four")
}

func four() {
	print("hello")
	for {
		fmt.Println("gr")
	}
}
