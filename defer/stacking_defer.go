package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

// counting done 9 8 7 6 5 4 3 2 1 0

// Deferred function calls are pushed onto a stack. When a function returns,
// its deferred calls are executed in last-in-first-out order.
// To learn more about defer statements read this
// https://go.dev/blog/defer-panic-and-recover
