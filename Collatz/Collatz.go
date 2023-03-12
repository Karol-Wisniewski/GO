package main

import (
	"fmt"
	"gnuplot"
	"log"
	"math"
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

func main() {
	// getLongestColltazSequence(1000, 2000)
	// sequencesUnder1k := generateCollatzSequences(100000)
	// Create a new gnuplot session
	plot, err := gnuplot.NewPlotter("", true, false)
	if err != nil {
		log.Fatal(err)
	}
	defer plot.Close()

	// Set the output file to PNG format
	plot.CheckedCmd("set terminal png")

	// Set the output filename
	plot.CheckedCmd("set output 'plot.png'")

	// Plot the sine function
	x := make([]float64, 0)
	y := make([]float64, 0)
	for i := 0; i < 100; i++ {
		x = append(x, float64(i)/10.0)
		y = append(y, math.Sin(float64(i)/10.0))
	}
	plot.PlotXY(x, y, "Sine wave")
}

//Liczba o najdłuższym ciągu z zakresów 1000-2000, 2000-3000, itd. jest zawsze nieparzysta
