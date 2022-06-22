package main

import "fmt"

func iiisv2(m map[uint]string) (int, []bool) {
	tab := make([]bool, 0, 51)
	highIndex := 0
	for key := range m {
		tab[key] = true
		if int(key) >= highIndex {
			highIndex = int(key)
		}
	}

	return highIndex, tab
}

func indexIsInSlice(m map[uint]string) []bool {
	// Select High Index
	highIndex := 0
	for key := range m {
		if int(key) >= highIndex {
			highIndex = int(key)
		}
	}

	// Check key exist?
	// var tab []bool
	tab := make([]bool, 0, highIndex)
	for key := range m {
		tab[key] = true
	}
	for i := 0; i <= highIndex; i++ {
		if _, ok := m[uint(i)]; ok {
			tab = append(tab, true)
		}
	}

	return tab
}

func main() {
	m1 := map[uint]string{1: "un", 4: "quatre"}
	r1 := indexIsInSlice(m1)
	fmt.Println(r1) // r1 = [false, true, false, false, true]

	m2 := map[uint]string{1: "un", 2: "deux", 4: "quatre", 50: "cinquante"}
	r2 := indexIsInSlice(m2)
	fmt.Println(r2)
}
