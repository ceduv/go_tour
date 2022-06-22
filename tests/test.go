package main

import "fmt"

func main() {
	test := "ced"
	fmt.Printf(test + "\n")
	fmt.Printf("%v \n", test)
	fmt.Printf("%v \n", test)

	tab := make([]bool, 6, 51)
	fmt.Println(tab)

	tab = tab[:4]
	fmt.Println(tab, len(tab), cap(tab))
}
