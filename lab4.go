package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func mod(a, b int) int {
	return a % b
}

func main() {
	for {
		fmt.Println("Mini Calculator")
		fmt.Println("1) Add 2) Sub 3) Mul 4) Div 5) Mod 6) Exit")
		fmt.Print("Choose: ")

		var choice int
		fmt.Scan(&choice)

		if choice == 6 {
			break
		}
		if choice < 1 || choice > 6 {
			fmt.Println("Invalid choice")
			continue
		}

		var a, b int
		fmt.Print("Enter a: ")
		fmt.Scan(&a)
		fmt.Print("Enter b: ")
		fmt.Scan(&b)

		if (choice == 4 || choice == 5) && b == 0 {
			fmt.Println("Error: division by zero")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Result:", add(a, b))
		case 2:
			fmt.Println("Result:", sub(a, b))
		case 3:
			fmt.Println("Result:", mul(a, b))
		case 4:
			fmt.Println("Result:", div(a, b))
		case 5:
			fmt.Println("Result:", mod(a, b))
		}
	}
}
