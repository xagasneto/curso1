package main

import (
	"errors"
	"fmt"
)

func main() {
	
	fmt.Println(valor)

}

func sum(numeros ... int) int{
	total := 0
	for_, numero := range numeros {
		total += numero
	} 
	return total
		

}
