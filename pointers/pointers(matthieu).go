package main

import "fmt"

type Square struct {
	b int
}

func (s Square) draw() {
	fmt.Printf("[in draw] j'ai comme pointeur: %p\n", &s)
	s.b = 2
}

func (s *Square) draw2() {
	fmt.Printf("[in draw2] j'ai comme pointeur: %p\n", &s)
	s.b = 2
}

func main() {
	s := Square{b: 1}
	fmt.Printf("[in main] j'ai comme pointeur: %p\n", &s)
	s.draw()
	fmt.Printf("s.b = %d\n", s.b) // still 1
	s.draw2()
	fmt.Printf("s.b = %d\n", s.b) // now 2

}
