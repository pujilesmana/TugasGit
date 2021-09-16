package main

import "fmt"

func minus(a, b int) int {
	return a - b
}

func plus() {

	fmt.Println("nice")
}

func main() {
	fmt.Println("Plus")
	hasil2 := minus(2, 3)
	plus()
	fmt.Printf("%d", hasil)
	fmt.Printf("%d", hasil2)
}
