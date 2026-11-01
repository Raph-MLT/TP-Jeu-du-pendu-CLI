package src

import "fmt"

func StartGame() {
	fmt.Printf("Veuillez entrer un nom.\n")
	var pseudo string
	fmt.Scan(&pseudo)
	
	mot := "mariany"

	fmt.Println(mot)
}