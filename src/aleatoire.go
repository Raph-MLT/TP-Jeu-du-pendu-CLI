package src

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

func ChargerMots(chemin string) ([]string, error) {
	fichier, err := os.Open(chemin)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("le fichier %s est introuvable : %w", chemin, err)
		}
		return nil, fmt.Errorf("impossible d'ouvrir %s : %w", chemin, err)
	}
	defer fichier.Close()

	mots := []string{}

	lecteur := bufio.NewScanner(fichier)
	for lecteur.Scan() {
		ligne := lecteur.Text()

		ligne = strings.TrimSpace(ligne)
		ligne = strings.ToLower(ligne)

		if motValide(ligne) {
			mots = append(mots, ligne)
		}
	}

	if err := lecteur.Err(); err != nil {
		return nil, fmt.Errorf("erreur pendant la lecture de %s : %w", chemin, err)
	}

	if len(mots) == 0 {
		return nil, fmt.Errorf("le fichier %s ne contient aucun mot valide", chemin)
	}

	return mots, nil
}

func motValide(mot string) bool {
	nbLettres := 0

	for _, c := range mot {
		if c >= 'a' && c <= 'z' {
			nbLettres++
		} else if c != '-' {
			return false
		}
	}

	return nbLettres > 0
}

func MotAuHasard(mots []string) string {
	position := rand.IntN(len(mots))
	return mots[position]
}
func LettreAuHasard(mot string) string {
	lettre := ""

	for lettre == "" {
		position := rand.IntN(len(mot))
		if mot[position] != '-' {
			lettre = string(mot[position])
		}
	}

	return lettre
}
func MotAleatoire() string {
	mots, err := ChargerMots("mot.txt")
	if err != nil {
		fmt.Println("Erreur :", err)
		os.Exit(1)
	}
	return MotAuHasard(mots)
}
