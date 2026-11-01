package src

import "fmt"

func StartGame() {
	fmt.Printf("Veuillez entrer un nom.\n")
	pseudo := pickUser()
	mot := "mariany"
	lettres := []string{}
	essais := 6

	fmt.Println(mot)
	

	for essais > 0 {
		Display(pseudo, mot, lettres, essais)

		var input string
		fmt.Scan(&input)
		if len(input) == 1 && []rune(input)[0] >= 'a' && []rune(input)[0] <= 'z' && []rune(input)[0] >= 'A' && []rune(input)[0] <= 'Z' {
			lettres = append(lettres, input)
			for _, v := range mot {
				if string(v) != input {
					essais -= 1
				}
			}
		}
	}

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