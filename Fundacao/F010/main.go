package main

// DECLARAÇÃO e ATRIBUIÇÃO

import "fmt"

const a = "Hello, World"

var (
	b bool // TAMBEM possivel por = true (No caso, ja colocar alguns valores que serao padrao)
	c int
	d string
	e float64
)

func main() {
	// b = true

	a := "Douglas"
	fmt.Println(a, b, c, d, e)
}
