package main

import (
	"fmt"
	big "math/big"
	"strconv"
	"strings"
)

func containsAllBytes(s string, intArray []int) bool {
	for _, i := range intArray {
		if !strings.Contains(s, strconv.Itoa(i)) {
			return false
		}
	}
	return true
}

func checkIfFactorialContainsAllBytes(intArray []int, n int) int {
	for {
		factorial := new(big.Int).MulRange(1, int64(n))
		factorialString := factorial.String()

		if containsAllBytes(factorialString, intArray) {
			return n
		} else {
			n++
		}
	}
}

func main() {
	var name string
	var surname string
	var slicedName string
	var intArray []int

	fmt.Println("Enter your name: ")
	fmt.Scan(&name)

	fmt.Println("Enter your surname: ")
	fmt.Scan(&surname)

	slicedName = strings.ToLower(name[0:3] + surname[0:3])
	intArray = []int{int(slicedName[0]), int(slicedName[1]), int(slicedName[2]), int(slicedName[3]), int(slicedName[4]), int(slicedName[5])}

	n := checkIfFactorialContainsAllBytes(intArray, 1)

	fmt.Println("Twoja silna liczba: " + strconv.Itoa(n))
}
