package main

import "fmt"

func main() {
	salarios := map[string]int{"NETO": 15000, "JOÃO": 20000, "MARIA": 25000}

	delete(salarios, "NETO")
	salarios["NET"] = 17000

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}

}
