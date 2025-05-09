package main

// MÉTODOS em STRUCT

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

func (c Cliente) Desativar() { // METODO para desativar o Ativo em Cliente, Tem que referenciar igual func (c Cliente)
	c.Ativo = false
	fmt.Printf("O cliente foi desativado %t\n", c.Ativo)
}

func main() {

	Douglas := Cliente{
		Nome:  "Douglas",
		Idade: 27,
		Ativo: true,
	}
	Douglas.Ativo = false
	Douglas.Desativar()

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", Douglas.Nome, Douglas.Idade, Douglas.Ativo)

}
