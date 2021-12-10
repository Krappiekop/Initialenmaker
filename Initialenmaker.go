package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	Inhoud, err := ioutil.ReadFile("Namen.txt")
	if err != nil {
		log.Fatal(err)
	}

	Naam := string(Inhoud)
	fmt.Print(Naam)

	Initiaal := strings.Split(Naam, " ")
	eersteWoord, tweedeWoord := Initiaal[0], Initiaal[1]
	eersteLetter := eersteWoord[0:1]
	laatsteLetter := tweedeWoord[0:1]

	fmt.Print(eersteLetter, ".", laatsteLetter, ".")

}
