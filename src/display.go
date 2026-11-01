package src

import "fmt"

func Display(pseudo, mot string, lettres []string, essais int) {
	fmt.Println(pseudo, mot, lettres, essais)

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
	fmt.Println("Mot: ", hiddenWord)
}
