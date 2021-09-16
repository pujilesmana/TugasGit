package main

import "fmt"

func minus(a, b int) int {
	return a - b
}

func plus(a, b int) int {

	return a + b
}

func main() {
	fmt.Println("Plus")
	hasil := plus(2, 3)
	hasil2 := minus(2, 3)

	fmt.Printf("%d", hasil)
	fmt.Printf("%d", hasil2)
}
