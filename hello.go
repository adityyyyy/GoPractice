package main

import "fmt"

func main() {
	fmt.Println("hello")
	one()
}

func one() {
	fmt.Println("World!")
	two()
}

func two() {
	fmt.Println("two")
}
