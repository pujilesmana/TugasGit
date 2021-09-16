package main

import "fmt"

func plus(a, b int) int {

	return a + b
}

func main() {
	fmt.Println("Plus")
	hasil := plus(2, 3)

	fmt.Printf("%d", hasil)
}
