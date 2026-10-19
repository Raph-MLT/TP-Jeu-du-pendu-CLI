package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) <= 1 {
		//afficher l'erreur ici
		os.Exit(1)
	}

	command := os.Args[1]

	gameSelect(command)

}

func gameSelect(c string) {
	switch c{
	case "jouer":
		fmt.Printf("ça joue\n")
	case "reprendre":
		fmt.Printf("parti repri")
	case "score":
		fmt.Printf("scoreeeeds")
	default:
		//Mettre une erreur ici car tu pues ta race
		os.Exit(1)
	}
}