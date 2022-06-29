package main

import "fmt"

func update(pointerOfNumber *int, value int) {
	*pointerOfNumber = value
}

func main() {
	number := 10
	fmt.Println(number)                                // 10
	fmt.Println("Adresse mémoire de number:", &number) // 0xc000

	myPointer := &number
	fmt.Println(myPointer)
	fmt.Printf("la valeur de l'adresse mémoire %v est %d.\n", myPointer, *myPointer)

	update(myPointer, 20)
	fmt.Printf("la valeur de l'adresse mémoire %v a changé %d.\n", myPointer, number)

}
