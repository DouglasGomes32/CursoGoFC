package main

// INTERFACE
// INTERFACE EM GO, PERMITE PASSAR APENAS METODOS
// TODO MUNDO QUE TIVER O METODO Desativar(), VAI CORRESPONDER O TIPO Pessoa

import "fmt"

type Endereco struct {
	Longadouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	// QUALQUER Struc que implemente o METODO Desativar, esta implementando a interface Pessoa
	// Se tiver 10k de interface que implemente o metodo, automaticamento implementa as 10k interface
	// Interface nos possibilida usar varios tipos, de uma forma cada vez mais simples
	// Poder de abstração maior
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	// METODO para desativar o Ativo em Cliente, Tem que referenciar igual func (c Cliente)
	// SIGNIFICA QUE A FUNÇÃO Desativar() tem que ser identica a interface(Se tiver assinatura)
	c.Ativo = false
	fmt.Printf("O cliente foi desativado %t\n", c.Ativo)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
	// AQUI ELE GARANTE QUE QUALQUER PESSOA QUE TIVER O METODO DESATIVAR, VAI PODER SER USADO AQUI
}

func main() {

	Douglas := Cliente{
		Nome:  "Douglas",
		Idade: 27,
		Ativo: true,
	}

	minhaEmpresa := Empresa{}
	Desativacao(minhaEmpresa)
	// TAMBEM FUNCIONA. PQ A EMPRESA, TEM O METODO DESATIVAR
	// E O METODO DESATIVAR, EXISTE NA INTERFACE PESSOA
	// QUALQUER STRUC QUE TENHA O METODO DESATIVAR, VAI ESTAR IMPLEMENTANDO TODAS INTERFACES QUE O METODO CHAMA

	Desativacao(Douglas)
	// PQ PODE PASSAR DOUGLAS, SE O TIPO DE Desativacao() é pessoa?
	// PQ DOUGLAS É UMA Pessoa interface() que implementa o Metodo Desativar()

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", Douglas.Nome, Douglas.Idade, Douglas.Ativo)

}
