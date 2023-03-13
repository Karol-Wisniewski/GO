package main

import (
	"fmt"
	"strconv"
	"strings"
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

func getLongestColltazSequence(l int, h int) {
	//slice to hold the Collatz sequence for each number in range
	temp := []int{}

	//length of the longest Collatz sequence in provided range
	var longest int

	//number that produces the longest Collatz sequence in provided range
	var number int

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

func generateCollatzSequences(n int) map[int]int {
	temp := make(map[int][]int)
	result := make(map[int]int)
	for i := 1; i < n+1; i++ {
		temp[i] = collatz(i, temp[i])
	}

	for i := 1; i < len(temp)+1; i++ {
		result[i] = len(temp[i])
	}

	return result
}

func generateCollatzLengths(n int) []int {
	temp := []int{}
	result := []int{}

	for i := 1; i < n+1; i++ {
		temp = collatz(i, temp)
		result = append(result, len(temp))
		temp = nil
	}

	return result
}

func flattenArray(arr [][]int) []int {

	// Initialize the flattened array
	flatArr := []int{}

	// Iterate through the rows of the array
	for _, row := range arr {

		// Append the current row to the flattened array
		flatArr = append(flatArr, row...)
	}
	return flatArr
}

func main() {
	// getLongestColltazSequence(1000, 2000)
	// sequencesUnder1k := generateCollatzSequences(100000)

	// file, err := os.Create("Collatz.dat")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for i := 1; i < len(sequencesUnder1k)+1; i++ {
	// 	file.WriteString(fmt.Sprintf("%d %d", i, sequencesUnder1k[i]))
	// 	file.WriteString("\n")
	// }

	// file.Close()

	lengths := generateCollatzLengths(10000)

	// fmt.Println(lengths)

	//remake "lengths" slice to hold the number of times each digit in each length occurs

	tempArray := []int{}

	for i := 0; i < len(lengths); i++ {

		tempStrArray := strings.Split(strconv.Itoa(lengths[i]), "")

		for i := 0; i < len(tempStrArray); i++ {
			temp, _ := strconv.Atoi(tempStrArray[i])
			tempArray = append(tempArray, temp)
		}
	}

	lengthsMap := make(map[int]int)

	for i := 0; i < len(tempArray); i++ {
		lengthsMap[tempArray[i]]++
	}

	fmt.Println(lengthsMap)

}

//Number with the longest Collatz sequence in the range 1000-2000, 2000-3000, etc. is always odd

//Probability of a number occuring in a Collatz sequence is simillar to Benford's Law
