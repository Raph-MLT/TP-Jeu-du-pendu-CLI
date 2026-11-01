package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const fichierScores = "scores.csv"

func enregistrerScore(pseudo string, statut string, essais int, mot string, lettres []string) error {
	fichierExiste := true
	_, err := os.Stat(fichierScores)
	if errors.Is(err, os.ErrNotExist) {
		fichierExiste = false
	} else if err != nil {
		return fmt.Errorf("impossible de vérifier %s : %w", fichierScores, err)
	}

	fichier, err := os.OpenFile(fichierScores, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("impossible d'ouvrir %s : %w", fichierScores, err)
	}
	defer fichier.Close()

	ecrivain := csv.NewWriter(fichier)

	if !fichierExiste {
		enTete := []string{"pseudo", "statut", "essais_restants", "mot", "lettres", "date"}
		if err := ecrivain.Write(enTete); err != nil {
			return fmt.Errorf("impossible d'écrire l'en-tête : %w", err)
		}
	}

	ligne := []string{
		pseudo,
		statut,
		fmt.Sprint(essais),
		mot,
		strings.Join(lettres, ""),
		time.Now().Format("2006-01-02 15:04"),
	}
	if err := ecrivain.Write(ligne); err != nil {
		return fmt.Errorf("impossible d'écrire dans %s : %w", fichierScores, err)
	}

	ecrivain.Flush()
	if err := ecrivain.Error(); err != nil {
		return fmt.Errorf("impossible d'écrire dans %s : %w", fichierScores, err)
	}

	return nil
}

func afficherScores() error {
	fichier, err := os.Open(fichierScores)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("aucun historique (%s introuvable) : %w", fichierScores, err)
		}
		return fmt.Errorf("impossible d'ouvrir %s : %w", fichierScores, err)
	}
	defer fichier.Close()

	lecteur := csv.NewReader(fichier)
	lecteur.FieldsPerRecord = -1
	lecteur.LazyQuotes = true

	parties := [][]string{}
	numero := 0

	for {
		ligne, err := lecteur.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("erreur pendant la lecture de %s : %w", fichierScores, err)
		}

		numero++

		if numero == 1 {
			continue
		}

		if len(ligne) != 6 {
			fmt.Printf("Ligne %d mal formée : %d colonnes au lieu de 6\n", numero, len(ligne))
			continue
		}

		parties = append(parties, ligne)
	}

	if len(parties) == 0 {
		return fmt.Errorf("aucune partie terminée dans %s", fichierScores)
	}

	fmt.Printf("%-15s %-10s %-8s %-15s %s\n", "PSEUDO", "STATUT", "ESSAIS", "MOT", "DATE")
	for _, p := range parties {
		fmt.Printf("%-15s %-10s %-8s %-15s %s\n", p[0], p[1], p[2], p[3], p[5])
	}

	return nil
}
