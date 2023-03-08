package main

import (
	"fmt"
	"math"
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

func fibFrequency(n int) int {

	cache[n]++

	switch n {
	case 0, 1:
		return n
	default:
		return fibFrequency(n-1) + fibFrequency(n-2)
	}
}

func fibFrequencyClosestToStrongNumber(cache map[int]int, strongNumber int) int {
	minDiff := float64(strongNumber - cache[len(cache)-1])
	small := len(cache) - 1

	for key, value := range cache {
		diff := strongNumber - value
		if math.Abs(float64(diff)) < minDiff {
			minDiff = math.Abs(float64(diff))
			small = key
		}
	}

	return small
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

	fmt.Println("Your 'strong' number: " + strconv.Itoa(n))

	fibFrequency(30)

	// fmt.Println(cache)

	fmt.Println("Your 'weak' number: " + strconv.Itoa(fibFrequencyClosestToStrongNumber(cache, n)))

}

var cache = make(map[int]int)

var minDiff int
