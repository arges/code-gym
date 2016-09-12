package main

import (
	"fmt"
)

func bin2int(bin string) int {
	val := 0
	for i, k := range bin {
		if string(k) == "1" {
			val += 2 ^ i
		}
	}
	return val
}

func main() {
	bin := "010101010101"
	fmt.Printf("%#x\n", bin2int(bin))
}
