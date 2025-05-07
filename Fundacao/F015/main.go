package main

// Maps

import "fmt"

func main() {

	salarios := map[string]int{
		"Douglas": 1000,
		"João":    2000,
		"Maria":   300,
	}

	fmt.Println(salarios)
	delete(salarios, "Maria")
	fmt.Println(salarios)
	salarios["Maria"] = 200
	fmt.Println(salarios)
	fmt.Println(salarios["Maria"])

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %v\n", nome, salario)
	}

	for _, salario := range salarios { // Usando o _ para ignorar o valor e nao quebrar nada
		fmt.Printf("O salario é %v\n", salario)
	}

	sal := make(map[string]int) // Diferença com make() é a sintax, mas vai dar certo tambem sem o make
	sal1 := map[string]int{}    // igual aqui nesse exemplo
	sal1["Gomes"] = 1000
	sal["Pedro"] = 2000

	fmt.Println(sal)
	fmt.Println(sal1)
}
