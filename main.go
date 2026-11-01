package main

import (
	"fmt"
	"os"
	"pendu/src"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Println(src.ErrorsNombre)
		os.Exit(1)
	}

	command := os.Args[1]

	gameSelect(command)

}

func gameSelect(c string) {
	switch c {
	case "jouer":
		fmt.Printf("ça joue\n")
		src.StartGame()
	case "reprendre":
		fmt.Println(src.ChargerSauvegarde()) 
	case "score":
		score()
	default:
		fmt.Println(src.ErrorsBadArgument)
		os.Exit(1)
	}
}

func score() {
	/* contenu, err := os.ReadFile("scores.txt")
	if err != nil {
		fmt.Println(src.ErrorsScore)
		os.Exit(1)
	}
	
	fmt.Println("Scores")
	fmt.Println(string(contenu)) */
	src.AfficherScores()
}
