package main

import "fmt"

const a = "hello word"

type ID int

var (
	b bool    = true
	c int     = 10
	e float64 = 1.4
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo de E é %T", e)
}