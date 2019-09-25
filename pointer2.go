package main

import "fmt"

func main() {
	var p = f()
	fmt.Println(p) //each time p() function is called return different value
}
func f() *int {
	v := 1
	return &v
}
