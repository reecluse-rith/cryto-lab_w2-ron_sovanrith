package main

import (
	"bufio"
	"fmt"
	"os"
)

func xorEncrypt(text string, key byte) string {
	data := []byte(text)
	for i := range data {
		data[i] ^= key
	}
	return string(data)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter message: ")
	scanner.Scan()
	text := scanner.Text()

	fmt.Print("Enter key: ")
	scanner.Scan()
	key := scanner.Text()

	if key == "" {
		fmt.Println("Key cannot be empty")
		return
	}

	encrypted := xorEncrypt(text, key[0])
	decrypted := xorEncrypt(encrypted, key[0])

	fmt.Printf("Encrypted: %x\n", encrypted)
	fmt.Println("Decrypted:", decrypted)
}
