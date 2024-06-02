package main

import "fmt"
import "rsc.io/quote"

func main() {
	fmt.Println("hello")
	one()
}

func one() {
	fmt.Println("World!")
	two()
	fmt.Println(quote.Go())
}

func two() {
	fmt.Println("two")
}
