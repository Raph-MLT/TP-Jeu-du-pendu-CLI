package main

import (
	"fmt"
	"pendu/src"
	"testing"
)

func TestMaPartie(t *testing.T) {
	mots, err := src.ChargerMots("mot.txt")
	if err != nil {
		t.Fatal(err)
	}
	mot := src.MotAuHasard(mots)
	fmt.Println("Mot tiré :", mot, "- lettre révélée :", src.LettreAuHasard(mot))

	err = src.enregistrerScore("Junior, Alex", "victoire", 4, mot, []string{"a", "b"})
	if err != nil {
		t.Fatal(err)
	}
	if err := src.afficherScores(); err != nil {
		t.Fatal(err)
	}
}
