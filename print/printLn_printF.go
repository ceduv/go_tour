package main

import "fmt"

func main() {
	m, n, p := 15, 25, 40

	fmt.Println(
		"(m + n = p) :", m, "+", n, "=", p,
	)

	fmt.Printf(
		"(m + n = p) : %d + %d = %d\n", m, n, p,
	)
}

// (m + n = p) : 15 + 25 = 40
// (m + n = p) : 15 + 25 = 40
