package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	c.nome = "Francisco Neto"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func main() {
	neto := Cliente{
		nome: "Neto",
	}
	neto.andou()
	fmt.Printf("O valor da struct com o nome é %v ", neto.nome)

}
