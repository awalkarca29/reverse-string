package main

import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s) // Menggunakan rune untuk menangani karakter multi-byte
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Menukar karakter
	}
	return string(runes)
}

func main() {
	input := "Halo, Dunia!"
	reversed := reverseString(input)
	fmt.Println(reversed) // Output: !ainuD ,olaH
}
