package main

import (
	"fmt"
)

func even(n int) (b bool) {
	if n%2 == 0 {
		b = true
	} else {
		b = false
	}
	return b
}

func odd(n int) (b bool) {
	if even(n) {
		b = false
	} else {
		b = true
	}
	return b
}

func collatz(n int, temp []int) []int {
	switch {
	case n == 1:
		break

	case even(n):
		n = n / 2
		temp = append(temp, n)
		return collatz(n, temp)

	case odd(n):
		n = (n * 3) + 1
		temp = append(temp, n)
		return collatz(n, temp)
	}
	return temp
}

func main() {
	//slice to hold the Collatz sequence for each number in range
	temp := []int{}

	//length of the longest Collatz sequence in provided range
	var longest int

	//number that produces the longest Collatz sequence in provided range
	var number int

	//lower bound of range to check
	var l int

	//upper bound of range to check
	var h int

	fmt.Println("Enter the first number of range to check: ")
	_, errl := fmt.Scan(&l)
	if errl != nil {
		fmt.Println(errl)
		return
	}

	fmt.Println("Enter the last number of range to check: ")
	_, errh := fmt.Scan(&h)
	if errh != nil {
		fmt.Println(errh)
		return
	}

	for i := l; i < h+1; i++ {
		temp = collatz(i, temp)
		if len(temp) > longest {
			longest = len(temp)
			number = i
		}
		temp = nil
	}

	fmt.Println("-----------------------------------------------")
	fmt.Println("Lower bound:", l, "Upper bound:", h)
	fmt.Println("Longest:", longest, "Number:", number)
	fmt.Println("-----------------------------------------------")
}

//Liczba o najdłuższym ciągu z zakresów 1000-2000, 2000-3000, itd. jest zawsze nieparzysta
