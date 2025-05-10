package main

// CONDICIONAIS - if
// SWITCH / CASES - QUANDO VOCE PODE TER DIVERSAS OPÇÕES

func main() {
	a := 3
	b := 2
	c := 5

	if a > b || c > a { // || -> SIGNIFICA "ou"
		println(a)
	} else {
		println(b)
	}

	if a > b && c > a { // && SIGINIFCA "e"
		println(a)
	} else {
		println(b)
	}

	// if a > b {
	// 	println(a)
	// } else {
	// 	println(b)
	// }

}
