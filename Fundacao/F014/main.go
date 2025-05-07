package main

import "fmt"

// SLICE PARECIDO COM UM ARRAY INFINITO
// cap() CAPACIDADE DO SLICE

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:]) // Ignorando o começo, a capacidade diminuiu

	s = append(s, 12)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// Depois de adicionar(append) o SLICE vai dobrar o tamanho antes do append

}
