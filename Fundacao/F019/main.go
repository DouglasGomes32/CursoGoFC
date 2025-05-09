package main

// STRUCT

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	Douglas := Cliente{
		Nome:  "Douglas",
		Idade: 27,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", Douglas.Nome, Douglas.Idade, Douglas.Ativo)

}
