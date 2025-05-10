package main

// PACOTES e MODULOS

import (
	"fmt"

	"github.com/DouglasGomes32/CursoGoFC/Fundacao/F029/matematica"
	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Println(s)
	fmt.Println(uuid.New())
}
