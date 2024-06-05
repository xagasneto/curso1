package main

import (
	"fmt"
	"github/xagasneto/curso1/matematica"
	// tem que trazer aqui pra dentro o módulo go.mod criado no github.com
	// a partir daí, consegue acessar os pacotes que estão em outros diretórios
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Println("O resultado:", s)

}
