package main

import (
	"fmt"

	"exo_matt/change"
	likereflect "exo_matt/like_reflect"
)

func main() {
	main2()
}

func main2() {
	var changes []change.Change

	s1 := likereflect.NewS_old(1, 2, 0)
	// s1 := likereflect.NewS("structure1", 1, 2, 0)
	fmt.Println(s1.PrintStruct())

	s3 := likereflect.NewS_old(2, 2, 0)
	// s3 := likereflect.NewS("structure2", 2, 2, 0)
	fmt.Println(s3.PrintStruct())

	changes = change.Compare(s1, s3)

	fmt.Println("===> compare s1/s3 :")

	change.PrintChanges(changes)

	fmt.Println("==== exo :")
	fmt.Println("s1 :")

	if dec, err := s1.PrintDeclaration(); err == nil {
		fmt.Println(dec)
	} else {
		fmt.Printf("/!\\ Got an error: %+v\n", err)
	}

	fmt.Println("s3 :")
	fmt.Println(s3.PrintDeclaration())

	fmt.Println(s1.Field(0))

}
