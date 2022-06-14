package main

import "fmt"

// return x AND y
func swap(x, y string) (string, string) {
	return y, x
}

// named == "naked" return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
