package main

import (
	"fmt"
	// "math"
	// big "math/big"
	// "strconv"
	// "strings"
	"reflect"
)

func main() {
	var tablica [5]string = [5]string{"a", "b", "c", "d", "e"}

	tablica[0] = "f"

	var tab = make([]string, 0, 5)

	tab = append(tab, "a", "b", "c", "d", "e")

	for i, v := range tablica {
		fmt.Println(i, v)
	}

	for _, v := range tab {
		fmt.Println(v)
	}

	//print type of tab
	fmt.Println(reflect.ValueOf(tab).Kind().String())
	//print capacity of tab
	fmt.Println(cap(tab))
	fmt.Println(len(tab))
}
