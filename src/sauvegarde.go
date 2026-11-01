package src

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Sauvegarde struct {
	Pseudo  string   `json:"pseudo"`
	Mot     string   `json:"mot"`
	Lettres []string `json:"lettres"`
	Essais  int      `json:"essais"`
}

func sauvegarder(pseudo string, mot string, lettres []string, essais int) error {
	s := Sauvegarde{Pseudo: pseudo, Mot: mot, Lettres: lettres, Essais: essais}

	donnees, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("impossible de créer le JSON : %w", err)
	}

	err = os.WriteFile("save.json", donnees, 0644)
	if err != nil {
		return fmt.Errorf("impossible d'écrire save.json : %w", err)
	}

	return nil
}

func chargerSauvegarde() (Sauvegarde, error) {
	var s Sauvegarde

	donnees, err := os.ReadFile("save.json")
	if errors.Is(err, os.ErrNotExist) {
		return s, fmt.Errorf("aucune partie sauvegardée (save.json introuvable) : %w", err)
	}
	if err != nil {
		return s, fmt.Errorf("impossible de lire save.json : %w", err)
	}

	err = json.Unmarshal(donnees, &s)
	if err != nil {
		return s, fmt.Errorf("save.json est illisible : %w", err)
	}

	if s.Pseudo == "" || s.Mot == "" || len(s.Lettres) == 0 || s.Essais < 1 || s.Essais > 6 {
		return s, errors.New("save.json est incomplet")
	}

	err = os.Remove("save.json")
	if err != nil {
		return s, fmt.Errorf("impossible de supprimer save.json : %w", err)
	}

	return s, nil
}
