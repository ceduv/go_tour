package main

type ms struct {
	field1 string ``                 // Non Tracké
	field2 string `watch:""`         // Compare le champ
	field3 ms2    `watch:"in_depth"` // permet de tracké les champs watch ds la structure liée
	field4 string `watch:"always"`   // enregistre la modif ds tous les cas
}

type ms2 struct {
	i1 int
	i2 int
}

func main() {

}
