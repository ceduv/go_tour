package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

// 1 2 true false no!

// A var declaration can include initializers, one per variable.

// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
