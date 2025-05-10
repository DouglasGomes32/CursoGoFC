package main

// GENERIC
// comparable SÓ COMPARA
// TRABALHANDO COM TIPAGEM DE FORMA CORRETA, SEM REPETIR CODIGO IGUAL A PARTE DEBAIXO AQUI

type Number interface {
	int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Douglas": 1000, "Pedro": 1234, "Andressa": 5678}
	f := map[string]float64{"Douglas": 1.0, "Pedro": 2.0, "Andressa": 3.0}

	M3 := map[string]int{"Douglas": 1000, "Pedro": 1234, "Andressa": 5678}

	println(Soma(m))
	println(Soma(f))
	println(Soma(M3))
	println(Compara(10, 10))

}

// .
// ..
// ..
// ..
// ..
// ..
// ..
// ..
// ..
// ..
// ..
// ..
// ..
// ..
// ..
// ..
// ..
// .

// Jeito errado. Pois, esta criando duas funções só pq o jeito é diferente

// func SomaInteiro(m map[string]int) int {
// 	var soma int
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

// func SomaFloat(m map[string]float64) float64 {
// 	var soma float64
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

// func main() {
// 	m := map[string]int{"Douglas": 1000, "Pedro": 1234, "Andressa": 5678}
// 	f := map[string]float64{"Douglas": 1.0, "Pedro": 2.0, "Andressa": 3.0}
// 	println(SomaInteiro(m))
// 	println(SomaFloat(f))

// }
