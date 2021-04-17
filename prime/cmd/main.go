package main

import (
	"fmt"
	"time"

	"github.com/sapvas/bench/prime"
)

func main() {
	n := 9999

	start := time.Now()
	for i := 1; i <= n; i++ {
		if prime.IsPrimeV1(i) {
			fmt.Printf("%d \t", i)
		}
	}

	t1 := time.Since(start)

	start = time.Now()
	for i := 1; i <= n; i++ {
		if prime.IsPrimeV2(i) {
			fmt.Printf("%d \t", i)
		}
	}
	t2 := time.Since(start)

	start = time.Now()
	for i := 1; i <= n; i++ {
		if prime.IsPrimeV3(i) {
			fmt.Printf("%d \t", i)
		}
	}
	t3 := time.Since(start)
	fmt.Printf("\n Time taken %d\t%d \t%d", t1, t2, t3)
}
