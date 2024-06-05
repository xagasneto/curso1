package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10, 12, 15}
	fmt.Printf("Len=%d, cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("Len=%d, cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("Len=%d, cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	fmt.Printf("Len=%d, cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 25)

	fmt.Printf("Len=%d, cap=%d %v\n", len(s[2:]), cap(s), s[2:])
	fmt.Printf("Len=%d, cap=%d %v\n", len(s), cap(s), s)

}
