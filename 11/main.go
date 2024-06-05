package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	neto := Cliente{
		Nome:  "Neto",
		Idade: 40,
		Ativo: true,
	}

	fmt.Printf(neto.Nome)

}
