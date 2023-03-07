package main

import (
	"fmt"
	big "math/big"
	"strings"
)

func calculateFactorial(n int) *big.Int {
	factorial := big.NewInt(1)
	for i := 2; i <= n; i++ {
		factorial.Mul(factorial, big.NewInt(int64(i)))
	}
	return factorial
}

func containsAllBytes(s string, byteArray []byte) bool {
	for _, b := range byteArray {
		if !strings.Contains(s, string(b)) {
			return false
		}
	}
	return true
}

func checkIfFactorialContainsAllBytes(byteArray []byte, n int) int {
	for {
		factorial := calculateFactorial(n)
		factorialString := factorial.String()

		if containsAllBytes(factorialString, byteArray) {
			return n
		}

		n++
	}
}

func main() {
	var name string
	var surname string
	var slicedName string
	var byteArray []byte

	fmt.Println("Enter your name: ")
	fmt.Scan(&name)

	fmt.Println("Enter your surname: ")
	fmt.Scan(&surname)

	slicedName = strings.ToLower(name[0:3] + surname[0:3])
	byteArray = []byte(slicedName)

	fmt.Println(byteArray)

	n := checkIfFactorialContainsAllBytes(byteArray, 1)

	fmt.Println(n)
}
