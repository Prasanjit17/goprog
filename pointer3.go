package main

import (
	"fmt"
)

func main() {
	v := 1
	incr(&v)
	fmt.Println(incr(&v))
	fmt.Println(v)
}
func incr(p *int) int {
	*p++
	fmt.Println("Inside", *p)
	return *p
}
