package main

import (
	"fmt"
)

func main() {
	x := 10
	changeValue(&x)
	fmt.Println(x)
	hello()
}
func changeValue(x *int) {
	*x = 7
}
