package src

import "fmt"

func Display(pseudo, mot string, lettres []string, essais int) {
	//fmt.Println(pseudo, mot, lettres, essais)

	hiddenWord := ""
	outer:
	for i := 0; i < len(mot); i++ {
		for _, l := range lettres {
			if string(mot[i]) == l {
				hiddenWord += l + " "
				continue outer
			}
		}
		hiddenWord += "_ "
	}
	fmt.Printf("JOUEUR: %s\n\n", pseudo)
	fmt.Printf("Mot: %s\n", hiddenWord)
	fmt.Println("Lettres proposés: ", lettres)
	fmt.Printf("Essais restant: %d\n", essais)
}
