package main

import "fmt"

// QUANDO USAR PONTEIRO
// NAO USAR PARA COISA RAPIDA E RAZA,
// USAR QUANDO FOR ALGO MUTAVEL
// PONTEIRO
// MEMÓRIA -> ENDEREÇO -> VALOR

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	soma(&minhaVar1, &minhaVar2)
	fmt.Println(minhaVar1)

}
