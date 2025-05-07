package main

import "fmt"

// FUNÇÕES ANONIMAS (CLOSURES) - DENTRO DE UMA FUNÇÃO, TER OUTRA FUNÇÃO

func main() {

	total := func() int {
		return sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0

	for _, numeros := range numeros {
		total += numeros
	}
	return total
}
