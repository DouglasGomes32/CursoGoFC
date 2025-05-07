package main

// PERCORRENDO ARRAYS

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

	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	for i, v := range meuArray {
		fmt.Printf("O valor do indice %d é %d\n", i, v)

	}

}
