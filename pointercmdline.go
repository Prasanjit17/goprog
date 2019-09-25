package main

import (
	"flag"
	"fmt"
)

var n = flag.Bool("n", false, "omit trailing new lline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(string.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
