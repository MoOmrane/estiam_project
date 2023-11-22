package main

import (
	"bufio"
	"fmt"
	"os"

	"dictionary/dictionary" 
)

func main() {
	filePath := "dictionary_data.txt"
	dict := dictionary.New(filePath)

	// Ajout d'entrées au dictionnaire
	err := dict.Add("opel", "corsa")
	if err != nil {
		fmt.Println("Erreur lors de l'ajout :", err)
	}

	// Obtention d'une définition
	entry, err := dict.Get("opel")
	if err != nil {
		fmt.Println("Erreur lors de l'obtention de la définition :", err)
	} else {
		fmt.Println("Définition de 'opel' :", entry.String())
	}

	// Suppression d'une entrée
	err = dict.Remove("opel")
	if err != nil {
		fmt.Println("Erreur lors de la suppression :", err)
	}

	// Liste de toutes les entrées
	words, entries, err := dict.List()
	if err != nil {
		fmt.Println("Erreur lors de la liste des entrées :", err)
	} else {
		fmt.Println("Mots dans le dictionnaire :")
		for _, word := range words {
			fmt.Printf("%s: %s\n", word, entries[word].String())
		}
	}
}
