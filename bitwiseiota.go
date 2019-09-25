package main

import (
	"fmt"
)

const (
	isAdmin = 1 << iota
	isHeadquaters
	canSeeFinancial

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorth
	canSeeSouth
)

func main() {
	var roles byte = isAdmin | canSeeFinancial | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("%b\n", isAdmin)
	fmt.Printf("Is Admin ?%v\n", isAdmin&roles == isAdmin)
}
