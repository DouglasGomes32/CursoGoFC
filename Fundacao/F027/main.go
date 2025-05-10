package main

// TYPE ASSERTATION:

import "fmt"

func main() {
	var minhaVar interface{} = "Douglas"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é: %v", res, ok)
	res2 := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é: %v", res2, ok)

}
