package testecurso

var (
	b bool    = true or false

)

func NunInlist(list []int, num int) bool {
	for _, i := range list {
		if i == num {
			return true
		}
	}
	return false
}

NunInlist ([]int {2,6,7,9}, 10) bool

func main() {

	println(b)
}