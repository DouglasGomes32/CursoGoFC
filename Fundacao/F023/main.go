package main

import "fmt"

// PONTEIRO
// MEMÓRIA -> ENDEREÇO -> VALOR

func main() {

	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	fmt.Println(&a)
	b := &a
	*b = 30
	fmt.Println(a)

}
