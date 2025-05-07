package main

// CRIAÇÃO de TIPOS

import "fmt"

const a = "Hello, World"

type ID int

var (
	b bool // TAMBEM possivel por = true (No caso, ja colocar alguns valores que serao padrao)
	c int
	d string
	e float64
	f ID
)

func main() {
	// b = true

	a := "Douglas"

	fmt.Println(a, b, c, d, e, f)
}
