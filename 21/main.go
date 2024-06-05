package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	//T representa o generic
	var soma T

	for _, v := range m {
		soma += v
	}
	return soma
}

func compara[T Number](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Weslwy": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Weslwy": 1000.1, "João": 2000.2, "Maria": 3000.3}

	m3 := map[string]MyNumber{"Weslwy": 1000, "João": 2000, "Maria": 3000}
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))

}
