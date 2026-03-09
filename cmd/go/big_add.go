package main

import (
	"fmt"
	"os"
	"strings"
)

// addBigNumbers adds two large integers represented as strings
func addBigNumbers(num1, num2 string) string {
	// Remove leading zeros
	num1 = strings.TrimLeft(num1, "0")
	num2 = strings.TrimLeft(num2, "0")

	if num1 == "" {
		num1 = "0"
	}
	if num2 == "" {
		num2 = "0"
	}

	// Reverse strings for easier digit-by-digit addition
	runes1 := []rune(num1)
	runes2 := []rune(num2)

	result := []rune{}
	carry := 0
	maxLen := len(runes1)

	if len(runes2) > maxLen {
		maxLen = len(runes2)
	}

	for i := 0; i < maxLen; i++ {
		d1 := 0
		d2 := 0

		if i < len(runes1) {
			d1 = int(runes1[len(runes1)-1-i]) - '0'
		}
		if i < len(runes2) {
			d2 = int(runes2[len(runes2)-1-i]) - '0'
		}

		total := d1 + d2 + carry
		result = append(result, rune(total%10+'0'))
		carry = total / 10
	}

	if carry > 0 {
		result = append(result, rune(carry+'0'))
	}

	// Reverse result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func main() {
	if len(os.Args) != 4 || os.Args[2] != "+" {
		fmt.Fprintf(os.Stderr, "Usage: go run big_add.go <num1> + <num2>\n")
		os.Exit(1)
	}

	num1 := strings.TrimSpace(os.Args[1])
	num2 := strings.TrimSpace(os.Args[3])

	result := addBigNumbers(num1, num2)
	fmt.Println(result)
}