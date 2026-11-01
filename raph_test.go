package main

import (
	"fmt"
	"testing"
	"pendu/src"
)

func TestMaPartie(t *testing.T) {
	mots, err := src.ChargerMots("mot.txt")
	if err != nil {
		t.Fatal(err)
	}
	mot := src.MotAuHasard(mots)
	fmt.Println("Mot tiré :", mot, "- lettre révélée :", src.LettreAuHasard(mot))

	err = enregistrerScore("Junior, Alex", "victoire", 4, mot, []string{"a", "b"})
	if err != nil {
		t.Fatal(err)
	}
	if err := afficherScores(); err != nil {
		t.Fatal(err)
	}
}
