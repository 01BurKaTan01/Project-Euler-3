package main

import (
	"fmt"
	"math"
)

func main() {
	sayim := 600851475143
	asalim := 0
	for sayim != 1 {
		asalim = siradakiAsal(asalim)
		if sayim%asalim == 0 {
			sayim = sayim / asalim
			fmt.Println(asalim)
			asalim = 2
		}
	}
}

func asalMi(sayi int) bool {
	if sayi < 2 {
		return false
	}
	ust := int(math.Sqrt(float64(sayi)))
	for i := 2; i <= ust; i++ {
		if sayi%i == 0 {
			return false
		}
	}
	return true
}

func siradakiAsal(n int) int {
	if n < 2 {
		return 2
	}

	for {
		n++
		if asalMi(n) {
			return n
		}
	}
}
