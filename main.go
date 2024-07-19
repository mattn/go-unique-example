package main

import (
	"fmt"
	"unique"
)

func main() {
	a := [6]int{2, 1, 4, 3}
	b := unique.Make(a).Value()
	fmt.Println(a == b)
	a[0] = 3
	fmt.Println(a == b)
}
