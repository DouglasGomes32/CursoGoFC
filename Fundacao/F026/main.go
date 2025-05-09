package main

import "fmt"

// INTERFACE VAZIA
// INTERFACE EM GO, PERMITE PASSAR APENAS METODOS
// TODO MUNDO QUE TIVER O METODO Desativar(), VAI CORRESPONDER O TIPO Pessoa

// INTERFACE VAZIA SIGNIFICA QUE ELA IMPLEMENTA TODO TIPO DE VARIAVEL
type x interface{}

func main() {

	var x interface{} = 10
	var y interface{} = "Hello, World!"

	showType(x)
	showType(y)

}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é: %T e o valor é %v\n", t, t)
}
