package main

import "fmt"

func iiisv2(m map[uint]string) []bool {
	tab := make([]bool, 0, 51)
	highIndex := 0
	for key := range m {
		tab[key] = true
		if int(key) >= highIndex {
			highIndex = int(key)
		}
	}

	return tab
}

func main() {
	m1 := map[uint]string{1: "un", 4: "quatre"}
	r1 := iiisv2(m1)
	fmt.Println(r1) // r1 = [false, true, false, false, true]

	m2 := map[uint]string{1: "un", 2: "deux", 4: "quatre", 50: "cinquante"}
	r3 := iiisv2(m2)
	fmt.Println(r3)
}
