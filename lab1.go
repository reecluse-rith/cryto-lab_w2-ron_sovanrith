package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)

	fmt.Println("a =", a)
	fmt.Println("b =", b)

	r := a
	r += b
	fmt.Println("+= :", r)

	r = a
	r -= b
	fmt.Println("-= :", r)

	r = a
	r *= b
	fmt.Println("*= :", r)

	if b != 0 {
		r = a
		r /= b
		fmt.Println("/= :", r)

		r = a
		r %= b
		fmt.Println("%= :", r)
	} else {
		fmt.Println("Cannot divide by zero")
	}
}
