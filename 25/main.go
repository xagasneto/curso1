package main

import "os"

func main() {
	f, err := os.Create("Arquivo.txt")
	if err != nil {
		panic(err)
	}
	f.Close()

}
