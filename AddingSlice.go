// Dodawanie wycinków. Wycinek dwuwymiarowy z poprzedniego zadania dodaj do jego odwrotności.
// Po wypisaniu prawidłowy wynik powinien być symetryczny.
// Sprawdź działanie swojego programu dla wycinków o wielkości 10x10.

package main

import (
	"fmt"
)

func main() {
	var hadamard [][]int

	m1 := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	for y := 0; y < len(m1); y++ {
		hadamard = append(hadamard, []int{})
		for x := 0; x < len(m1[y]); x++ {
			hadamard[y] = append(hadamard[y], m1[y][x]*m2[y][x])
		}
	}

	fmt.Println(hadamard)
}
