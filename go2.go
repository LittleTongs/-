package main

import (
	"fmt"
)

func qwe(a int, b int) int {
	return a + b
}

func rty(a int, b int, qwe func(int, int) int) int {
	return qwe(a, b)
}

func main() {
	a := 0
	b := 0
	fmt.Scanf("%d %d\n", &a, &b)
	c := rty(a, b, qwe)
	fmt.Printf("%d\n", c)
}
