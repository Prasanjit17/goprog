package main

import (
	"fmt"
)

func main() {
	var i int = 42
	fmt.Println("%v, %T", i, i)

	var j float32
	j = float32(i)
	fmt.Println("%v %T", j, j)
}
