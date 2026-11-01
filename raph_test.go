package main

import (
	"fmt"
	"testing"
)

func TestMaPartie(t *testing.T) {
	mots, err := ChargerMots("mot.txt")
	if err != nil {
		t.Fatal(err)
	}
	mot := motAuHasard(mots)
	fmt.Println("Mot tiré :", mot, "- lettre révélée :", lettreAuHasard(mot))

	err = enregistrerScore("Junior, Alex", "victoire", 4, mot, []string{"a", "b"})
	if err != nil {
		t.Fatal(err)
	}
	if err := afficherScores(); err != nil {
		t.Fatal(err)
	}
}
