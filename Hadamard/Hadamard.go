// Iloczyn Hadamarda. Po pierwsze, sprawdź czym jest ten iloczyn.
// Następnie napisz w swoich Notatkach, co udało Ci się na ten temat zrozumieć,
// i napisz krótki program, który ilustruje zastosowanie tego mnożenia.
// Możesz to zrobić na tablicach o niewielkiej długości, np. do 5 elementów.

package main

import (
	"fmt"
)

func main() {
	var hadamard [][]int

	m1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	m2 := [][]int{
		{2, 2, 2},
		{2, 2, 2},
		{3, 3, 4},
	}

	for y := 0; y < len(m1); y++ {
		hadamard = append(hadamard, []int{})
		for x := 0; x < len(m1[y]); x++ {
			hadamard[y] = append(hadamard[y], m1[y][x]*m2[y][x])
		}
	}

	fmt.Println(hadamard)
}
