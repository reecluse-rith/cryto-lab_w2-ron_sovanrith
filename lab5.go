package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter text: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	binary := ""
	for _, b := range []byte(text) {
		binary += fmt.Sprintf("%08b", b)
	}

	fmt.Println("Binary:", binary)
	fmt.Printf("Hex: %x\n", text)
	fmt.Println("Base64:", base64.StdEncoding.EncodeToString([]byte(text)))
}
