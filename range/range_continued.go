package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for key, value := range pow {
		fmt.Printf("%d - %d \n", key, value)
	}
}

// 0-1 1-2 2-4 3-8 4-16 5-32 6-64
// 7-128 8-256 9-512

// You can skip the index or value by assigning to _.

// for i, _ := range pow
// for _, value := range pow
// If you only want the index, you can omit the second variable.

// for i := range pow

// The super (possibly over) simplified definition is just that
// << is used for "(X) times 2" and
// >> is for "divided by 2" and
// the number after it is how many times.
