package main

// !!! UTILISEZ LA LIBRARY LOG PERSO
import (
	"log" // !!!

	"gitlab.zebestof.com/zbo/reactor/libs/utils/log"
)

func main() {
	name := "ced"
	log.Errorf("test %+v", name)
}
