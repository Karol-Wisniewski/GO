package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

type T struct{ a, b, c float64 }

func p(t T) (r [2]float64, e bool) {
	d := t.b*t.b - 4*t.a*t.c
	if d < 0 {
		return
	}
	e = true
	s := math.Sqrt(d)
	r[0], r[1] = (-t.b+s)/(2*t.a), (-t.b-s)/(2*t.a)
	return
}
func main() {
	a, b, c := flag.Float64("a", 1, ""), flag.Float64("b", 1, ""), flag.Float64("c", 1, "")
	flag.Parse()
	if *a == 0 {
		os.Exit(1)
	}
	r, e := p(T{*a, *b, *c})
	if !e {
		fmt.Println("Brak")
	} else {
		fmt.Printf("x1=%.2f, x2=%.2f\n", r[0], r[1])
	}
}
