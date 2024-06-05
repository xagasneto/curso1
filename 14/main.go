package main

func main() {
	a := 10
	//Memória -> Endereço -> Valor
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	println(*b)

}
