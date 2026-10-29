package main

import "fmt"

func myAnd(a, b int) int {
	return a & b
}

func myOr(a, b int) int {
	return a | b
}

func myXor(a, b int) int {
	return a ^ b
}

func myNot(a int) int {
	return ^a
}

func main() {
	var a, b int

	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)

	fmt.Println("AND:", myAnd(a, b))
	fmt.Println("OR:", myOr(a, b))
	fmt.Println("XOR:", myXor(a, b))
	fmt.Println("NOT a:", myNot(a))
	fmt.Println("Left shift:", a<<1)
	fmt.Println("Right shift:", a>>1)

	x := a
	x += b
	fmt.Println("+= :", x)

	x = a
	x -= b
	fmt.Println("-= :", x)

	x = a
	x *= b
	fmt.Println("*= :", x)

	if b != 0 {
		x = a
		x /= b
		fmt.Println("/= :", x)

		x = a
		x %= b
		fmt.Println("%= :", x)
	} else {
		fmt.Println("Cannot divide by zero")
	}
}
