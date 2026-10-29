package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)

	fmt.Println("Both positive:", a > 0 && b > 0)
	fmt.Println("One greater:", a > b || b > a)
	fmt.Println("Not equal:", !(a == b))
}
