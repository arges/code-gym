package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["blah"] = 1
	m["blah blah"] = 2
	for k, v := range m {
		fmt.Println(v, k)
	}
	fmt.Println(m)
}
