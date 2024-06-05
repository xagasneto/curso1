package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 3, 5, 45, 68, 456, 2369, 254))

}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero

	}
	return total
}
