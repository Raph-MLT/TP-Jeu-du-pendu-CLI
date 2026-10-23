package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Println(ErrorsNombre)
		os.Exit(1)
	}

	command := os.Args[1]

	gameSelect(command)

}

func gameSelect(c string) {
	switch c {
	case "jouer":
		fmt.Printf("ça joue\n")
	case "reprendre":
		fmt.Printf("parti repri")
	case "score":
		score()
	default:
		fmt.Println(ErrorsBadArgument)
		os.Exit(1)
	}
}

func score() {
	contenu, err := os.ReadFile("scores.txt")
	if err != nil {
		fmt.Println(ErrorsScore)
		os.Exit(1)
	}

	fmt.Println("Scores")
	fmt.Println(string(contenu))
}
