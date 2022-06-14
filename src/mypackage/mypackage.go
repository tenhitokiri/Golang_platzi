package mypackage

import (
	"fmt"
)

//CarPublic is a public Struct
type CarPublic struct {
	Name string
	Brand string
	Year int
}

//CarPublic is a public Struct
type carPrivate struct {
	name string
	brand string
	year int
}

//PrintMessage Imprime un mensaje en pantalla
func PrintMessage(message string) {
	fmt.Println(message)
}

func printMessage1(message string) {
	fmt.Println(message)
}
