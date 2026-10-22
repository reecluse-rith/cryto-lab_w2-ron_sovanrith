package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter a: ")
	fmt.Scan(&a)

	fmt.Print("Enter b: ")
	fmt.Scan(&b)

	result1 := a
	fmt.Println("a =", result1)

	result2 := b
	fmt.Println("b =", result2)

	result1 = a
	result1 += b
	fmt.Println("+= :", result1)

	result1 = a
	result1 -= b
	fmt.Println("-= :", result1)

	result1 = a
	result1 *= b
	fmt.Println("*= :", result1)

	if b != 0 {
		result1 = a
		result1 /= b
		fmt.Println("/= :", result1)

		result1 = a
		result1 %= b
		fmt.Println("%= :", result1)
	} else {
		fmt.Println("Cannot divide by zero")
	}
}