package main

// PONTEIRO COM STRUCT
// QUANDO USAR PONTEIRO
// NAO USAR PARA COISA RAPIDA E RAZA,
// USAR QUANDO FOR ALGO MUTAVEL
// PONTEIRO
// MEMÓRIA -> ENDEREÇO -> VALOR

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) simular(valor int) int {

	c.saldo += valor
	println(c.saldo)
	return c.saldo

}

func main() {
	conta := Conta{saldo: 100}
	conta.simular(200)
	println(conta.saldo)

}
