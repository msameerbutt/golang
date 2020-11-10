package main

import (
	"fmt"
)

func main() {
	var power int 

	power = 9000
	fmt.Printf("It's over %d\n", power)

	magic := double()
	fmt.Printf("It's over %d\n", magic)
}

func double() int {
	return 5000
}