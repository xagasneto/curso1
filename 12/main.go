package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {

	neto := Cliente{
		Nome:  "Neto",
		Idade: 40,
		Ativo: true,
	}
	neto.Cidade = "Fortaleza"

	fmt.Printf("Cidade: %s", neto.Cidade)

}
