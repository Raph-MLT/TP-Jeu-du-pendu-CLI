package src

import (
	"fmt"
	"os"
)

func StartGame() {
	pseudo := pickUser()
	mot := MotAleatoire()
	lettres := []string{LettreAuHasard(mot)}
	essais := 6

	//fmt.Println(mot)

	for essais > 0 {
		Display(pseudo, mot, lettres, essais)

		var input string
		fmt.Scan(&input)

		var letterAlreadyFound bool = false
		for _, l := range lettres {
			if input == l {
				letterAlreadyFound = true
			}
		}
		if letterAlreadyFound {
			continue
		}
		if input == "!save" {
			fmt.Printf("Partie enregistrer ! à très vite :]\n")
			EnregistrerScore(pseudo, "?", essais, mot, lettres)
			os.Exit(0)
		}

		//fmt.Println([]rune(input)[0])
		var letterFound bool
		if len(input) == 1 && ([]rune(input)[0] >= 'a' && []rune(input)[0] <= 'z') || ([]rune(input)[0] >= 'A' && []rune(input)[0] <= 'Z') {
			lettres = append(lettres, input)
			for _, v := range mot {
				if string(v) == input {
					letterFound = true
				}
			}
			if !letterFound {
				essais -= 1
			}
		}
		if checkWork(mot, lettres) {
			fmt.Printf("Partie gagné !!!\n")
			EnregistrerScore(pseudo, "gagné", essais, mot, lettres)
			os.Exit(0)
			break
		}
	}
	fmt.Printf("Partie perdue :(\n")
	EnregistrerScore(pseudo, "perdu", essais, mot, lettres)
}

func pickUser() string {
	pseudo := ""

	fmt.Printf("Veuillez entrer un nom.\n")
	var input string
	fmt.Scan(&input)
	hasLetter := false

	for _, v := range input {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			pseudo += string(v)
			hasLetter = true
		} else if v == ' ' || v == ',' {
			pseudo += string(v)
		}
	}

	if !hasLetter {
		pickUser()
	}

	return pseudo
}

func checkWork(m string, l []string) bool {
outer:
	for _, r := range m {
		for _, w := range l {
			if r == []rune(w)[0] {
				continue outer
			}
		}
		return false
	}
	return true

}
