package main

import "fmt"

// Endereco representa um endereço genérico.
type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// Pessoa é uma interface que define o método Desativar.
// Qualquer struct que implementar o método Desativar automaticamente implementa essa interface.
type Pessoa interface {
	Desativar()
}

// Empresa é uma struct que implementa o método Desativar.
type Empresa struct {
	Nome string
}

// Desativar da Empresa apenas satisfaz a interface Pessoa, mas não faz nenhuma ação específica.
func (e Empresa) Desativar() {
	fmt.Println("Empresa desativada (sem ação específica).")
}

// Cliente é uma struct que representa um cliente e também implementa a interface Pessoa.
type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco // Incorporação da struct Endereco (composição).
}

// Desativar desativa o cliente (define Ativo como false).
func (c Cliente) Desativar() {
	fmt.Println("Desativando cliente...")
	c.Ativo = false
	fmt.Printf("Cliente desativado: %t\n", c.Ativo)
}

// Desativacao recebe qualquer valor que implemente a interface Pessoa e chama o método Desativar.
func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	// Criando um cliente chamado Douglas
	Douglas := Cliente{
		Nome:  "Douglas",
		Idade: 27,
		Ativo: true,
	}

	// Criando uma empresa qualquer
	minhaEmpresa := Empresa{Nome: "Minha Empresa"}

	// >>> PONTO DE DEBUG IDEAL AQUI <<<
	fmt.Println("🔍 Debug: Antes de desativar empresa e cliente")

	// Desativando empresa (mesmo que não faça nada visivelmente)
	Desativacao(minhaEmpresa)

	// Desativando cliente
	Desativacao(Douglas)

	// Exibindo dados do cliente após desativação
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", Douglas.Nome, Douglas.Idade, Douglas.Ativo)
}
