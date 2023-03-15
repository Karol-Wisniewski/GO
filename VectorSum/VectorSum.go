// Sumowanie wektorów. Wektorem nazywa się często tablicę elementów o ustalonej długości.
// Napisz program, który tworzy w pamięci dwie tablice o rozmiarze 20 elementów, zawierające liczby rzeczywiste 64-bitowe.
// Do pierwszej tablicy wpisz w pętli wartości 2.0. Do drugiej tablicy wpisz wartości 3.0.
//  Następnie w pętli wygeneruj sumę tych dwóch tablic,
// dodając do siebie każdy element z osobna. To ćwiczenie jest "na rozgrzewkę".

package main

import (
	"fmt"
)

func main() {
	var v1 [20]float64
	var v2 [20]float64
	var sum [20]float64

	for i := 0; i < 20; i++ {
		v1[i] = 2.0
		v2[i] = 3.0
	}

	for i := 0; i < 20; i++ {
		sum[i] = v1[i] + v2[i]
	}

	fmt.Println("v1: ", v1)
	fmt.Println("v2: ", v2)
	fmt.Println("sum: ", sum)

}
