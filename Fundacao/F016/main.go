package main

// FUNÇÕES

import (
	"errors"
	"fmt"
)

func main() {

	// Essa forma de fazer, é bem comum. Geralmente fazemos ja para tratar o erro

	valor, err := sum(5, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)

}

func sum(a, b int) (int, error) {
	//Fazendo isso, retorna dois valores na função, caso queira um só
	// ------------- Só deixar um, ex: func sum(a, b int) int

	if a+b >= 50 {
		return a + b, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
