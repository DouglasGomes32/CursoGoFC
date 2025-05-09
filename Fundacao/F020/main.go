package main

// COMPOSIÇÃO de STRUCT

import "fmt"

type Endereco struct {
	Longadouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {

	Douglas := Cliente{
		Nome:  "Douglas",
		Idade: 27,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", Douglas.Nome, Douglas.Idade, Douglas.Ativo)

}
