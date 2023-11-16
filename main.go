	package main

	import (
		"fmt"
		"bufio"

		"estiam/dictionary"
	)

	func main() {
		// bufio.NewReader(os.Stdin)
		fmt.Println("Hello there")
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
	