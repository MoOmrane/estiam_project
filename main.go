	package main

	import (
		"fmt"
		"bufio"
		"strings"
		"estiam/dictionary"
		)

	func main() {
		dict := dictionary.New()
	
		dict.Add("opel", "corsa")
		dict.Add("renault", "clio")
	
		// definitions
		word := "opel"
		entry, err := dict.Get(word)
		if err != nil {
			fmt.Printf("Error getting definition for '%s': %s\n", word, err)
		} else {
			fmt.Printf("Definition of '%s': %s\n", word, entry.String())
		}
	
		// Supprimer un element
		wordToRemove := "renault"
		dict.Remove(wordToRemove)
		fmt.Printf("'%s' has been removed from the dictionary.\n", wordToRemove)
	
		// List
		words, entries := dict.List()
		fmt.Println("Words in the dictionary:")
		for _, word := range words {
			fmt.Printf("%s: %s\n", word, entries[word].String())
		}
	}
	

	func actionAdd(d *dictionary.Dictionary, reader *bufio.Reader) {
		fmt.Print("Mot : ")
		word, _ := reader.ReadString('\n')
		fmt.Print("Définition : ")
		definition, _ := reader.ReadString('\n')
		d.Add(strings.TrimSpace(word), strings.TrimSpace(definition))
	}
		

	func actionDefine(d *dictionary.Dictionary, reader *bufio.Reader) {
		fmt.Print("Mot à définir : ")
		word, _ := reader.ReadString('\n')
		entry, err := d.Get(strings.TrimSpace(word))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Définition de '%s' : %s\n", word, entry.String())
	}

	func actionRemove(d *dictionary.Dictionary, reader *bufio.Reader) {
		fmt.Print("Mot à supprimer : ")
		word, _ := reader.ReadString('\n')
		d.Remove(strings.TrimSpace(word))
		fmt.Printf("'%s' a été supprimé du dictionnaire.\n", strings.TrimSpace(word))
	}	

	func actionList(d *dictionary.Dictionary) {
		words, entries := d.List()
		fmt.Println("Mots dans le dictionnaire :")
		for _, word := range words {
			fmt.Printf("%s: %s\n", word, entries[word].String())
		}
	}
	